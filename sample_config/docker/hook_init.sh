#!/bin/bash

# `init` hook
# no parameters

echo "Doing any preparation before go-live"

docker kill nodeos-bios

true
