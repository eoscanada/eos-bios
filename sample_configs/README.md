# Sample configs for booting a mainnet

This directory contains sample files required to boot a mainnet in a
distributed fashion.

In is an example setup that a Candidate Block Producer would prepare
to participate in a launch.

The different hooks configured here allow a local instance to be
bootstrapped. The hooks can be modified at will to fit your
infrastructure's automation procedures.

## Work in progress

As of March 31st, the encryption with pgp/keybase is not implemented,
so you don't need to fiddle with your keys just yet.

This is a work in progress. Please participate in the discussion in
the
[BIOS Boot Telegram channel](https://t.me/joinchat/GSUv1UaI5QIuifHZs8k_eA)

## Your setup

There are two main configuration files:

1. `config.yaml`, which is your local infrastructure configuration. It
   points to files on your disk, configurations you want to vote for,
   your nodes' IP addresses, your wallet's host:port and the different
   hooks you configured.

1. `launch.yaml` is the file you agreed upon with your fellow BPs,
   with whom you're ready to participate in a launch.

You can grab a sample `snapshot.csv` file from
https://raw.githubusercontent.com/eosdac/airdrop/master/snapshot_282.csv
thanks to our friends at https://github.com/eosdac

When ready, run:

    eos-bios --local-config config.yaml --launch-data launch.yaml

Rinse and repeat.


## Hooks

There are two types of hooks: `exec` and `url`.  The `exec` hook
executes the program configured in the hook, with arguments specific
to each hooks.  See `hooks.go` for the parameters of each.


### `init`

When all basic checks pass, the `init` hook is run.

### `config_ready`

When the BIOS thinks you should start your node, and it has gathered
configuration, it will dispatch the `config_ready` hook, allowing your
node be configured as block producer (or not), inject ephemeral keys
(or not), configure some remote peers, write an appropriate
`genesis.json` file.

At the end of this hook, your node should be running and ready to
respond to HTTP queries on the _host:port_ configured in your
`config.yaml`.
