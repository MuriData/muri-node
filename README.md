# muri-node

**muri-node** is the prover and node infrastructure for MuriData storage providers. 
It interfaces with MuriChain (via `murichain-avalanchego`) to listen for assignments, serve replicas, and generate ZK proofs for challenge rounds.

## Features

- **Standard Storage & PoI**: Tracks standard replica assignments and generates parallelized 8-leaf Proof of Integrity (PoI) Groth16 proofs.
- **Merkle Tree Building**: Builds sparse Merkle trees for the 16KB chunk file splits.
- **On-chain Synchronization**: Packages compressed proof responses and commits them for on-chain submission (via `FileMarket` interactions).

## Responsibilities

1. **Staking & Assignment**: Nodes lock collateral and accept file replica spots via the order system.
2. **Proving**:
   - Downloads/builds the 16KB chunk Merkle trees.
   - Computes witness assignments and outputs Groth16 proofs for the challenge verification window.
   - Pushes finalized `submitProof()` transactions rapidly to remain within the 50-block (~100s) challenge window.

## Running the Node

Requires the `.key` proving files from `muri-artifacts` for standard PoI to operate properly.
