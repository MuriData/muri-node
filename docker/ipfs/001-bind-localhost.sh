#!/bin/sh
# MuriData IPFS (Kubo) configuration — optimized for storage provider nodes.
# Runs inside the Kubo container before the daemon starts.

# ── Security: bind API/gateway to localhost only ──────────────────────
# Swarm (4001) stays on 0.0.0.0 for peer discovery.
ipfs config Addresses.API /ip4/127.0.0.1/tcp/5001
ipfs config Addresses.Gateway /ip4/127.0.0.1/tcp/8080

# ── Routing: Accelerated DHT + IPNI for fast content discovery ────────
# Custom routing: DHT (WAN + LAN) in parallel with the IPNI delegated
# router (cid.contact / delegated-ipfs.dev) for global content indexing.
# Provides go through both DHT and IPNI so content is discoverable via
# both the DHT crawl and the Network Indexer.
ipfs config --json Routing '{
  "Type": "custom",
  "Routers": {
    "dht-wan": {
      "Type": "dht",
      "Parameters": {
        "Mode": "auto",
        "PublicIPNetwork": true,
        "AcceleratedDHTClient": true
      }
    },
    "dht-lan": {
      "Type": "dht",
      "Parameters": {
        "Mode": "server",
        "PublicIPNetwork": false
      }
    },
    "ipni": {
      "Type": "http",
      "Parameters": {
        "Endpoint": "https://delegated-ipfs.dev"
      }
    },
    "parallel-provide": {
      "Type": "parallel",
      "Parameters": {
        "Routers": [
          { "RouterName": "dht-wan" },
          { "RouterName": "ipni", "IgnoreErrors": true }
        ]
      }
    },
    "parallel-find": {
      "Type": "parallel",
      "Parameters": {
        "Routers": [
          { "RouterName": "dht-lan", "IgnoreErrors": true },
          { "RouterName": "dht-wan" },
          { "RouterName": "ipni", "IgnoreErrors": true, "Timeout": "5s" }
        ]
      }
    }
  },
  "Methods": {
    "provide":        { "RouterName": "parallel-provide" },
    "find-providers": { "RouterName": "parallel-find" },
    "find-peers":     { "RouterName": "parallel-find" },
    "get-ipns":       { "RouterName": "dht-wan" },
    "put-ipns":       { "RouterName": "dht-wan" }
  }
}'

# ── Provide: re-announce all pinned content every 12h ─────────────────
# Default is 22h. Shorter interval keeps provider records fresh in the
# DHT (records expire after ~24h). "pinned" strategy only announces CIDs
# we've explicitly pinned (our stored files), not transient cache blocks.
# (Kubo 0.33+: Reprovider fields migrated to Provide)
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

# ── Experimental: faster provide ──────────────────────────────────────
# OptimisticProvide: speeds up DHT provide by not waiting for all
# closest peers — stores records at the first reachable ones.
# OptimisticProvideJobsPoolSize: parallel provide operations.
ipfs config --json Experimental.OptimisticProvide true
ipfs config --json Experimental.OptimisticProvideJobsPoolSize 60

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
