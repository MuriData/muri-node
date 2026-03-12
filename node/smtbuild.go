package node

import (
	"context"
	"fmt"
	"math/big"
	"runtime"
	"sync"
	"time"

	"github.com/MuriData/muri-node/prover"
	"github.com/MuriData/muri-zkproof/circuits/poi"
	"github.com/MuriData/muri-zkproof/pkg/merkle"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/rs/zerolog/log"
)

// fetchTimeout computes a generous IPFS fetch deadline based on file size.
// Base of 5 minutes + 3 seconds per MB accounts for CatChunked's per-segment
// retry (each segment can take up to 2 min response header timeout + retries).
func fetchTimeout(numChunks uint32) time.Duration {
	const chunkSize = 16384 // 16 KB
	sizeMB := uint64(numChunks) * chunkSize / (1024 * 1024)
	if sizeMB < 1 {
		sizeMB = 1
	}
	return 5*time.Minute + time.Duration(sizeMB)*3*time.Second
}

// buildSMTStreaming downloads a file from IPFS and builds an SMT by hashing
// chunks on the fly. The full file is never buffered in memory — only one
// download segment (~1 MB) plus the leaf hashes are held at any time.
//
// Peak memory: ~12 MB for a 1 GB file, ~124 MB for a 10 GB file
// (vs ~1 GB / ~10 GB with the old CatChunked + BuildSMT approach).
func (n *Node) buildSMTStreaming(ctx context.Context, ref string, numChunks uint32) (*merkle.SparseMerkleTree, error) {
	type hashJob struct {
		index int
		data  []byte
	}

	jobs := make(chan hashJob, 128)

	// Pre-allocate the results slice. Each index is written by exactly one
	// worker (disjoint indices), so no synchronization is needed for writes.
	hashCap := int(numChunks) + 1
	hashes := make([]fr.Element, hashCap)

	// Start parallel hash workers
	workers := runtime.NumCPU()
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				if j.index < hashCap {
					hashes[j.index] = poi.HashChunk(j.data)
				}
			}
		}()
	}

	// Stream download: fetch each 16 KB chunk and dispatch for hashing.
	// CatChunkedTo holds only ~1 MB (one download segment) at a time.
	totalChunks, err := n.ipfs.CatChunkedTo(ctx, ref, poi.FileSize, func(index int, chunk []byte) {
		// Copy chunk data — the slice is only valid during this callback.
		data := make([]byte, len(chunk))
		copy(data, chunk)
		select {
		case jobs <- hashJob{index: index, data: data}:
		case <-ctx.Done():
		}
	})
	close(jobs)
	wg.Wait()

	if err != nil {
		return nil, fmt.Errorf("streaming download: %w", err)
	}

	if totalChunks > hashCap {
		return nil, fmt.Errorf("chunk count %d exceeds expected %d", totalChunks, numChunks)
	}
	hashes = hashes[:totalChunks]

	log.Debug().Int("chunks", totalChunks).Msg("streaming hash complete")

	smt, err := merkle.BuildSMTFromLeafHashes(hashes, poi.MaxTreeDepth, n.prover.ZeroLeafHash())
	if err != nil {
		return nil, fmt.Errorf("build SMT: %w", err)
	}

	return smt, nil
}

// buildSMTFromRawBlock fetches a raw IPFS block (e.g. a directory DAG node)
// via block/get and builds an SMT from its bytes. Raw blocks are typically
// small (< 1 MB) so no streaming is needed.
func (n *Node) buildSMTFromRawBlock(ctx context.Context, ref string) (*merkle.SparseMerkleTree, error) {
	data, err := n.ipfs.BlockGetWithRetry(ctx, ref)
	if err != nil {
		return nil, fmt.Errorf("block get: %w", err)
	}

	chunks := merkle.SplitIntoChunks(data, poi.FileSize)
	if len(chunks) == 0 {
		return nil, fmt.Errorf("empty block data for %s", ref)
	}

	hashes := make([]fr.Element, len(chunks))
	for i, chunk := range chunks {
		hashes[i] = poi.HashChunk(chunk)
	}

	smt, err := merkle.BuildSMTFromLeafHashes(hashes, poi.MaxTreeDepth, n.prover.ZeroLeafHash())
	if err != nil {
		return nil, fmt.Errorf("build SMT: %w", err)
	}

	return smt, nil
}

// proveSelective generates a proof by fetching only the 8 challenged chunks
// (via IPFS byte-range requests) instead of downloading the entire file.
// For raw block URIs, fetches the whole block (typically tiny) and slices locally.
func (n *Node) proveSelective(ctx context.Context, randomness *big.Int, smt *merkle.SparseMerkleTree, ref string, rawBlock bool) (*prover.ProofResult, error) {
	// Derive the 8 leaf indices the circuit will open
	indices := prover.DeriveLeafIndices(randomness, smt.NumLeaves)

	// Collect unique indices (small files may repeat via modular wrapping)
	unique := make(map[int]struct{})
	for _, idx := range indices {
		unique[idx] = struct{}{}
	}

	fetchCtx, fetchCancel := context.WithTimeout(ctx, 2*time.Minute)
	var chunkMap map[int][]byte
	var err error
	if rawBlock {
		chunkMap, err = n.fetchChunksFromBlock(fetchCtx, ref, unique)
	} else {
		chunkMap, err = n.fetchChunksSelective(fetchCtx, ref, unique)
	}
	fetchCancel()
	if err != nil {
		return nil, fmt.Errorf("selective fetch: %w", err)
	}

	log.Debug().
		Int("unique_chunks", len(unique)).
		Int("tree_leaves", smt.NumLeaves).
		Msg("selective fetch complete")

	// Build sparse chunks slice — only the 8 needed indices are populated.
	sparseChunks := make([][]byte, smt.NumLeaves)
	for idx, data := range chunkMap {
		sparseChunks[idx] = data
	}

	return n.prover.GenerateProofFromSMT(n.secretKey, randomness, sparseChunks, smt)
}

// fetchChunksSelective fetches specific file chunks by index from IPFS using
// byte-range requests. Each chunk is poi.FileSize bytes at offset index*FileSize.
// Fetches run in parallel; returns a map of index → padded chunk data.
func (n *Node) fetchChunksSelective(ctx context.Context, ref string, indices map[int]struct{}) (map[int][]byte, error) {
	type result struct {
		idx  int
		data []byte
		err  error
	}

	ch := make(chan result, len(indices))
	for idx := range indices {
		go func(i int) {
			offset := int64(i) * int64(poi.FileSize)
			data, err := n.ipfs.CatRangeWithRetry(ctx, ref, offset, int64(poi.FileSize))
			if err != nil {
				ch <- result{i, nil, err}
				return
			}
			// Zero-pad last chunk if shorter (same as SplitIntoChunks)
			if len(data) < poi.FileSize {
				padded := make([]byte, poi.FileSize)
				copy(padded, data)
				data = padded
			}
			ch <- result{i, data, nil}
		}(idx)
	}

	m := make(map[int][]byte, len(indices))
	for range indices {
		r := <-ch
		if r.err != nil {
			return nil, fmt.Errorf("chunk %d: %w", r.idx, r.err)
		}
		m[r.idx] = r.data
	}
	return m, nil
}

// fetchChunksFromBlock fetches a raw IPFS block and extracts the requested
// chunk indices locally. Raw blocks are typically tiny (< 1 KB for directory
// DAG nodes), so fetching the whole block is cheaper than byte-range requests.
func (n *Node) fetchChunksFromBlock(ctx context.Context, ref string, indices map[int]struct{}) (map[int][]byte, error) {
	data, err := n.ipfs.BlockGetWithRetry(ctx, ref)
	if err != nil {
		return nil, fmt.Errorf("block get: %w", err)
	}

	allChunks := merkle.SplitIntoChunks(data, poi.FileSize)
	m := make(map[int][]byte, len(indices))
	for idx := range indices {
		if idx < len(allChunks) {
			m[idx] = allChunks[idx]
		} else {
			// Zero chunk for indices beyond data (padding)
			m[idx] = make([]byte, poi.FileSize)
		}
	}
	return m, nil
}
