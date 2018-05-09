EOS.IO Software-based blockchain boot orchestration tool
--------------------------------------------------------

[Chinese version](./README-cn.md)

`eos-bios` is a command-line tool for people who want to kickstart a
blockchain using EOS.IO Software.

It implements the following:
* Multi-staged launch of the mainnet
* Booting local development environments
* Booting testnets
* Booting consortium or private networks

The first you need to know is the **discovery protocol**. See an introduction here:

https://youtu.be/8aNZ_ZnKS-A

Jump directly to the [sample configurations](./sample_config) if you
know what you're doing.


Local development environment
-----------------------------

[Download `eos-bios` from the releases section here on GitHub](https://github.com/eoscanada/eos-bios/releases),
clone this repository and copy the `sample_config` to a directory of
your choice.

Modify the `my_discovery_file.yaml` to point to a local address:

```
target_http_address: http://localhost:8888
```

Then run:

    ./eos-bios boot --single

This gives you a fully fledged development environment, a chain loaded
with all system contracts, very similar to what you will get on the
main network once launched.

The sample configuration sets up a single node, as it doesn't point to
other block producer candidates (skips the `peers` discovery).



Orchestrate a community launch
------------------------------

When the time comes to orchestrate a launch, *everyone* will run:

    ./eos-bios orchestrate

According to an algorithm, and using the network discovery data, each
team will be assigned a role deterministically:

1. The _BIOS Boot node_, which will, alone, execute the equivalent of `eos-bios boot`.
2. An _Appointed Block Producer_, which executes the equivalent of `eos-bios join --verify`
3. An _other participant_, which executes the equivalent of `eos-bios join`


### Practicing to join


Ask anyone on the [latest seed network](#seed-networks) to invite you. They will run:

    ./eos-bios invite [youraccount] [your pubkey]

This will create your account on the seed network. Tweak your `my_discovery_file.yaml`:

* `seed_network_account_name`, make it match what you provided for an invite.
* `seed_network_http_address`, this should be the address of the seed network you want to orchestrate from.

Also add your seed account private key to `privkey.keys`.  This will allow you to run:

    ./eos-bios publish

You will want to tweak those values in your `my_discovery_file.yaml` also:

* `target_http_address` is the address to reach the node you're booting for the next stage.
* `target_p2p_address` is a publicly reachable address advertised to mesh the network
* `target_appointed_block_producer_signing_key` and `target_initial_authority`: these will be injected on the newly created blockchain as `target_account_name`.
* `seed_network_peers` this one warrants [its own section](#network-peers).

Other notable fields:
* `seed_network_launch_block` is the target block on the seed network which will unleash an `orchestrate`d launch.
* `target_contents` are all the pieces of content we need to agree on that will make it into the chain, like system contracts, ERC-20 snapshots, etc..



### Practicing to boot

Review your `my_discovery_file.yaml` (see previous section).  Then run:

    ./eos-bios boot

This will test your `hook_boot_*.sh` hooks.


Discovering the network
-----------------------

When you do point to a seed network (in your config), you can:

    ./eos-bios discover

and this will print the peers, their relative weights, etc..


Network peers
-------------

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


Seed networks
-------------

| Seed network endpoint | Stage | Booted |
| --------------------- | -----:| ------:|
| http://stage0.eoscanada.com | 0 | May 9th 2018 |
| http://stage1.eoscanada.com | 1 | May 9th 2018 |




Example flow and interventions in the orchestrated launch
---------------------------------------------------------

1. Everyone runs `eos-bios orchestrate`.
1. `eos-bios` downloads the network topology pointed to by your `my_discovery_file.yaml`, as does everyone.
1. The network topology is sorted by weight according to how people voted in their `peers` section.
1. The `launch_ethereum_block` is taken from the top 20% in the topology: if they all agree, with continue with that number. Otherwise, we wait until they do (and periodically retraverse the network graph)




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

* output.log -> output EVERYTHING to a file, hook a Tee on `os.Stderr`
  and `os.Stdout`.

* Find out what we do for the chain_id.. do we vote for it too ?
  Top 20% must agree on the chain_id ?
  Top 20% must agree on the constitution ?

* boot_connect_mesh: Make sure we don't mesh with the first BIOS boot..
  it's most probably not running..

* Do connectivity checks when doing `discovery`.. and get a report upon orchestration
  that the peers are up ?

* Make sure `setprods` is called properly... and does affect the producer schedule.

* How do we handle previous `genesis` entries in the blockchain ?
  * Someone joining will be happy to get it.. but if the boot node restarts, everyone needs
    to orchestrate *AFTER*..
  * Clean-up ?
    * We need a way to delete our data in the `eosio.disco` contract..


Role       Seed Account  Target Acct   Weight  Contents          Launch block (local time)
----       ------------  -----------   ------  ----------------  ------------
BIOS NODE  eosmama       eoscanadacom  10      1112211           500 (Fri Nov 7th 23:36, local time)
ABP 01     eosmarc       eosmarc       5       1111112           572 (Fri Nov 8th 00:25, local time)
ABP 02     eosrita       eosrita       2       1111111           572 (Fri Nov 8th 00:25, local time)
ABP 03     eosguy        eosguy        1       1111111           572 (Fri Nov 8th 00:25, local time)
ABP 04     eosbob        eosbob        1

Contents disagreements:
* About column 1: `boot_sequence.yaml`
  * eosmarc, eoscanadacom, eospouet says 1: /ipfs/Qmakjdsflakjdslfkjaldsfk
  * eosmama says 2: /ipfs/Qmhellkajdlakjdsflkj
