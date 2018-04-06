#!/bin/bash

# do things to connect to bios..
# $1 = p2p_address
# $2 = public key used by BIOS
# $3 = private key used by BIOS
# $4 = genesis_json

echo "Killing running nodes"
killall nodeos

echo "Phasing out any previous blockchain from disk"
rm -rf ~/.eos/blocks ~/.eos/shared_mem

echo "Copying base config"
cp base_config.ini ~/.eos/config.ini

echo "Writing genesis.json"
echo "$4" > ~/.eos/genesis.json

echo "plugin = eosio::producer_plugin" >> ~/.eos/config.ini
echo "producer-name = eoscanada2" >> ~/.eos/config.ini
echo "p2p-peer-address = $1" >> ~/.eos/config.ini
echo "private-key = [\"$NODEOS_PRODUCER_PUBLIC_KEY\",\"$NODEOS_PRODUCER_PRIVATE_KEY\"]" >> ~/.eos/config.ini
echo "resync-blockchain = true" >> ~/.eos/config.ini

echo "Restart your process my friend, then press ENTER"
read
