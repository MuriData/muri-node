#!/usr/bin/env bash
#
# MuriData Node — One-Click Setup
#
# Usage:
#   ./setup.sh              # interactive setup
#   ./setup.sh --force      # overwrite existing setup
#
# What it does:
#   1. Installs Docker + Docker Compose if missing
#   2. Builds the murid container image
#   3. Generates EVM wallet + ZK secret key
#   4. Downloads PoI prover/verifier keys (~200 MB)
#   5. Prompts for storage capacity
#   6. Writes murid.toml
#   7. Starts avago (local Avalanche L1 node) + IPFS (Kubo)
#   8. Waits for funding, then stakes and starts the daemon
#
# Idempotent: skips steps that are already done (existing keys, config, etc.)

set -euo pipefail

# ─── Colors ───────────────────────────────────────────────────────────
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
BOLD='\033[1m'
NC='\033[0m'

info()  { echo -e "${CYAN}[INFO]${NC}  $*"; }
ok()    { echo -e "${GREEN}[OK]${NC}    $*"; }
warn()  { echo -e "${YELLOW}[WARN]${NC}  $*"; }
fail()  { echo -e "${RED}[FAIL]${NC}  $*"; exit 1; }
step()  { echo -e "\n${BOLD}── $* ──${NC}"; }

# ─── Defaults ─────────────────────────────────────────────────────────
MIN_PRICE="1000"
SKIP_KEY_DOWNLOAD=false
FORCE=false

# ─── Parse flags ──────────────────────────────────────────────────────
while [[ $# -gt 0 ]]; do
    case "$1" in
        --min-price)          MIN_PRICE="$2";          shift 2 ;;
        --skip-key-download)  SKIP_KEY_DOWNLOAD=true;  shift ;;
        --force)              FORCE=true;              shift ;;
        -h|--help)
            echo "Usage: $0 [flags]"
            echo ""
            echo "Flags:"
            echo "  --min-price WEI        Min price wei/chunk/period (default: 1000)"
            echo "  --skip-key-download    Don't download PoI prover keys"
            echo "  --force                Overwrite existing config/keys"
            echo ""
            exit 0
            ;;
        *)  fail "Unknown flag: $1 (use --help)" ;;
    esac
done

# ─── Resolve paths ────────────────────────────────────────────────────
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
DOCKER_DIR="${SCRIPT_DIR}/docker"
NODE_DIR="${DOCKER_DIR}/node"
KEYS_DIR="${NODE_DIR}/keys"
DATA_DIR="${NODE_DIR}/data"
CONFIG_FILE="${NODE_DIR}/murid.toml"
COMPOSE_FILES="-f ${DOCKER_DIR}/docker-compose.yml -f ${DOCKER_DIR}/docker-compose.avago.yml"

cd "$SCRIPT_DIR"

echo -e "${BOLD}"
echo "  ╔══════════════════════════════════════╗"
echo "  ║   MuriData Node — One-Click Setup    ║"
echo "  ╚══════════════════════════════════════╝"
echo -e "${NC}"

# ─── Step 1: Install Docker ──────────────────────────────────────────
step "Checking Docker"

install_docker() {
    info "Installing Docker..."
    case "$(uname -s)" in
        Linux)
            if command -v apt-get &>/dev/null; then
                sudo apt-get update -qq
                sudo apt-get install -y -qq ca-certificates curl gnupg
                sudo install -m 0755 -d /etc/apt/keyrings
                curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg 2>/dev/null
                sudo chmod a+r /etc/apt/keyrings/docker.gpg
                DISTRO=$(. /etc/os-release && echo "$ID")
                echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/${DISTRO} $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
                    sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
                sudo apt-get update -qq
                sudo apt-get install -y -qq docker-ce docker-ce-cli containerd.io docker-compose-plugin
            elif command -v dnf &>/dev/null; then
                # Fedora/CentOS/RHEL/Rocky/Alma
                sudo dnf -y install dnf-plugins-core
                DNF_DISTRO=$(. /etc/os-release && echo "$ID")
                case "$DNF_DISTRO" in
                    centos|rhel|rocky|almalinux) DNF_DISTRO="centos" ;;
                    fedora) DNF_DISTRO="fedora" ;;
                    *) DNF_DISTRO="centos" ;;
                esac
                sudo dnf config-manager --add-repo "https://download.docker.com/linux/${DNF_DISTRO}/docker-ce.repo"
                sudo dnf -y install docker-ce docker-ce-cli containerd.io docker-compose-plugin
                sudo systemctl enable --now docker
            else
                fail "Unsupported Linux distro. Install Docker manually: https://docs.docker.com/engine/install/"
            fi
            if ! groups | grep -q docker; then
                sudo usermod -aG docker "$USER" 2>/dev/null || true
                warn "Added $USER to docker group. You may need to log out and back in."
            fi
            ;;
        Darwin)
            if command -v brew &>/dev/null; then
                info "Installing Docker Desktop via Homebrew..."
                brew install --cask docker
                echo ""
                warn "Docker Desktop installed but needs to be started."
                warn "Open Docker Desktop from Applications, then re-run this script."
                exit 1
            else
                fail "Install Docker Desktop from https://docs.docker.com/desktop/install/mac-install/"
            fi
            ;;
        *)
            fail "Unsupported OS. Install Docker manually: https://docs.docker.com/engine/install/"
            ;;
    esac
}

if ! command -v docker &>/dev/null; then
    install_docker
fi

# Verify docker is running
if ! docker info &>/dev/null; then
    if [[ "$(uname -s)" == "Linux" ]]; then
        sudo systemctl start docker 2>/dev/null || true
        sleep 2
    fi
    if ! docker info &>/dev/null; then
        fail "Docker is installed but not running. Start Docker and re-run this script."
    fi
fi
ok "Docker $(docker --version | grep -oE '[0-9]+\.[0-9]+\.[0-9]+' | head -1)"

if docker compose version &>/dev/null; then
    ok "Docker Compose $(docker compose version --short 2>/dev/null || echo 'available')"
else
    fail "Docker Compose plugin not found. Install: https://docs.docker.com/compose/install/"
fi

# ─── Step 2: Create directories ──────────────────────────────────────
step "Preparing directories"

mkdir -p "$KEYS_DIR" "$DATA_DIR"
ok "Node directory: $NODE_DIR"

# ─── Step 3: Build murid image ───────────────────────────────────────
step "Building murid container"

MURID_IMAGE=$(docker compose ${COMPOSE_FILES} images murid -q 2>/dev/null || echo "")
if [[ -n "$MURID_IMAGE" ]] && [[ "$FORCE" == false ]]; then
    ok "murid image already built (use --force to rebuild)"
else
    info "Building murid Docker image (this may take a minute)..."
    docker compose ${COMPOSE_FILES} build murid
    ok "murid image built"
fi

# ─── Helper: run murid command in container ───────────────────────────
murid() {
    docker compose ${COMPOSE_FILES} run --rm -T murid "$@"
}

# ─── Step 4: Generate keys ───────────────────────────────────────────
step "Generating keys"

PRIV_KEY_PATH="/app/keys/node.key"
SECRET_KEY_PATH="/app/keys/secret.key"

# EVM private key
EVM_ADDR=""
if [[ -f "${KEYS_DIR}/node.key" ]] && [[ "$FORCE" == false ]]; then
    ok "EVM key already exists"
else
    info "Generating EVM wallet..."
    EVM_OUTPUT=$(murid wallet new -out "$PRIV_KEY_PATH" -force 2>&1)
    echo "$EVM_OUTPUT" | grep -v "^$" | while IFS= read -r line; do echo "     $line"; done
    EVM_ADDR=$(echo "$EVM_OUTPUT" | grep -oiE '0x[0-9a-fA-F]{40}' | head -1 || echo "")
    ok "EVM wallet generated"
fi

# ZK secret key
if [[ -f "${KEYS_DIR}/secret.key" ]] && [[ "$FORCE" == false ]]; then
    ok "ZK secret key already exists"
else
    info "Generating ZK secret key..."
    ZK_OUTPUT=$(murid keygen -out "$SECRET_KEY_PATH" -force 2>&1)
    echo "$ZK_OUTPUT" | grep -v "^$" | while IFS= read -r line; do echo "     $line"; done
    ok "ZK secret key generated"
fi

# ─── Step 5: Download PoI prover keys ────────────────────────────────
step "PoI prover keys"

if [[ "$SKIP_KEY_DOWNLOAD" == true ]]; then
    info "Skipping key download (--skip-key-download)"
elif [[ -f "${KEYS_DIR}/poi_prover.key" ]] && [[ -f "${KEYS_DIR}/poi_verifier.key" ]] && [[ "$FORCE" == false ]]; then
    PROVER_SIZE=$(wc -c < "${KEYS_DIR}/poi_prover.key" | tr -d ' ')
    if [[ "$PROVER_SIZE" -gt 1000000 ]]; then
        ok "PoI keys already downloaded (prover: $(du -h "${KEYS_DIR}/poi_prover.key" | cut -f1))"
    else
        warn "Prover key looks truncated, re-downloading..."
        murid download-keys -out /app/keys
        ok "PoI keys downloaded"
    fi
else
    info "Downloading PoI prover/verifier keys (~200 MB)..."
    murid download-keys -out /app/keys
    ok "PoI keys downloaded"
fi

# ─── Step 6: Prompt for capacity ─────────────────────────────────────
step "Storage capacity"

# Stake math: 1 GB = 65536 chunks, StakePerChunk = 0.0004 MURI → 26.2144 MURI/GB
echo ""
echo -e "  Each GB of storage capacity requires ${BOLD}~26.21 MURI${NC} as collateral."
echo -e "  Examples:  1 GB = ~26.21 MURI  |  10 GB = ~262.14 MURI  |  100 GB = ~2621.44 MURI"
echo ""

while true; do
    read -p "  Storage capacity in GB [10]: " CAPACITY_INPUT
    CAPACITY_GB="${CAPACITY_INPUT:-10}"
    # Validate: must be a positive number
    if [[ "$CAPACITY_GB" =~ ^[0-9]+\.?[0-9]*$ ]] && [[ $(echo "$CAPACITY_GB > 0" | bc -l 2>/dev/null || echo "0") == "1" ]]; then
        break
    fi
    # Fallback validation without bc
    if [[ "$CAPACITY_GB" =~ ^[1-9][0-9]*$ ]]; then
        break
    fi
    warn "Please enter a positive number."
done

# Calculate chunks and stake requirement
# chunks = GB * 1024 * 1024 * 1024 / 16384 = GB * 65536
# stake_muri = chunks * 0.0004
CHUNKS=$(echo "$CAPACITY_GB * 65536" | bc -l 2>/dev/null | cut -d. -f1 || echo "")
if [[ -z "$CHUNKS" ]]; then
    # Fallback for integer GB without bc
    CHUNKS=$(( ${CAPACITY_GB%.*} * 65536 ))
fi
STAKE_MURI=$(echo "scale=4; $CHUNKS * 0.0004" | bc -l 2>/dev/null || echo "unknown")

ok "Capacity: ${CAPACITY_GB} GB (${CHUNKS} chunks) — requires ${BOLD}${STAKE_MURI} MURI${NC} stake"

# ─── Step 7: Write config ────────────────────────────────────────────
step "Writing config"

if [[ -f "$CONFIG_FILE" ]] && [[ "$FORCE" == false ]]; then
    ok "Config already exists at murid.toml (use --force to overwrite)"
else
    cat > "$CONFIG_FILE" << TOML
# murid — generated by setup.sh (local avago mode)

[chain]
rpc_url = "http://127.0.0.1:9650/ext/bc/inP2vNhcVSABGmq39UHwuB9tDxUUWp3g6gpRwdE6TqtAtAWmu/rpc"
chain_id = 97981
market_address = "0xaab9f94671d6b22eee60509b5c3149e90a78fb54"
staking_address = ""   # auto-resolved from market contract
gas_limit = 2000000
max_gas_price = 100
gas_priority = 2
gas_escalation = 1.25
max_retries = 3
confirmation_blocks = 1
listen_mode = "poll"
ws_url = ""

[ipfs]
api_url = "http://127.0.0.1:5001"
timeout = "30s"
pin_files = true

[node]
private_key_path = "${PRIV_KEY_PATH}"
data_dir = "/app/data"
keys_dir = "/app/keys"
secret_key_path = "${SECRET_KEY_PATH}"

[storage]
max_capacity_gb = ${CAPACITY_GB}
min_price = ${MIN_PRICE}

[auto_execute]
enabled = true
poll_interval = "30s"
max_orders_to_fill = 5

[challenge]
poll_interval = "4s"
safety_margin = 10

[log]
level = "info"
pretty = true
TOML

    ok "Config written to docker/node/murid.toml"
fi

# ─── Step 8: Start infrastructure ────────────────────────────────────
step "Starting services"

info "Starting avago (Avalanche L1 node) + IPFS..."
docker compose ${COMPOSE_FILES} up -d avago ipfs
ok "avago + IPFS started"

# Wait for IPFS health
info "Waiting for IPFS to be healthy..."
IPFS_READY=false
for i in $(seq 1 30); do
    if docker compose ${COMPOSE_FILES} exec -T ipfs ipfs id &>/dev/null 2>&1; then
        IPFS_READY=true
        break
    fi
    sleep 2
done
if [[ "$IPFS_READY" == true ]]; then
    ok "IPFS is healthy"
else
    warn "IPFS not yet healthy — it may still be initializing."
fi

# ─── Step 9: Wait for funding, then stake ────────────────────────────
step "Fund & Stake"

# Resolve EVM address if not captured during keygen
if [[ -z "$EVM_ADDR" ]] && [[ -f "${KEYS_DIR}/node.key" ]]; then
    EVM_ADDR=$(murid wallet balance -config /app/murid.toml 2>&1 | grep -oiE '0x[0-9a-fA-F]{40}' | head -1 || echo "")
fi

echo ""
echo -e "  To stake ${BOLD}${CAPACITY_GB} GB${NC}, your node wallet needs at least ${BOLD}${STAKE_MURI} MURI${NC}."
echo ""
if [[ -n "$EVM_ADDR" ]]; then
    echo -e "  ${BOLD}Send MURI to: ${EVM_ADDR}${NC}"
else
    echo -e "  Run ${CYAN}murid wallet balance${NC} to see your address."
fi
echo ""
echo -e "  Waiting for sufficient balance... (press Ctrl+C to skip and do it later)"
echo ""

# Poll balance until sufficient
FUNDED=false
while true; do
    BALANCE_OUTPUT=$(murid wallet balance -config /app/murid.toml 2>&1 || echo "")
    # Extract balance line like "Balance: 123.456789 MURI"
    BALANCE_LINE=$(echo "$BALANCE_OUTPUT" | grep -i "balance" || echo "")

    if [[ -n "$BALANCE_LINE" ]]; then
        # Extract the numeric balance
        BALANCE_VAL=$(echo "$BALANCE_LINE" | grep -oE '[0-9]+\.[0-9]+' | head -1 || echo "0")
        if [[ -n "$BALANCE_VAL" ]]; then
            # Compare balance >= stake_muri
            ENOUGH=$(echo "$BALANCE_VAL >= $STAKE_MURI" | bc -l 2>/dev/null || echo "0")
            if [[ "$ENOUGH" == "1" ]]; then
                ok "Balance: ${BALANCE_VAL} MURI — sufficient for staking"
                FUNDED=true
                break
            else
                echo -ne "\r  Balance: ${BALANCE_VAL} MURI (need ${STAKE_MURI} MURI)... waiting    "
            fi
        fi
    else
        echo -ne "\r  Checking balance... (avago may still be syncing)    "
    fi
    sleep 10
done

# Stake
if [[ "$FUNDED" == true ]]; then
    echo ""
    info "Registering node on-chain with ${CAPACITY_GB} GB..."
    STAKE_OUTPUT=$(murid stake -capacity-gb "$CAPACITY_GB" 2>&1) || {
        echo "$STAKE_OUTPUT"
        fail "Staking failed. You can retry manually: docker compose ${COMPOSE_FILES} run --rm murid stake -capacity-gb ${CAPACITY_GB}"
    }
    echo "$STAKE_OUTPUT" | grep -v "^$" | while IFS= read -r line; do echo "     $line"; done
    ok "Node staked successfully!"
fi

# ─── Step 10: Start murid daemon ─────────────────────────────────────
step "Starting murid daemon"

docker compose ${COMPOSE_FILES} up -d murid
ok "murid daemon started"

# ─── Done ─────────────────────────────────────────────────────────────
step "Setup complete"

COMPOSE_CMD="cd docker && docker compose -f docker-compose.yml -f docker-compose.avago.yml"

echo ""
echo -e "  ${GREEN}Your MuriData storage node is now running!${NC}"
echo ""
echo -e "  ${CYAN}Config:${NC}      docker/node/murid.toml"
echo -e "  ${CYAN}Keys dir:${NC}    docker/node/keys/"
echo -e "  ${CYAN}Data dir:${NC}    docker/node/data/"
echo -e "  ${CYAN}Capacity:${NC}    ${CAPACITY_GB} GB (${CHUNKS} chunks)"
echo -e "  ${CYAN}RPC:${NC}         Local avago (127.0.0.1:9650)"
echo -e "  ${CYAN}Network:${NC}     Chain ID 97981"
if [[ -n "$EVM_ADDR" ]]; then
    echo -e "  ${CYAN}Node addr:${NC}   ${EVM_ADDR}"
fi
echo ""
echo -e "  Useful commands:"
echo -e "     ${CYAN}${COMPOSE_CMD} logs -f murid${NC}              # follow logs"
echo -e "     ${CYAN}${COMPOSE_CMD} run --rm murid status${NC}      # node status"
echo -e "     ${CYAN}${COMPOSE_CMD} run --rm murid wallet balance${NC}"
echo -e "     ${CYAN}${COMPOSE_CMD} down${NC}                       # stop everything"
echo -e "     ${CYAN}${COMPOSE_CMD} up -d${NC}                      # restart everything"
echo ""
