#!/bin/bash

# `start_bios_boot` hook
# $1 genesis JSON
# $2 ephemeral public key
# $3 ephemeral private key
#
# This process must not BLOCK.

echo "Killing running nodes"
killall nodeos

echo "Phasing out any previous blockchain from disk"
mkdir -p ~/.eos
rm -rf ~/.eos/blocks ~/.eos/shared_mem

echo "Copying base config"
# Your base_config.ini shouldn't contain any `producer-name` nor `private-key` nor `enable-stale-production` statements.
cp base_config.ini ~/.eos/config.ini

echo "Writing genesis.json"
echo $1 > ~/.eos/genesis.json

echo "plugin = eosio::producer_plugin" >> ~/.eos/config.ini
echo "producer-name = eosio" >> ~/.eos/config.ini
echo "enable-stale-production = true" >> ~/.eos/config.ini
echo "private-key = [\"$2\",\"$3\"]" >> ~/.eos/config.ini

# Replace this by some automated command to restart the node.
echo "CONFIGURATION DONE: Re/start nodeos, and press ENTER"
# You could run, on another tab:
#    nodeos --data-dir $HOME/.eos --genesis-json $HOME/.eos/genesis.json --config-dir $HOME/.eos
read
