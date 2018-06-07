EOS.IO Software-based blockchain boot tool
------------------------------------------

[点击查看中文](./README-cn.md)

`eos-bios` is a command-line tool for people who want to kickstart a
blockchain using EOS.IO Software.

It implements the following:
* Booting local development environments
* Booting testnets
* Booting consortium or private networks

See [sample configurations](./sample_config).


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

On Telegram through this invite link:
https://t.me/joinchat/GSUv1UaI5QIuifHZs8k_eA (EOSIO BIOS Boot channel)



Previous propositions
---------------------

See the previous, deprecated proposition in README.v1.md

See the previous previous, deprecated and never implemented
proposition in README.v0.md
