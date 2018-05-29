#!/bin/bash -e

# THIS IS A SAMPLE FILE. PLEASE TWEAK FOR YOUR INFRASTRUCTURE.

# `boot_node` hook
# $1 genesis JSON
# $2 ephemeral public key
# $3 ephemeral private key
# $4 = p2p_address_statements like "p2p-peer-address = 1.2.3.4\np2p-peer-address=2.3.4.5"
# $5 = p2p_addresses to connect to, split by comma
#
# This process must not BLOCK.

echo "Copying base config"
cp base_config.ini config.ini

echo "Writing genesis.json"
echo $1 > genesis.json

echo "producer-name = eosio" >> config.ini
echo "enable-stale-production = true" >> config.ini
echo "private-key = [\"$2\",\"$3\"]" >> config.ini
echo "$4" >> config.ini

echo "Removing old nodeos data (you might be asked for your sudo password)..."
sudo rm -rf /tmp/nodeos-data

echo "Running 'nodeos' through Docker."
docker run -ti --rm --detach --name nodeos-bios \
       -v `pwd`:/etc/nodeos -v /tmp/nodeos-data:/data \
       -p 8888:8888 -p 9876:9876 \
       eoscanada/eos:dawn-v4.2.0 \
       /opt/eosio/bin/nodeos --data-dir=/data \
                             --config-dir=/etc/nodeos \
                             --genesis-json=/etc/nodeos/genesis.json \
                             --p2p-listen-endpoint=127.0.0.1:65432 \
                             --max-transaction-time=5000

#~/build/eos/build/programs/nodeos/nodeos --data-dir /tmp/nodeos-data --genesis-json `pwd`/genesis.json --max-transaction-time=5000 --p2p-listen-endpoint=127.0.0.1:65432 --config-dir `pwd` &

# Reasons for options:
#
# --genesis-json to initialize the chain, can only be put the FIRST boot, take out after.
# --p2p-listen-endpoint is a quick way to make sure your node is NOT reachable during the boot.
#                       don't open or forward traffic to that point.
# --max-transaction-time is to avoid timeouts when doing the initial actions insertion.
#
# All three options can be removed when you're ready to mesh with the network.

echo ""
echo "   View logs with: docker logs -f nodeos-bios"
echo ""

echo "Waiting 3 secs for nodeos to launch through Docker"
sleep 3

echo "Hit ENTER to continue"
read
