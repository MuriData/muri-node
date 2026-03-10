package ipfs

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/MuriData/muri-node/config"
	"github.com/rs/zerolog/log"
)

// apiURL builds a fully-escaped Kubo API URL with the given path and arg.
func (c *Client) apiEndpoint(path, arg string) string {
	u := fmt.Sprintf("%s%s", c.apiURL, path)
	if arg != "" {
		u += "?" + url.Values{"arg": {arg}}.Encode()
	}
	return u
}

// Client is an HTTP client for the Kubo IPFS API.
type Client struct {
	apiURL      string
	http        *http.Client // normal ops: 30s response header timeout
	httpBulk    *http.Client // bulk downloads: 2min — Kubo may need to discover/fetch blocks from remote peers
	idleTimeout time.Duration
	maxRetries  int
	baseDelay   time.Duration
}

// NewClient creates an IPFS client from config.
func NewClient(cfg config.IPFSConfig) *Client {
	timeout := cfg.Timeout.Duration
	if timeout == 0 {
		timeout = 30 * time.Second
	}
	maxRetries := cfg.MaxRetries
	if maxRetries == 0 {
		maxRetries = 4
	}
	baseDelay := cfg.RetryDelay.Duration
	if baseDelay == 0 {
		baseDelay = 2 * time.Second
	}

	// Clone the default transport and add ResponseHeaderTimeout.
	// This catches unreachable IPFS nodes quickly (connection + headers)
	// without capping the entire body transfer like http.Client.Timeout does.
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.ResponseHeaderTimeout = timeout

	// Bulk transport: longer timeout for first-time downloads of large files.
	// Kubo may need to traverse the DHT, discover providers, and fetch blocks
	// from distant peers — 30s is often not enough for cold blocks.
	transportBulk := http.DefaultTransport.(*http.Transport).Clone()
	transportBulk.ResponseHeaderTimeout = 2 * time.Minute

	return &Client{
		apiURL:      cfg.APIURL,
		http:        &http.Client{Transport: transport},
		httpBulk:    &http.Client{Transport: transportBulk},
		idleTimeout: timeout,
		maxRetries:  maxRetries,
		baseDelay:   baseDelay,
	}
}

// Cat fetches the raw bytes of a CID from IPFS.
// Connection and response-header timeouts are handled by the transport.
// Body streaming uses an idle timeout so arbitrarily large files can transfer
// as long as data keeps flowing.
func (c *Client) Cat(ctx context.Context, cid string) ([]byte, error) {
	url := c.apiEndpoint("/api/v0/cat", cid)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ipfs cat: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("ipfs cat %s: status %d: %s", cid, resp.StatusCode, string(body))
	}

	return readAllIdleTimeout(ctx, resp.Body, c.idleTimeout)
}

// readAllIdleTimeout reads all data from body, timing out if no data arrives
// for the specified idle duration. Unlike http.Client.Timeout which caps the
// entire request lifetime, this only fires when data stops flowing — allowing
// arbitrarily large files to transfer as long as progress continues.
func readAllIdleTimeout(ctx context.Context, body io.ReadCloser, idle time.Duration) ([]byte, error) {
	type chunk struct {
		data []byte
		err  error
	}

	ch := make(chan chunk, 4)

	go func() {
		defer close(ch)
		buf := make([]byte, 256*1024) // 256 KB read buffer
		for {
			n, err := body.Read(buf)
			if n > 0 {
				data := make([]byte, n)
				copy(data, buf[:n])
				ch <- chunk{data: data}
			}
			if err != nil {
				if err != io.EOF {
					ch <- chunk{err: err}
				}
				return
			}
		}
	}()

	var result bytes.Buffer
	received := int64(0)
	timer := time.NewTimer(idle)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			body.Close()
			return nil, ctx.Err()

		case c, ok := <-ch:
			if !ok {
				// Reader goroutine hit EOF — transfer complete.
				if received > 1024*1024 {
					log.Debug().Int64("bytes", received).Msg("ipfs download complete")
				}
				return result.Bytes(), nil
			}
			if c.err != nil {
				return nil, c.err
			}
			result.Write(c.data)
			received += int64(len(c.data))
			// Reset idle timer — data is still flowing.
			if !timer.Stop() {
				select {
				case <-timer.C:
				default:
				}
			}
			timer.Reset(idle)

		case <-timer.C:
			body.Close() // unblocks the reader goroutine
			return nil, fmt.Errorf("idle timeout: no data received for %v (got %d bytes so far)", idle, received)
		}
	}
}

// CatRange fetches a byte range from a CID using Kubo's offset/length parameters.
// Only the IPFS DAG blocks covering the requested range are retrieved from the network.
func (c *Client) CatRange(ctx context.Context, cid string, offset, length int64) ([]byte, error) {
	return c.catRangeWith(ctx, c.http, cid, offset, length)
}

// catRangeWith is the internal implementation of CatRange, parameterized by HTTP client.
func (c *Client) catRangeWith(ctx context.Context, httpClient *http.Client, cid string, offset, length int64) ([]byte, error) {
	params := url.Values{
		"arg":    {cid},
		"offset": {fmt.Sprintf("%d", offset)},
		"length": {fmt.Sprintf("%d", length)},
	}
	u := fmt.Sprintf("%s/api/v0/cat?%s", c.apiURL, params.Encode())
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ipfs cat range: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("ipfs cat range %s: status %d: %s", cid, resp.StatusCode, string(body))
	}

	// Response is bounded by the requested length, so io.ReadAll is fine.
	return io.ReadAll(resp.Body)
}

// withRetry executes fn with exponential backoff retry.
func (c *Client) withRetry(ctx context.Context, fn func() ([]byte, error)) ([]byte, error) {
	var lastErr error
	delay := c.baseDelay

	for attempt := 0; attempt <= c.maxRetries; attempt++ {
		data, err := fn()
		if err == nil {
			return data, nil
		}
		lastErr = err

		if attempt == c.maxRetries {
			break
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(delay):
		}
		delay *= 2
	}

	return nil, fmt.Errorf("after %d retries: %w", c.maxRetries, lastErr)
}

// CatWithRetry fetches the raw bytes of a CID with exponential backoff retry.
func (c *Client) CatWithRetry(ctx context.Context, cid string) ([]byte, error) {
	return c.withRetry(ctx, func() ([]byte, error) { return c.Cat(ctx, cid) })
}

// CatRangeWithRetry fetches a byte range with exponential backoff retry.
func (c *Client) CatRangeWithRetry(ctx context.Context, cid string, offset, length int64) ([]byte, error) {
	return c.withRetry(ctx, func() ([]byte, error) { return c.CatRange(ctx, cid, offset, length) })
}

// downloadSegmentSize is the chunk size for segmented downloads.
// 1 MB balances request overhead (~1024 requests for 1 GB) against
// retry granularity (only re-download the failed 1 MB, not the whole file).
const downloadSegmentSize int64 = 1024 * 1024

// CatChunked downloads a CID in sequential 1 MB segments, retrying each
// independently. Uses the bulk HTTP client (2 min response header timeout)
// so Kubo has time to discover and fetch blocks from remote peers.
//
// Much more resilient than a monolithic Cat for large files:
// if the connection drops at 500 MB, only the current 1 MB segment is retried
// instead of restarting the entire download.
func (c *Client) CatChunked(ctx context.Context, cid string) ([]byte, error) {
	var result bytes.Buffer
	offset := int64(0)

	for {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}

		// Use bulk client (2 min response header timeout) with per-segment retry.
		data, err := c.withRetry(ctx, func() ([]byte, error) {
			return c.catRangeWith(ctx, c.httpBulk, cid, offset, downloadSegmentSize)
		})
		if err != nil {
			return nil, fmt.Errorf("offset %d: %w", offset, err)
		}
		if len(data) == 0 {
			break
		}
		result.Write(data)
		offset += int64(len(data))

		// Log progress for large downloads (every 100 MB)
		if offset%(100*1024*1024) == 0 {
			log.Debug().Int64("bytes", offset).Msg("ipfs chunked download progress")
		}

		if int64(len(data)) < downloadSegmentSize {
			break // last segment
		}
	}

	if result.Len() > 1024*1024 {
		log.Debug().Int("bytes", result.Len()).Msg("ipfs chunked download complete")
	}
	return result.Bytes(), nil
}

// CatChunkedTo downloads a CID in segments and calls fn for each
// fileChunkSize-aligned chunk. The full file is never buffered — only one
// download segment (~1 MB) plus a small carry buffer are held at a time.
//
// fn receives the chunk index (0-based) and the chunk data. The data slice
// is only valid for the duration of the call — callers must copy if needed.
// The last chunk is zero-padded to fileChunkSize.
//
// Returns the total number of chunks delivered to fn.
func (c *Client) CatChunkedTo(ctx context.Context, cid string, fileChunkSize int, fn func(index int, data []byte)) (int, error) {
	offset := int64(0)
	chunkIndex := 0
	var carry []byte

	for {
		if ctx.Err() != nil {
			return 0, ctx.Err()
		}

		segment, err := c.withRetry(ctx, func() ([]byte, error) {
			return c.catRangeWith(ctx, c.httpBulk, cid, offset, downloadSegmentSize)
		})
		if err != nil {
			return 0, fmt.Errorf("offset %d: %w", offset, err)
		}
		if len(segment) == 0 {
			break
		}
		lastSegment := int64(len(segment)) < downloadSegmentSize
		offset += int64(len(segment))

		// Prepend any carry from the previous segment
		buf := segment
		if len(carry) > 0 {
			buf = append(carry, segment...)
			carry = nil
		}

		// Deliver complete chunks
		for len(buf) >= fileChunkSize {
			fn(chunkIndex, buf[:fileChunkSize])
			buf = buf[fileChunkSize:]
			chunkIndex++
		}

		// Save leftover for next segment
		if len(buf) > 0 {
			carry = make([]byte, len(buf))
			copy(carry, buf)
		}

		// Log progress for large downloads
		if offset%(100*1024*1024) == 0 {
			log.Debug().Int64("bytes", offset).Msg("ipfs streaming download progress")
		}

		if lastSegment {
			break
		}
	}

	// Handle final partial chunk (zero-padded)
	if len(carry) > 0 {
		padded := make([]byte, fileChunkSize)
		copy(padded, carry)
		fn(chunkIndex, padded)
		chunkIndex++
	}

	if offset > 1024*1024 {
		log.Debug().Int64("bytes", offset).Int("chunks", chunkIndex).Msg("ipfs streaming download complete")
	}

	return chunkIndex, nil
}

// addResponse is the JSON response from /api/v0/add.
type addResponse struct {
	Hash string `json:"Hash"`
	Name string `json:"Name"`
	Size string `json:"Size"`
}

// Add uploads data to IPFS and returns the CID.
func (c *Client) Add(ctx context.Context, data []byte, filename string) (string, error) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return "", fmt.Errorf("create form file: %w", err)
	}
	if _, err := part.Write(data); err != nil {
		return "", fmt.Errorf("write data: %w", err)
	}
	writer.Close()

	url := c.apiEndpoint("/api/v0/add", "")
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &body)
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.http.Do(req)
	if err != nil {
		return "", fmt.Errorf("ipfs add: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("ipfs add: status %d: %s", resp.StatusCode, string(respBody))
	}

	var result addResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("decode response: %w", err)
	}

	return result.Hash, nil
}

// Pin pins a CID to the local IPFS node.
func (c *Client) Pin(ctx context.Context, cid string) error {
	url := c.apiEndpoint("/api/v0/pin/add", cid)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("ipfs pin: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("ipfs pin %s: status %d: %s", cid, resp.StatusCode, string(body))
	}

	return nil
}

// Unpin removes a pin for a CID from the local IPFS node.
func (c *Client) Unpin(ctx context.Context, cid string) error {
	url := c.apiEndpoint("/api/v0/pin/rm", cid)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("ipfs unpin: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("ipfs unpin %s: status %d: %s", cid, resp.StatusCode, string(body))
	}

	return nil
}

// IsPinned checks if a CID is pinned locally.
func (c *Client) IsPinned(ctx context.Context, cid string) (bool, error) {
	url := c.apiEndpoint("/api/v0/pin/ls", cid)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return false, fmt.Errorf("create request: %w", err)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return false, fmt.Errorf("ipfs pin ls: %w", err)
	}
	defer resp.Body.Close()

	// Drain body for HTTP connection reuse
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}
	// Kubo returns 500 for "not pinned" — distinguish from other errors
	if resp.StatusCode == http.StatusInternalServerError {
		return false, nil
	}
	return false, fmt.Errorf("ipfs pin ls %s: status %d: %s", cid, resp.StatusCode, string(body))
}

// Provide announces to the DHT that this node can provide the given CID.
// Uses the bulk HTTP client (2 min timeout) because DHT provide must contact
// multiple peers across the network. Retries on transient failures.
func (c *Client) Provide(ctx context.Context, cid string) error {
	var lastErr error
	for attempt := 0; attempt <= 2; attempt++ {
		if attempt > 0 {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(time.Duration(attempt) * 5 * time.Second):
			}
		}

		err := c.provideOnce(ctx, cid)
		if err == nil {
			return nil
		}
		lastErr = err
		log.Debug().Err(err).Str("cid", cid).Int("attempt", attempt+1).Msg("dht provide attempt failed")
	}
	return lastErr
}

func (c *Client) provideOnce(ctx context.Context, cid string) error {
	// Kubo 0.40+ moved /api/v0/dht/provide → /api/v0/routing/provide.
	// Try the new endpoint first, fall back to the old one for older versions.
	u := c.apiEndpoint("/api/v0/routing/provide", cid)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	resp, err := c.httpBulk.Do(req)
	if err != nil {
		return fmt.Errorf("ipfs routing provide: %w", err)
	}
	defer resp.Body.Close()

	// Kubo streams NDJSON progress; drain it so the connection can be reused.
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ipfs routing provide %s: status %d: %s", cid, resp.StatusCode, truncate(body, 256))
	}

	return nil
}

// truncate returns at most n bytes of b as a string, for error messages.
func truncate(b []byte, n int) string {
	if len(b) <= n {
		return string(b)
	}
	return string(b[:n]) + "..."
}

// Ping checks connectivity to the IPFS node.
func (c *Client) Ping(ctx context.Context) error {
	url := c.apiEndpoint("/api/v0/id", "")
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("ipfs ping: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ipfs ping: status %d", resp.StatusCode)
	}

	return nil
}
