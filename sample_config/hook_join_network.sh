#!/bin/bash -e

# `join_network` hook:
# $1 = genesis_json
# $2 = p2p_address_statements like "p2p-peer-address = 1.2.3.4\np2p-peer-address=2.3.4.5"
# $3 = p2p_addresses to connect to, split by comma
# $4 = producer-name statements, like: "producer-name = hello\nproducer-name = hello.a"
#      You will have many only when joining a net with less than 21 producers.
# $5 = producer-name you should handle, split by comma


# WARN: this is SAMPLE keys configuration to get your keys into your config.
#       You'll want to adapt that to your infrastructure, `cat` it from a file,
#       use some secrets management software or whatnot.
#
#       They need to reflect your `target_initial_authority`
#       strucuture in your `my_discovery_file.yaml`.
#
PUBKEY=EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV
PRIVKEY=5KQwrPbwdL6PhXujxW37FSSQZ1JiwsST4cqQzDeyXtP79zkvFD3


echo "Killing running nodes"
killall nodeos || true

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
       eosio/eos:dawn-v4.0.0 \
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
