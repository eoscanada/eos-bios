#!/bin/bash

echo "Compile the wast file"
/home/abourget/build/eos/build/wat2wasm /home/abourget/build/eos/build/contracts/eosio.system/eosio.system.wast -o /home/abourget/build/eos/build/contracts/eosio.system/eosio.system.wasm
