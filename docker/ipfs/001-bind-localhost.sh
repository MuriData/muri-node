#!/bin/sh
# Bind API and gateway to localhost only.
# Swarm (4001) stays on 0.0.0.0 for peer discovery.
# Runs after the Docker image sets both to 0.0.0.0.
ipfs config Addresses.API /ip4/127.0.0.1/tcp/5001
ipfs config Addresses.Gateway /ip4/127.0.0.1/tcp/8080

# Use auto DHT mode: server when publicly reachable, client when behind NAT.
# The "server" profile forces dhtserver which fails if peers can't connect inbound.
ipfs config Routing.Type dht

# Enable relay client so the node is reachable via relay when behind NAT.
# The "server" profile disables this.
ipfs config --bool Swarm.RelayClient.Enabled true
