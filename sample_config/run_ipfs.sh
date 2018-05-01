#!/bin/bash

export IPFS_PATH=./.ipfs-data

if [ ! -f ./ipfs ]; then
    echo "ipfs binary missing in `pwd`"
    echo "Please install it in this directory and try again."
    echo "To do so, download 'go-ipfs' from https://dist.ipfs.io/#go-ipfs and extract the 'ipfs' binary here."
    exit 1
fi

#./ipfs config Addresses.Gateway /ip4/127.0.0.1/tcp/8081  # instead of 8080
#./ipfs config Addresses.API /ip4/127.0.0.1/tcp/5002      # instead of 5001, you'll need to use --ipfs-api-address /ip4/

./ipfs daemon --init
