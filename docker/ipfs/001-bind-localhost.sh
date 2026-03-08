#!/bin/sh
# Bind API and gateway to localhost only.
# Swarm (4001) stays on 0.0.0.0 for peer discovery.
# Runs after the Docker image sets both to 0.0.0.0.
ipfs config Addresses.API /ip4/127.0.0.1/tcp/5001
ipfs config Addresses.Gateway /ip4/127.0.0.1/tcp/8080
