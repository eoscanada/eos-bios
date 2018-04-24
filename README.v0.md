THIS FILE IS KEPT FOR REFERENCE ONLY.  IT IS NOT THE PLAN ANYMORE. THE
PLAN HAS CHANGED, AS ALWAYS!

---

This repository follows up on [Thomas Cox's post](https://medium.com/eosio/bios-boot-eosio-blockchain-2b58b8a978a1) about booting EOS.IO Software.


An EOS BIOS proposal
--------------------

The https://github.com/eoscanada/eos-bios-launch-data contains a simple
file with something like this:

```
launch_btc_block_height: 525123  # Approx June 3rd 2018 9am EST, 6am PST, 3pm UTC.

opening_balances_snapshot_hash: ec2fe55229c3ef2232b5b7fa57175243bf5cf16cb7bbe4a6f8750274f7a56f9a

contract_hashes:
  bios: de6b010347d6ed6f56a5caf406332fbe61a1a485990c5a43484323269ba6b5dd
  system: 21ebeb718e516e727cae6851cb87b0dd040e3cdc1a57dca8d90f88cd8fc1d315
  msig: 8608b380ab76eaa8f8dbe9ebedb4e09f3d4c496d366711d3f09d8865fc0efcb4
  token: c470675ba4809eb1899739394c7913f2582582860ff137854eada1151d1e180c

producers:
- account_name: example
  authority:
    owner:
      threshold: 1
      keys:
      - public_key: EOS8NijGLHT8WyDmt2nqMwfP1hr8EiYx5JCYBWSP9S26WgbeugvSJ
        weight: 1
    active:
      threshold: 1
      keys:
      - public_key: EOS8NijGLHT8WyDmt2nqMwfP1hr8EiYx5JCYBWSP9S26WgbeugvSJ
        weight: 1
  initial_block_signing_key: EOS8NijGLHT8WyDmt2nqMwfP1hr8EiYx5JCYBWSP9S26WgbeugvSJ
  keybase_user: abourget
  organization_name: Example Org
  urls:
  - https://example.com
  - https://twitter.com/example
  - https://github.com/example
```

The current repository drafts a tentative BIOS program, that strives
to streamline and automate the process of kickstarting a new EOS
network.

It can be installed by downloading Go from https://golang.org/dl and running:

    go get github.com/eoscanada/eos-bios

It will build and install a binary in `~/go/bin/eos-bios`.

We will also publish releases for convenience, but building it
yourself is recommended.


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

* Sync your system clock with the rest of the world (run `ntpdate`).


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
         --eosio-system-code ./eosio-system.wasm     \
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

  * That no `eosio_account_name` equal `eosio`, `eosio.auth`,
    `eosio.system` or a few other names that wouldn't be cool.

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

  * The **BIOS Boot node**'s `eos-bios` continues:

    * Generates a new keypair, displays it. Let's call that one the
      `ephemeral key` (to contrast with the producer's key passed
      through `--eosio-secret-key`)

    * Generates a `genesis.json` file, which includes:

      * `initial_key`, set to the generated _ephemeral key_.  This is
        the key used to sign the first transactions and setup the
        chain (see `chain_initializer::get_chain_start_producers`).

      * `initial_timestamp` will be reset to the time of the BTC
        block, or `now()`.

      * `initial_chain_id` will be set to [insert something not dumb]
        (encoded title of a news article of the day?! :)

    * The operator sets these values in his node's `config.ini` (`producer-name = eosio` and `private-key = ["EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV","5KQwrPbwdL6PhXujxW37FSSQZ1JiwsST4cqQzDeyXtP79zkvFD3"]`)

    * The operator boots the node, which starts producing.

    * `eos-bios` is now capable of injecting the system contracts,
      setup the initial producers (unrewarded ABPs). Any transaction
      herein is signed with the `ephemeral key` generated on boot, and
      passed as `initial_key` in the `genesis.json`:

      * `eos-bios` uses the `--bp-api-address` to submit a `setcode`
        transaction to inject the code for the `eosio` account (with
        both `--eosio-system-code` and `--eosio-system-abi`).

      * it also `create account [producer's eosio_account_name]
        [producer's eosio_public_key] [producer's eosio_public_key]`
        for **all producers** listed in `launch.yaml`, in order of the
        shuffle. This is to simplify the takeoff after votes come in.

        * delegate 1 EOS net/bandwidth to all of these, from the
          account on the first line of `snapshot.csv` (the B1
          account), so the BPs listed in `launch.yaml` can write a
          transaction from the get go, if necessary.

      * it `issue`s all opening balances in `eosio` with the contents
        of `snapshot.csv` and creates all the corresponding accounts
        and assigns the pubkeys.  These actions can be batched in a
        few transactions hopefully, maxed at
        `chain_config.max_block_size` (currently 1024*1024 bytes)
        minus some overhead.

        * REVISE: assign unregistered accounts for Ethereum claims post-launch.

        * REVISE with issue / transfer split.
          * Stake 50/50 net band / cpu band, to EVERYONE, so they can vote but not yet transfer.

        * Assign PRIVILEGED to eosio.msig account.

        * DOUBLE-CHECK: any assignation of powers to producers?
          something with multisig powers? (ref. Thomas's mind)

        * TODO: We need to figure out how the Ethereum addresses come
          into play.  For those who haven't registered/claimed, that's
          sort of the last call!


      * is pushes an `updateauth` on the account `eosio` he had
        control over, with a public key similar to
        `EOS0000000000000000000000000000000000000000000000000000000000`,
        rendering the `eosio` account unusable.

        * PERHAPS we should have something more intrinsic, that would
          make that key null, either a privileged primitive that skips
          the `updateauth` checks (that verify the owner key is valid,
          thresholds are sufficient, etc..), and render the account
          permanently disabled.

      * `eos-bios` will then create the _Kickstart data_ file, encrypt it
        for the 21 other ABPs and print it on screen.

      * The operator will publish that content on its social media
        properties.

      * Then, the BIOS Boot node has done its job. He then reverts as
        being only one of the 50+ waiting since the beginning, with
        the sole exception that he knows the address of one of the
        nodes, and can watch the other ones connect.

  * While the Boot node does the steps above, the other 21 ABPs run this process in parallel:

    * Their node is started with production disabled (for now),
      keeping the capacity to interconnect with peers.

    * They print out an encrypted _Kickstart data_ with their OWN
      `p2p_address`, and encrypt it with Keybase/GPG (same as the BIOS
      Node) for a number of ABP (3, 5, all?), in a deterministic
      fashion.. as to create an interesting topology once everyone
      inter-connects.

    * They publish it on their social / web properties.

    * `LOOP HERE`: They pick up the other folks' Kickstart data, and paste it in
      (waiting on stdin).  `eos-bios` would then decrypt it (if it's
      able to).

      * This would reveal the IP of some other nodes, potentially the
        BIOS Boot node.

      * The incoming _Kickstart data_ would be from the BIOS Boot node,
        **or** from one of the ABPs.  What distinguishes the BISO Boot is
        the presence of the private key.

      * This reveals the location of the remote node.

    * `eos-bios` then does one of:

      1. If enabled, issue a call to `/v1/net/connect` on their
         `--bp-api-address` to add the BIOS Boot node address, and
         starts to sync.

      2. If the `eosio::net_api_plugin` isn't enabled, `eos-bios`
         would also print the `config.ini` snippet needed, the
         operator does it manually and boots is node, which would
         connect to the Boot node.

      3. Use the `eos-bios` hooks to automate re-config and restart.

    * At this point:

      * If the BIOS Boot Node is connected through the mesh, the
        network syncs, otherwise, we're just waiting on the BIOS to
        publish it's kickstart data, and we're going back to `LOOP
        HERE`, to connect more nodes.

      * `get account eosio`, and verify that the account has been
        disabled, which marks the end of the BIOS's Boot Node process.

      * The 21 ABPs poll their node (through `--bp-api-address`) until
        they obtain the hash of block 1. They used the
        `private_key_used` in the _Kickstart data_ to validate the
        signare on block 1, proving it was from the BIOS Boot node.

        * If it wasn't, sabotage the network (see below). A few good
          rehearsals should prevent this.

    * At this point, the chain is sync'd from the BIOS Boot, and a
      decently solid network is established.

    * The 21 verify that all of the 21 that were voted have their
      account properly set up with the pubkey in the `launch.yaml`
      file, otherwise they sabotage the network (if they can and
      they're not the ones that were left out with no account/key)

    * The 21 verify the integrity of the Opening Balances in the new
      nascent network, against the locally loaded `snapshot.csv`.

      * `eos-bios` takes a snapshot of `eosio`'s `currency` table and
        compares it locally with `snapshot.csv`.

      * Any failure in verifications would trigger a sabotage.

    * The `eos-bios` program pushes a signed transactions to `eosio`
      system contract, with the `regproducer` action (with
      `--eosio-my-account` and the matching `eosio_public_key` in the
      matching `producers` definition in `launch.yaml`), effectively
      registering the producer on the chain.

    * TODO: when do we check the hashes for the different `setcode` ?

    * TODO: and add other things.. any checks need to be done here

  * At this point, BIOS Boot node is back to normal, as one of the 50+
    persons waiting for which nothing has happened (except perhaps
    seeing who were the ABPs and the BIOS Boot node). They're waiting
    on standard input for the next stage.

  * We come to a point where anyone feeling comfortable can start
    publishing addresses for the whole world to connect (or publishing
    the _Kickstart data_ unencrypted).

    * This would allow all the 50+ who were still waiting, to join in
      using the same logic, albeit with validation disabled (so they
      wouldn't sabotage their account!)

  * `eos-bios` quits, and says thanks or something.

  * The rest of the steps in Thomas Cox's would probably be handled a
    posteri, or by the system contract itself. Some code still needs
    to be written to clarify it all.



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
bitcoin_merkle_root: abcdef123123
```

The `bitcoin_merkle_root` would correspond to the height agreed upon
in `launch.yaml`, so that everyone can check this was created after
Bitcoin's block, and cannot be replayed.

The `private_key_used` would only be present in the _Kickstart data_
coming from the BIOS Boot node.  Other ABPs would not have that.

Sabotaging the network
----------------------

Sabotaging the network means rendering their BP account useless (just
like the `eosio` account is being rendered useless by replacing the
permissions with known-to-be-unknown keys, like
EOS11111111111111111111111111...).

If all ABPs run the BIOS software, they should all sabotage the
network together, and if you falsely sabotage the network, you lost
your chance of being a BP !




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




To be fleshed out
-----------------

* Figure our where `genesis.json` fits in.. perhaps in
  https://github.com/eoscanada/eos-bios-launch-data agreed upon by the
  community.

  * We could add a check by all ABPs

* Regarding initial inflation, and BP average:

  * Good chances that inflation is set a posteri, when the
    constitution kicks in or something, and real Block Producers are
    voted with stakes.. then an avg can be made on their proposition.


Hooks
-----

To ensure a speedy ignition, we've designed a *super simple* remote
control protocol, and provided a binary for you to hook in.  It's
totally optional but can help you speed up the deployment.

By running `eos-bios-rc` near your node and securely port-forwarding
(using `ssh -L`, `kubectl port-forward` or other VPN solution), you
can it react to the different steps in the booting process, for ex:

* reset the storage upon boot
* write some `config.ini` bits and restart your node
* update a `genesis.json` file remotely, and kickstart the node

See `hooks.go`

WARNING: you are on the hook (ha ha) to do any input validation. If a
rogue BP writes an exploit to the `Kickstart data`, it could execute
things on your infrastructure if you haven't checked your things.
`eos-bios` will validate as much as possible its inputs, but be
mindful on your end.



From WAST to WASM
-----------------

The specs for WASM are moving and some syntax for WAST changed during
the development of Dawn 3.

The change occurred in
https://github.com/WebAssembly/wabt/commit/500b617b1c8ea88a2cf46f60205071da9c7569bc
.. changing the syntax for `call_indirect`.

Building with
https://github.com/eoscanada/wabt/tree/prior-to-call_indirect-syntax-change
will produce a `wabt` binary that can read the older format and
produce a valid `.wasm` file. This is simply a fork of `wabt` with a
tag, to facilitate building.
