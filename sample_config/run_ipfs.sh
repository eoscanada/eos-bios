#!/bin/bash

export IPFS_PATH=./.ipfs-data

if [ ! -f ./ipfs ]; then
    echo "ipfs binary missing in `pwd`"
    echo "Please install it in this directory and try again."
    echo "To do so, download 'go-ipfs' from https://dist.ipfs.io/#go-ipfs and extract the 'ipfs' binary here."
    exit 1
fi

# IF YOU ALTER THESE, EXPORT THOSE VARS FOR
# eos-bios TO PICK UP AND COMMUNICATE WITH
# THE RIGHT IPFS INSTANCE:
#
#./ipfs config Addresses.Gateway /ip4/127.0.0.1/tcp/8081  # instead of 8080
#./ipfs config Addresses.API /ip4/127.0.0.1/tcp/5002      # instead of 5001, you'll need to use --ipfs-api-address /ip4/
#
# In the terminal where you call `eos-bios` and want to link to this IPFS:
#
#     export EOS_BIOS_IPFS_LOCAL_GATEWAY_ADDRESS=http://127.0.0.1:8081
#     export EOS_BIOS_IPFS_API_ADDRESS=/ip4/127.0.0.1/tcp/5002
#

./ipfs daemon --init --enable-namesys-pubsub
