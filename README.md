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

## Quick Start

The one-click setup script installs Docker, builds the node, generates keys, downloads prover files, and starts all services (local Avalanche L1 node + IPFS + murid):

```bash
./setup.sh
```

The script will:
1. Install Docker if missing
2. Build the murid container image
3. Generate an EVM wallet and ZK secret key
4. Download PoI prover/verifier keys (~200 MB)
5. Prompt for storage capacity (each GB requires ~26.21 MURI collateral)
6. Write `murid.toml` and start avago + IPFS
7. Wait for you to fund the wallet, then auto-stake and start the daemon

Options:
```bash
./setup.sh --force              # overwrite existing setup
./setup.sh --skip-key-download  # skip prover key download
./setup.sh --min-price 2000     # set min price per chunk/period
```

## Manual Setup

If you prefer to set up step by step:

```bash
# Build
go build -o murid ./cmd/murid/

# Interactive setup (generates keys, writes config)
./murid init

# Download prover keys (~200 MB)
./murid download-keys -out ./keys

# Fund the node address, then stake
./murid stake -capacity-gb 10

# Start the daemon
./murid run
```

### Docker (manual)

```bash
cd docker
docker compose -f docker-compose.yml -f docker-compose.avago.yml run --rm murid init
docker compose -f docker-compose.yml -f docker-compose.avago.yml run --rm murid stake -capacity-gb 10
docker compose -f docker-compose.yml -f docker-compose.avago.yml up -d
```

## Requirements

- **Docker** (installed automatically by `setup.sh`)
- **MURI tokens** for staking collateral (~26.21 MURI per GB)
- The `.key` proving files from `muri-artifacts` (downloaded automatically)
