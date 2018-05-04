#!/bin/bash

# `boot_connect_mesh.sh` hook
# $1 p2p-peer-address statements (like the `join_network` hook)
# $2 comma-separated peer address list
#
# This hook is called when you 'eos-bios boot' or are seleted as BIOS
# Boot in an orchestrated launch.
#
# It should connect your boot node

echo "Adding p2p-peer-address'es to config.ini"

echo "$1" >> config.ini


echo "Restarting boot node"

docker restart nodeos-bios

sleep 2
