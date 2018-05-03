EOS.IO Software-based blockchain boot orchestration tool
--------------------------------------------------------
[chinese version](./README-cn.md)

`eos-bios` is a command-line tool for people who want to kickstart a
blockchain using EOS.IO Software.

It implements the following:
* Booting a mainnet
* Booting testnets
* Booting local development environments
* Booting consortium or private networks
* Orchestrates launches with many participants through the Network Discovery Protocol.

This program relies on you publishing a `discovery` file.

NOTE: Jump directly to the [sample configurations](./sample_config) if
you know what you're doing.


Launch a local node with a single command
-----------------------------------------

[Download `eos-bios`](https://github.com/eoscanada/eos-bios/releases),
clone this repo, go to `sample_config/docker`:

    git clone https://github.com/eoscanada/eos-bios
    cd eos-bios/sample_config
    wget https://github.com/eoscanada/eos-bios/releases/download/......tar.gz  # Pick the right one for you
    tar -zxvf eos-bios*tar.gz
    mv eos-bios_*/eos-bios .

Then run:

    ./eos-bios boot

Enjoy.

This gives you a fully fledged development environment, a chain loaded
with all system contracts, very similar to what you will get on the
main network once launched.

The sample configuration sets up a single node, as it doesn't point to
other block producer candidates (through the `peers` property).


Boot a network
--------------

By tweaking your configuration, you can grow the network.  To force
the `boot` operation, run:

    ./eos-bios boot

Have your friends run:

    ./eos-bios join --verify

Once your `boot` call finishes properly, it will display the
_Kickstart Data_ that you can provide your friends with.  Once they
paste that in their terminal, they will join your network.

For this to work, you need to properly configure your `config.yaml`
and have properly disseminated your `discovery` files, and given the
URL to your friends.



Join an existing network
------------------------

To join a network, tweak your discovery file to point to the network you're trying to join and publish it. Make sure other participants in the network link to your discovery file as their `peers`.

* Read the [annotated sample configuration file](sample_config/config.yaml).
* Read [sample discovery file here](https://github.com/eoscanada/network-discovery)

Run:

    eos-bios join [--verify]

The `--verify` option toggles the boot sequence verification.


Orchestrate a community launch
------------------------------

When the time comes to orchestrate a launch, *everyone* will run:

    eos-bios orchestrate

According to an algorithm, and using the network discovery data, each
team will be assigned a role deterministically:

1. The _BIOS Boot node_, which will, alone, execute the equivalent of `eos-bios boot`.
2. An _Appointed Block Producer_, which executes the equivalent of `eos-bios join --verify`
3. An _other participant_, which executes the equivalent of `eos-bios join`

The same hooks are used in `boot`, `join` and `orchestrate`, so get
them right and practice.


Example flow and interventions in the orchestrated launch
---------------------------------------------------------

1. Everyone runs `eos-bios orchestrate`.
1. `eos-bios` downloads the network topology pointed to by your `my_discovery_file.yaml`, as does everyone.
1. The network topology is sorted by weight according to how people voted in their `peers` section.
1. The `launch_ethereum_block` is taken from the top 20% in the topology: if they all agree, with continue with that number. Otherwise, we wait until they do (and periodically retraverse the network graph)



Network Discovery Protocol
--------------------------

The Network Discovery Protocol starts with a simple file (see
`my_discovery_file.yaml` in the `sample_config` dir), which is
published to IPFS, and linked through IPNS (like a DNS on IPFS). This
provides a reference that points to your `my_discovery_file.yaml` from
anyone connected to the IPFS network. It looks like:
`/ipns/QmYRsQNxAZFvx8djAxKsgurJT1RF47MhAEQ2sLz1MunnXH`. The last part
is a hash of the public key of the `ipfs` instance running on
someone's computer (which holds the corresponding private key).

In the `my_discovery_file.yaml`, there is (under `launch_data`) a
`peers` key. This allow you to point to other block producer's
published `/ipns/Qm...` link. You can also `weight` the link.

By traversing these links, we can build an in-memory graph of all the
peers connected to one another. Everyone is free to publish when they
want. No need for centralized spreadsheet.

The graph that is created this way can be sorted according to who was
voted for the most. It is public (so don't try to screw anyone), and a
public commitment of whom you're willing to launch with. It is a
decentralized way to vouch for other Block Producers.


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



TODO
----

* In Orchestrate, compute the LaunchData by the most votes, weighted by the highest Weight

* Implement more ethereum sources, so we're not blocked by DDoS of those websites.
  * Could we RPC directly to a swarm of nodes ?
  * We could also load some from the disk, like `ethereum_swarm.txt`

* output.log -> output EVERYTHING to a file, hook a Tee on `os.Stderr`
  and `os.Stdout`.

* No publishing of `secret-p2p-address`, only a remote control of
  `/v1/net/connect` to some ABPs who have published and established
  the network.
  * hook_boot_publish_genesis.sh
  * hook_boot_node.sh
  * hook_boot_connect_mesh.sh
  * hook_boot_publish_privkey.sh

  * hook_join_network.sh  # add connect_count
  * hook_done.sh [role]

* Find out what we do for the chain_id.. do we vote for it too ?
  Top 20% must agree on the chain_id ?
  Top 20% must agree on the constitution ?
