#!/bin/sh
# MuriData IPFS (Kubo) configuration — optimized for storage provider nodes.
# Runs inside the Kubo container before the daemon starts.

# ── Security: bind API/gateway to localhost only ──────────────────────
# Swarm (4001) stays on 0.0.0.0 for peer discovery.
ipfs config Addresses.API /ip4/127.0.0.1/tcp/5001
ipfs config Addresses.Gateway /ip4/127.0.0.1/tcp/8080

# ── Routing: Accelerated DHT + IPNI for fast content discovery ────────
# "auto" mode handles DHT (WAN + LAN) automatically — starts as client,
# switches to server when publicly reachable. AcceleratedDHTClient scans
# the full DHT every hour for ~10x faster lookups (useful for storage
# providers that need fast challenge responses). Kubo 0.40 fixed the
# interaction with DHT sweep so provide waits for initial FullRT crawl.
# DelegatedRouters adds IPNI (cid.contact) for global content indexing.
# (Kubo 0.35+: DelegatedRouters replaces verbose Routing.Type=custom)
ipfs config --json Routing.Type '"auto"'
ipfs config --json Routing.AcceleratedDHTClient true
ipfs config --json Routing.DelegatedRouters '["https://delegated-ipfs.dev"]'

# ── Provide: re-announce pinned content via DHT sweep ─────────────────
# Kubo 0.39+ enables DHT sweep by default (Provide.DHT.SweepEnabled=true),
# which spreads reprovide operations smoothly across the interval instead
# of batching them. "pinned" strategy only announces CIDs we've explicitly
# pinned (our stored files), not transient cache blocks.
# Interval of 12h (default 22h) keeps provider records fresh in the DHT
# (records expire after ~24h).
ipfs config --json Provide.Strategy '"pinned"'
ipfs config --json Provide.DHT.Interval '"12h"'

# ── Swarm: connection manager tuning ──────────────────────────────────
# Storage nodes benefit from more connections for better block exchange
# and faster content retrieval. Higher limits than default (96/32).
ipfs config --json Swarm.ConnMgr '{
  "Type": "basic",
  "HighWater": 400,
  "LowWater": 100,
  "GracePeriod": "30s"
}'

# ── NAT traversal & relay ────────────────────────────────────────────
# RelayClient: discover and use relay nodes when behind NAT.
# RelayService: also serve as a relay for other nodes if publicly reachable.
# Hole punching: enables NAT hole-punching via DCUtR for direct connections.
ipfs config --bool Swarm.RelayClient.Enabled true
ipfs config --bool Swarm.RelayService.Enabled true
ipfs config --bool Swarm.EnableHolePunching true

# ── Garbage collection: reclaim space from unpinned transient blocks ──
# GC runs automatically when the repo exceeds StorageMax. Only unpinned
# blocks are removed — all pinned order CIDs are safe. GCPeriod controls
# how often the check runs (not how long GC takes).
# StorageMax should be well above the node's committed storage capacity
# to leave headroom for transient Bitswap blocks and DAG traversal cache.
ipfs config --json Datastore.StorageMax '"50GB"'
ipfs config --json Datastore.GCPeriod '"1h"'

# ── Resource manager: raise limits for storage workload ───────────────
# Default limits can throttle under heavy block exchange. Raise the
# system-wide limits to accommodate many concurrent transfers.
ipfs config --json Swarm.ResourceMgr.Enabled true
ipfs config --json Swarm.ResourceMgr.MaxMemory '"2GB"'

# ── Peering: persistent connections to other MuriData nodes ───────────
# Static peering bypasses DHT for known nodes — blocks transfer directly.
# Add known MuriData bootstrap/relay peers here as the network grows.
# Example (uncomment and fill with real multiaddrs):
# ipfs config --json Peering.Peers '[
#   { "ID": "12D3KooW...", "Addrs": ["/ip4/X.X.X.X/tcp/4001"] }
# ]'
