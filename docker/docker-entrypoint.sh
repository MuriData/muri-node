#!/bin/sh
set -e

CONFIG="${MURID_CONFIG:-/app/murid.toml}"

# Ensure subdirectories exist (bind-mount may overlay Dockerfile-created dirs)
mkdir -p /app/keys /app/data

# Pass through to murid, injecting -config for commands that need it
case "$1" in
    init)
        # Interactive setup — pass -config so output goes to the right place
        shift
        exec murid init -config "$CONFIG" "$@"
        ;;
    run|status|stake|unstake|increase-capacity|decrease-capacity)
        CMD="$1"
        shift
        exec murid "$CMD" -config "$CONFIG" "$@"
        ;;
    wallet|keygen|download-keys|help|-h|--help)
        exec murid "$@"
        ;;
    *)
        exec murid "$@"
        ;;
esac
