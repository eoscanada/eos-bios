This repository follows up on [Thomas Cox's post](https://medium.com/eosio/bios-boot-eosio-blockchain-2b58b8a978a1) about booting EOS.IO Software.


An EOS BIOS proposal
--------------------

The https://github.com/eosioca/eos-bios-launch-data contains a simple
file with something like this:

```
launch_btc_block_height: 525123  # Approx June 3rd 2018 9am EST, 6am PST, 3pm UTC.

opening_balances_snapshot_hash: abcdef123123123

system_contract_hash: 123123abcdef

producers:
- eosio_account_name: acctname
  eosio_public_key: EOSexample
  keybase_user: example
  agent_name: Example EOS.IO BP
  url: https://example.com
```

The current repository drafts a tentative BIOS program, that strives
to streamline and automate the process of kickstarting a new EOS
network.

It can be installed by downloading Go from https://golang.org/dl and running:

    go get github.com/eosioca/eos-bios

It will build and install a binary in `~/go/bin/eos-bios`.

We will also publish releases for convenience, but building it
yourself is recommended (although we are grateful that you trust us to
build public releases).


### Preparation

* Have your desired the `launch.yaml` file (from
  `eos-bios-launch-data`), either by pulling from GitHub, or pulling from
  your friends from the community in a P2P way (in case GitHub gets
  DDoS'd), or however you like.

  * This file would contain an agreed upon Bitcoin block number for
    randomization seed, `launch_btc_block_height`.
  * It contains hashes of the openning balances snapshot csv file, as
    well as of the compiled system contract.
  * It contains the list of block producers that you want in the network.

* A fresh `node` with hardware and network ready, but an empty
  blockchain, with a (compiled and tested) release version of `EOS.IO
  Software` from Block.one

* A freshly dumped ERC-20 token balances snapshot (`snapshot.csv`),
  which matches the `opening_balances_snapshot_hash` in `launch.yaml`.
  See https://github.com/eosio/genesis/tree/0.3.0-beta

* Established DDoS-proof communication channels to send info between
  ABPs. (See below)


### Go-Live

Everyone trying to participate in the Go-Live would execute `eos-bios`
this way:

```bash
eos-bios --launch-data ./launch.yaml                 \
         --eosio-my-account acctname                 \
         --eosio-private-key ./eospriv.key           \
         --keybase-key ./file.key                    \
         --bp-api-address http://1.2.3.4:8888        \
         --bp-p2p-address 1.2.3.4:9876               \
         --eosio-system-code ./eosio-system.wast     \
         --eosio-system-abi ./eosio-system.abi       \
         --opening-balances-snapshot ./snapshot.csv
```

> * `--bp-api-address` is the target API endpoint for the locally booting node, a clean-slate node. It can be routable only from the local machine.
> * `--bp-p2p-address` is the endpoint which will be published at the end of the process
> * `--eosio-my-account` is what links the `launch.yml` with the current instance of the program.
> * `--eosio-private-key` must correspond to the `eosio_public_key` of the current instance's `producers` stanza in `launch.yaml`.
> * `--eosio-system-code` and `--eosio-system-abi` point to the compiled eosio.system contract
> * `--keybase-key` would point to the PGP key, or Keybase something, to decrypt payloads.


This process would:

* Verify that the `--opening-balances-snapshot` hashes to the value in
  `launch.yaml:opening_balances_snapshot_hash`.

* Verify that the `--eosio-system-code` and `--eosio-system-abi` hash
  to `launch.yaml:system_contract_hash` when concatenated. `eos-bios`
  will print the hashes on stdout in any case.. for you to adjust or
  verify with the community.

* Verify there are no duplicates within all these fields from `launch.yaml`:
  `eosio_account_name`, `keybase_user`, `agent_name`, `eosio_public_key`

* Verify there are at least 50 candidates in `producers` list.

* Fetch the Bitcoin block at height
  `launch.yaml:launch_btc_block_height`, take its Merkle Root, massage
  it into an `int64`.

  * We could have 3 sources, like https://blockexplorer.com/
    https://blockchain.info/ and https://live.blockcypher.com/btc/
    chosen randomly by the local process, or a connection to a local
    Bitcoin node.

  * At this point, we have a deterministically random number generator,
    with a value unknown before, fed to [rand.Seed](https://golang.org/pkg/math/rand/#Rand.Seed)

* `eos-bios` would then deterministically shuffle the list of
  producers from `launch.yaml:producers` and select the first 22.
  These are the **Appointed Block Producers** (ABPs). The first of
  them is the **BIOS Boot node**

  * Based on `--eosio-account-name`, your `eos-bios` instance knows if it
    is the Boot node or not.

    * `eos-bios` would print the name of the BIOS Boot node, and URL,
      and ask you to watch for that organization to publish the
      _Kickstart data_ (see below).

      * `eos-bios` could have fetched the other properties linked to
        the Keybase account listed in `launch.yaml`, to display them
        in case Keybase.io goes down while the launch is running.

  * The BIOS Boot node's `eos-bios` continues:

    * Generates a new keypair, displays it.

    * The operator sets these values along with the boot account (`eosio`) in his node's `config.ini` (`producer-name = eosio` and `private-key = ["EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV","5KQwrPbwdL6PhXujxW37FSSQZ1JiwsST4cqQzDeyXtP79zkvFD3"]`)

    * The operator boots the node, which starts producing.

    * `eos-bios` is now capable of injecting the system contracts, setup the initial producers (unrewarded ABPs):

      * `eos-bios` uses the `--bp-api-address` to submit a signed
        `setcode` transaction (it has the generated keys), to inject
        the code for the `eosio` account (both `--eosio-system-code`
        and `--eosio-system-abi`).

      * it also `create account [producer's eosio_account_name] [producer's eosio_public_key] [producer's eosio_public_key]` for all 21 ABPs (except himself) [NOTE: why not for *all* producers listed in the `launch.yaml` file ?]

      * it signs a transaction to `updateauth` on the account `eosio`
        he had control over, with a public key similar to
        `EOS0000000000000000000000000000000000000000000000000000000000`,
        rendering the `eosio` account unusable.

      * `eos-bios` will create the _Kickstart data_ file, encrypt it
        for the 21 other ABPs and print it on screen.

      * The operator will publish that content on its social media
        properties.

      * Then, the BIOS Boot node has done its job. He then reverts as
        being only one of the 50+ waiting since the beginning, with
        the sole exception that he knows the address of one of the
        nodes, and can watch the other ones connect.

  * While the Boot node does the steps above, the other 21 ABPs wait
    on standard input, for the operator to paste the _Kickstart data_
    (see below) from the BIOS Boot node, somewhere on the interwebs.

    * When your launch team discovers that data, it pastes it in
      stdin, and if you're part of the 21 ABPs, you will have the key
      (through `--keybase-key` or a locally running `keybase` instance
      or shelling to `pgp` or something) to decrypt the file and know
      how to continue.

      * This reveals the location of the BIOS Boot node, and the
        private key used to bootstrap that first node.

    * `eos-bios` then does one of:

      1. If enabled, issue a call to `/v1/net/connect` on their
         `--bp-api-address` to add the BIOS Boot node address, and
         starts to sync.

      2. If the `eosio::net_api_plugin` isn't enabled, `eos-bios`
         would also print the `config.ini` snippet needed, the
         operator does it manually and boots is node, which would
         connect to the Boot node.

    * At this point, the network syncs, shouldn't be too long.

    * The 21 ABPs poll their node (through `--bp-api-address`) until
      they obtain the hash of block 1. They used the
      `private_key_used` in the _Kickstart data_ to re-sign block 1,
      proving it was the BIOS Boot node.

      * If it wasn't, sabotage the network (see below). A few good
        rehearsals should prevent this.

    * The 21 verify that all of the 21 that were voted have their
      account properly set up with the pubkey in the `launch.yaml`
      file, otherwise they sabotage the network (if they can and
      they're not the ones that were left out with no account/key)

    * `eos-bios` pushes a signed transactions to `eosio` system
      contract, with the `regproducer` action (with
      `--eosio-my-account` and the matching `eosio_public_key` in the
      matching `producers` definition in `launch.yaml`), effectively
      registering the producer on the chain.

    * When all checks are done `eos-bios` will poll the node and try
      to discover all the other participants, and display them on
      screen.

  * At this point, BIOS Boot node is back to normal, as one of the 50+
    persons waiting for which nothing has happened (except perhaps
    seeing who were the ABPs and the BIOS Boot node). They're waiting
    on standard input for the next stage.

  * It is now the job of the 21 ABPs to finish validation, sabotage
    the network if they find anything odd.

    * The 21 interim BPs verify the integrity of the Opening Balances
      in the new nascent network, against the locally loaded
      `snapshot.csv`.

      * They could query their own node for a few seconds, taking a
        sample of accounts and failing if anything is missing or
        faulty.

      * They could take a full dump of the currency table and compare
        it with the local `snapshot.csv`.

      * Any failure in verifications would initiate a sabotage: like
        the BIOS Boot node did, overwrite their permissions with a
        null key, making this network unuseable, requiring the teams
        to start over.

  * We come to a point where anyone feeling comfortable can start
    publishing addresses for the whole world to connect, or publishing
    the _Kickstart data_ unencrypted.

    * This would allow all the 50+ who were still waiting, to join in
      using the same logic, albeit with validation disabled (so they
      wouldn't sabotage their account!)

  * `eos-bios` quits, and says thanks or something.


Sabotaging the network
----------------------

Sabotaging the network means rendering their BP account useless (just
like the `eosio` account is being rendered useless by replacing the
permissions with known-to-be-unknown keys, like
EOS000000000000000000...).

If all ABPs run the BIOS software, they should all sabotage the
network together, and if you falsely sabotage the network, you lost
your chance of being a BP !



Communications channels
-----------------------

Because of real risks of DDoS at the launch of the EOS Blockchain, the
communication in the setup would use public key cryptography, ideally
Keybase.io, and a few high-profile properties (Twitter, GitHub's gist,
pastebin.com) to share content between the Appointed Block Producers
kickstarting the network.

Each launch team, independently, would monitor the other block
producers' properties (verified beforehand, ideally through Keybase's
social-media vetting system) and see if they publish anything. Where
each team publishes wouldn't be known in advance, thus difficult to
attack.

The `eos-bios` program, could, if we want, have plugins for a few such
properties and automate some of the processes. Otherwise, the teams,
watching for the BIOS Boot Node instructions (in the form of an
encrypted payload) would simply paste the message in the waiting
`eos-bios`.

This is not optimal in terms of speed, as there would be human
intervention, but would satisfy the DDoS protection we need.

There are a few options to speed up comms and make them more
automatic, with varying degrees of resilience / feasibility:

1. Ad-hoc VPN between the nodes (with something like
   http://meshbird.com/, an ad-hoc VPN based on bittorrent-like DHT)

2. Pick some random chat room service

3. Pigeons anyone?


Kickstart data
--------------

The _Kickstart data_ would be an encrypted YAML / JSON, using the 21
Appointed Block Producers' public keys, but no one else's.

That data can be published anywhere, and when pasted in a waiting
`eos-bios` node, it can be decrypted and the process can continue on.

Sample contents:

```
bios_boot_node: 1.2.3.4:9876
private_key_used: 123123123123123123123123
```



To be fleshed out
-----------------

* Figure our where `genesis.json` fits in.. perhaps in
  https://github.com/eosioca/eos-bios-launch-data agreed upon by the
  community.

  * We could add a check by all ABPs



Block Producers publishing their intent
---------------------------------------

A neat way for block producers to publish their intent, or to prepare
a private launch, would be to host/fork their own
`eos-bios-launch-data` repositories and list in there only the
candidates with whom they wish to build the network.

Candidates listed in this repository could be less filtered, and it
would be up to the communities to agree on a common `launch.yaml`.  It
is to be expected that strong teams will want to partner with other
strong teams to build the strongest network.
