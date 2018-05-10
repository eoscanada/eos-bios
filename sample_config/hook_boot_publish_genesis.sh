#!/bin/bash -e

# `boot_publish_genesis` hook
# $1 = base64 encoded genesis json content
# $2 = raw genesis JSON content

echo "Please publish this genesis data to everyone:"
echo ""
echo "    $1"
echo ""

#
# You can also publish it through ipfs:
#
#echo $2 | ./ipfs add
#
# and publish the resulting URL
#

echo "Hit ENTER when done (or skip if launching a local node only)"
read
