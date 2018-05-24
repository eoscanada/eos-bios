# EOS.IO Software-based blockchain boot orchestration tool

[点击查看中文](./README.join-cn.md)

## 1. Clone the [eos-bios repo](https://github.com/eoscanada/eos-bios)
## 2. Download the latest [eos-bios release](https://github.com/eoscanada/eos-bios/releases)

OR build from source:
        
    go get -u -v github.com/eoscanada/eos-bios/eos-bios

## 3. Ask someone to invite you to the seed network
Provide them with your desired 12 character account name and public key. They will run 
        
    eos-bios invite YOUR_ACCOUNT_NAME YOUR_PUBKEY

## 4. Update your disco file
Copy the `sample_config` folder to a folder named `stageX` where X is the stage number we are launching. Edit `my_discovery_file.yaml`:
* `seed_network_account_name`, make it match what you provided for an invite.
* `seed_network_http_address`, this should be the address of the seed network you want to orchestrate from.
* `seed_network_peers` this one warrants [its own section](#network-peers).
* `target_http_address` is the address for `eos-bios` to reach your node
* `target_p2p_address` is a publicly reachable address advertised to mesh the network
* `target_account_name`, `target_appointed_block_producer_signing_key`, `target_initial_authority`: the values you want to use on the target network
* `target_contents` are all the pieces of content we need to agree on that will make it into the chain, like system contracts, ERC-20 snapshots, etc.. You will see consensus achieved with the members on your first orchestration. Use the sample_config values for now.
## 5. Update `privkeys.keys`
This file should contain the private key(s) to control your seed network account
## 6. Publish your discovery file

    eos-bios publish

## 7. Update `hook_init.sh` and `hook_join_network.sh` to your environement 
The sample config gives you Docker hooks. You can use systemd or Kubernetes!
In `hook_join_network.sh` you need to add your public and private keys
## 8. Orchestrate!
Run 

    eos-bios orchestrate
    
and wait for the boot to happen!

# Free Bonus

## 9. See what's happening
Run one of the following commands:
* `eos-bios discover`
* `eos-bios list`
* `eos-bios discover --serve`


# Network peers

The `seed_network_peers` section of your discovery file looks like this:

```
seed_network_peers:
- account: eosexample
  comment: "They are good"
  weight: 10  # Weights are between 0 and 100 (INT value)
- account: eosmore
  comment: "They are better"
  weight: 20
```

This means you are comfortable launching the network with both
`eosexample` (at 10% vote weight), and `eosmore` (at 20%). `eos-bios`
will compute a graph of the network based on that peering information.

These are all account names on the seed network used to boot a new
network.

# How to weight your peers

1. Do they fully understand the boot sequence ? Do they understand all
   actions that need to be processed in order to have a chain that
   qualifies as mainnet. (they decide on boot_sequence.yaml)

2. Can they compile system contracts and compare their source code,
   making sure that the proposed contracts are legit, do not contain
   rogue code, etc.. (they decide on target_contents)

3. Do they understand how to make sure the snapshot.csv is valid,
   up-to-date and reflect the last Ethereum snapshot ? (they decide on
   snapshot.csv)

4. Can they properly boot the network and have they practiced being
   the BIOS Boot node.

5. Can they properly boot a node and mesh into the network, have they
   practiced `join` ?


The reason for those is because of the design of `eos-bios` .. votes
determine who gets which role, and based on the role you have, you
have critical decisions to make and the community relies on you for
the critical things, in the order above.
