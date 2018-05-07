#!/bin/bash

# `join_network` hook:
# $1 = genesis_json
# $2 = p2p_address_statements like "p2p-peer-address = 1.2.3.4\np2p-peer-address=2.3.4.5"
# $3 = p2p_addresses to connect to, split by comma
# $4 = producer-name statements, like: "producer-name = hello\nproducer-name = hello.a"
#      You will have many only when joining a net with less than 21 producers.
# $5 = producer-name you should handle, split by comma

PUBKEY=EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV
PRIVKEY=`cat privkey-GDW5CV.key`

echo "Killing running nodes"
killall nodeos

echo "Removing old nodeos data (you might be asked for your sudo password)..."
sudo rm -rf /tmp/nodeos-data

echo "Writing genesis.json"
echo $1 > genesis.json

# Your base_config.ini shouldn't contain any `producer-name` nor `private-key`
# nor `enable-stale-production` statements.
echo "Copying base config"
cp base_config.ini config.ini
echo "$2" >> config.ini
echo "$4" >> config.ini
echo "private-key = [\"$PUBKEY\",\"$PRIVKEY\"]" >> config.ini

echo "Running 'nodeos' through Docker."
docker run -ti --detach --name nodeos-bios \
       -v `pwd`:/etc/nodeos -v /tmp/nodeos-data:/data \
       -p 8888:8888 -p 9876:9876 \
       eosio/eos:dawn3x \
       /opt/eosio/bin/nodeos --data-dir=/data \
                             --genesis-json=/etc/nodeos/genesis.json \
                             --config-dir=/etc/nodeos

echo ""
echo "   View logs with: docker logs -f nodeos-bios"
echo ""

echo "Waiting 3 secs for nodeos to launch through Docker"
sleep 3

echo "Hit ENTER to continue"
read
