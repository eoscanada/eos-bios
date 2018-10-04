EOS.IO Software-based blockchain boot tool
------------------------------------------

[点击查看中文](./README-cn.md)

`eos-bios` is a command-line tool for people who want to kickstart a
blockchain using EOS.IO Software. For example:

* Booting local development environments
* Booting testnets
* Booting consortium or private networks

See [sample configurations](./sample_config).


Local development environment
-----------------------------

[Download `eos-bios` from the releases section here on GitHub](https://github.com/eoscanada/eos-bios/releases),
clone this repository and copy the `bootseqs/release-v1.1` to a directory of
your choice.

In directory copied, modify the `base_config.ini` to fit your needs. Usually,
for development purposes, the bind address should be changed from `0.0.0.0` to
`127.0.0.1` for config keys `http-server-address`, `p2p-listen-endpoint`,
and `p2p-server-address` so you are not exposing you development node to
the external world.

Once configuration is done, simply run:

    ./eos-bios boot

This gives you a fully fledged development environment, a chain loaded
with all system contracts, very similar to what you will get on the
main network once launched.

The sample configuration sets up a single node, as it doesn't point to
other block producer candidates (skips the `peers` discovery).

Staged launches
---------------

We keep an updated list of the different stages launched with `eos-bios` here:

https://stages.eoscanada.com



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

Add `-u` to `go get` to pull updates.



Join the discussion
-------------------

On Telegram through this link:
https://t.me/joinchat/GSUv1UaI5QIuifHZs8k_eA (`EOSIO BIOS Boot` channel)



Previous propositions
---------------------

See the previous, deprecated proposition in README.v1.md

See the previous previous, deprecated and never implemented
proposition in README.v0.md
