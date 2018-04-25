EOS.IO Software-based blockchain boot orchestration tool
--------------------------------------------------------

`eos-bios` is a command-line tool for people who want to kickstart a
blockchain using EOS.IO Software.

It implements the following:
* Booting a mainnet
* Booting testnets
* Booting local development environments
* Booting consortium or private networks
* Orchestrates launches with many participants through the Network Discovery Protocol.

This program relies on you publishing a `discovery` file. See
https://github.com/eoscanada/network-discovery for more information
about the discovery file.

To discover a network, you need a minimal configuration file like this:

```
network:
  seed_discovery_url: https://abourget.keybase.pub/testnet-cbillett.yaml
  cache_path: /tmp/disco-cbillett
```

The `seed_discovery_url` indicates where to first look for, and `eos-bios` will start traversing the graph created by links to other discovery URLs.

Run `eos-bios` with:

```
$ eos-bios -c config.yaml discover
...
```

To use `eos-bios` to launch a network, you need to add two sections to your local config file:

```
peer:
  my_account: eoscanada   # This identifies you within the graph, matches eosio_account_name in there.
  api_address: http://localhost:8889  # For eos-bios to reach your node and interact with it.
  secret_p2p_address: localhost:19876  # The bit of information you're going to publish to other parties for them to join you, as the BOOT node.
  block_signing_private_key_path: ./privkey-GDW5CV.key  # The private key you want to pass to your hooks, depending on the boot scenario.

hooks:
  init:
    exec: ./hooks/local_init.sh
  boot_network:
    exec: ./hooks/local_boot_network.sh
  join_network:
    exec: ./hooks/local_join_network.sh
  done:
    exec: ./hooks/local_done.sh
```

and run one of the following:

```
$ eos-bios -c config.yaml boot  # Act as the boot node, and run those hooks, run the boot sequence

$ eos-bios -c config.yaml join  # Connect to another node, eventually --verify that actions on chain
                                # conform to what's in the boot_sequence

$ eos-bios -c config.yaml orchestrate  # Use the orchestration algorithms for a decentralized yet
                                       # deterministic way to boot a network together.
```

Install / Download
------------------

You can download the latest release here:
https://github.com/eoscanada/eos-bios/releases .. it is a single
binary that you can download on all major platforms. Simply make
executable and run. It has zero dependencies.

Alternatively, you can build from source with:

    go get -v github.com/eoscanada/eos-bios/eos-bios

This will install the binary in `~/go/bin` provided you have the Go
tool installed (quick install at https://golang.org/dl)


Join the discussion
-------------------

On Telegram through this invite link:
https://t.me/joinchat/GSUv1UaI5QIuifHZs8k_eA (EOSIO BIOS Boot channel)


Previous proposition
--------------------

See the previous proposition in this repo in README.v0.md
