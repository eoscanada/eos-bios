EOS.IO Software-based blockchain boot orchestration tool
--------------------------------------------------------

[点击查看中文](./README-cn.md)

`eos-bios` is a command-line tool for people who want to kickstart a
blockchain using EOS.IO Software.

It implements the following:
* Multi-staged launch of the mainnet
* Booting local development environments
* Booting testnets
* Booting consortium or private networks

The first you need to know is the **discovery protocol**. See an introduction here:

[![The Disco Dance](https://i.ytimg.com/vi/8aNZ_ZnKS-A/hqdefault.jpg?sqp=-oaymwEZCNACELwBSFXyq4qpAwsIARUAAIhCGAFwAQ==&rs=AOn4CLCZeGSGv9Qix8mHX77R4-d0rzDkgA)](https://youtu.be/8aNZ_ZnKS-A)

https://youtu.be/8aNZ_ZnKS-A

Jump directly to the [sample configurations](./sample_config) if you
know what you're doing.

Other **videos** explaining the different concepts of `eos-bios`:

| Description | Link |
| ----------- | ----:|
| Details of the Discovery file  | https://youtu.be/uE5J7LCqcUc |
| Network meshing algorithm  | https://youtu.be/I-PrnlmLNnQ |
| The boot sequence | https://youtu.be/SbVzINnqWAE |
| Code review of the core flow  | https://youtu.be/p1B6jvx25O0 |
| The hooks to integrate with your infrastructure  | https://youtu.be/oZNV6fUoyqM |

Some Q&A related to `eos-bios`:

| Description | Link |
| ----------- | ----:|
| Are accounts carried over from stage to stage?  | https://youtu.be/amyMm5gVpLg |
| How block producers agree on the content of the chain? | https://youtu.be/WIR7nab40qk |



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
2. An _Appointed Block Producer_, which executes the equivalent of `eos-bios join --validate`
3. An _other participant_, which executes the equivalent of `eos-bios join`

The orchestrate for all will wait the agreed upon
`seed_network_launch_block`, and then do the dance.


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


How to weight your peers
------------------------

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


Seed networks
-------------

We keep an updated list of the different stages launched with `eos-bios` here:

https://stages.eoscanada.com



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



Readiness checklist
-------------------

* Did I update my `target_p2p_address` to reflect the IP of the NEW network we're booting ?
* Did I update my `target_http_address` to point to my node, reachable from `eos-bios`'s machine ?
*


Troubleshooting
---------------

* Do your `PRIVKEY` and `PUBKEY` in `hook_join_network.sh` match what
  you published in your discovery file under
  `target_initial_authority` and `target_initial_block_signing_key` ?

* Forked ? Did someone point to an old p2p address ? If so, removed
  them from the network.



TODO
----

* In Orchestrate, compute the LaunchData by the most votes, weighted by the highest Weight

  * Find out what we do for the chain_id.. do we vote for it too ?
    Top 20%, max 5 must agree on the chain_id ?
    Top 20%, max 5 must agree on the constitution ?

* boot_connect_mesh: Make sure we don't mesh with the first BIOS
  boot..  it's most probably not running.. but perhaps it's neat if it
  connects when it is up.. at least to get the blocks.

* Dedupe the p2p-address in join and mesh..

* Display time of the launch block (or delta)

* Implement contents disagreement

* Upgrade to dawn4.1, docker hub, pub eosio.unregd, maintain `eos`
  with CORE_SYMBOL_NAME=EOS applied, publish Docker image, with disco
  and unregd contracts.

* Destroy `eosio`'s signing key also. What happens is someone votes
  for `eosio` enough that he's able to produce ? Then the original
  launcher would have that key.. well it would have been
  published.. so anyone really would have access to that (!!).

* Only download the IPFS refs from the top X that can take the
  decision.. so we're not slowed down by some who never push their
  files.

* p2p chain validation.. should wait until the `target_p2p_address` is
  reachable..  not fail right away.. and retry.

* eos-bios publish: should display the number of seconds left until launch block

* try to get blocks again.. now that ABI things are fixed.. ABI won't
  change during the schebang

* solve the p2p connection error, when doing validation (!!!)
  network-version-match mismatch connection closed because of that..

  * don't use the p2p, rather use the http endpoint for validation..
  * poke until it answers..

* sample_config: Docker image should use dawn 4.1 instead of 4.0.0

* readiness loop:
  * wait for launch block
  * wait for consensus on contents, launch block, chain_id / constitution hash ?
    * sleep 5 seconds
  * wait for participation levels ? publish participation ?

* ipfs: attempt to download from many ipfs gateways, first being `--ipfs`

* do direct meshing (outgoing) from bios boot to other nodes.. don't boot_connect_mesh.sh
  instead the bios boot has a firewalled entrypoint or something..
  connect_mesh could unlock firewall perhaps..

* implement a sort of `--json` output for those wanting to use the tool to do other things..
  * export the `genesis` for a given name, export the disco file for a given name.
  * perhaps simply give example `cleos` calls to do the same..

* Replace `b1` in the contract, or relace `b1b1b1b1b1b1` by `b1` in here.
  * Create the `b1` account before setting the `eosio.system` contract ?

* How will we call `setprods` if it's not in `eosio.bios` anymore ? (!!)
  * Added to eosio.system

* Don't download IPFS refs, unless it was voted on by the top X..
  * Anyone can block the flow right now, otherwise..

* Add a "participate" contract publishing.. so we consider for meshing only those who are
  actively running "orchestrate".

* FIX THE REVERSE target block launch number thing..

* Duplicate target account, fails at account creation.  We should ignore the nodes that
  have multiple KEYS and multiple TARGET ACCOUNTS..
  * Which one we choose ? The first updated at ? we don't want someone to be able to
    take you out merely because they publish a similar `target_account_name`.
  * Let the social graph fix it up?


Desired output for network:

```
Role       Seed Account  Target Acct   Weight  Contents          Launch block (local time)
----       ------------  -----------   ------  ----------------  ------------
BIOS NODE  eosmama       eoscanadacom  10      1112211           500 (Fri Nov 7th 23:36, local time)
ABP 01     eosmarc       eosmarc       5       1111112           572 (Fri Nov 8th 00:25, local time)
ABP 02     eosrita       eosrita       2       1111111           572 (Fri Nov 8th 00:25, local time)
ABP 03     eosguy        eosguy        1       1111111           572 (Fri Nov 8th 00:25, local time)
ABP 04     eosbob        eosbob        1

Contents disagreements:
* About column 4: `boot_sequence.yaml`
  * eosmarc, eoscanadacom, eospouet says 1: /ipfs/Qmakjdsflakjdslfkjaldsfk
  * eosmama says 2: /ipfs/Qmhellkajdlakjdsflkj
```


Things that need testing upon launch
------------------------------------

* Test the eosio.msig machinery, a bunch of BPs need to be able to pass in a proposed
  transaction.
* Cross validae the b1 vesting schedule.
