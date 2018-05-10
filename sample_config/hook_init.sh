#!/bin/bash -e

# `init` hook
# $1 = operation (either `join`, `boot` or `orchestrate`)

echo "Starting $1 operation"

docker kill nodeos-bios || true
docker rm nodeos-bios || true
