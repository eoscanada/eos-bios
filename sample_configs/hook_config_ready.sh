#!/bin/bash

# config_ready hook dispatched

echo "Killing running nodes"
killall nodeos

echo "Phasing out any previous blockchain from disk"
rm -rf ~/.eos/blocks ~/.eos/shared_mem

echo "Copying base config"
cp base_config.ini ~/.eos/config.ini

echo "Writing genesis.json"
echo $1 > ~/.eos/genesis.json

if [ $4 == "true" ]; then
    echo "plugin = eosio::producer_plugin" >> ~/.eos/config.ini
    echo "producer-name = $5" >> ~/.eos/config.ini

    # sed -i -e "s/^#producer-name\(.*\)/producer-name = $5/g" ~/.eos/config.ini
    if [ $5 == "eosio" ]; then
        echo "enable-stale-production = true" >> ~/.eos/config.ini
        echo "private-key = [\"$2\",\"$3\"]" >> ~/.eos/config.ini
    fi
fi

echo "Restart your process my friend, then press ENTER"
read
