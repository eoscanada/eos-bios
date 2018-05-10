#!/bin/bash -e

# `boot_publish_handoff` hook
# $1 = public key
# $2 = private key, proving you haven't kept any access to yourself.

echo "Publish this private key out in the world, to prove you have not kept any access for yourself:"
echo ""
echo "    Public key: $1"
echo "    Private key: $2"
echo ""

#
# You can also publish the private key to IPFS and share the ipfs ref:
# echo "Network boot handoff of ephemeral keys: $1 $2" | ./ipfs add
#

echo "Hit ENTER when done (or skip if you're a single node net)"
read
