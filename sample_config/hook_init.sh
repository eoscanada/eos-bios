#!/bin/bash

# `init` hook
# no parameters

echo "Doing any preparation before go-live"

mkdir -p ~/.eos

docker kill nodeos-bios

true
