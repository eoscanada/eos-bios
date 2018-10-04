These are detailed instructions for launching a mainnet
-------------------------------------------------------

If you are appointed Interim Block Producer, here is our experience and tips:

1. In the `sample_config`, tweak `hook_boot_node.sh` to reflect your environment.

1. Run `eos-bios boot --single`.

1. This will generate a `genesis.json` and `genesis.pub` and
   `genesis.key`.

1. If you ever need to restart the chain because injection fails or
   whatever, run: `eos-bios boot --single --reuse-genesis`. This way,
   you don't need to disseminate another `genesis.json` and require
   everyone to clear-up their storage.

1. The node your boot NEEDS --max-transaction-time=5000 for
   transactions not to fail.

1. You should not mesh to anyone during injection. There are already
   bots that are sending automated transactions.. if you are meshed
   with someone exposing an API endpoint, the boot sequence will
   surely be impacted, and validation will most probably fail.

1. Once injection succeeds, and validation passes (the `eos-bios`
   built-in validation) you can enable `author-whitelist = nobody` in
   the configuration, and add a few `p2p-peer-address` statements. Do
   *not* use `contract-whitelist` unless #3418 has been merged and you
   are running updated code.

1. Restart your node, and `resume` production if it was paused on
   startup (remove the `--genesis-json` flag).

1. Blocks should start propagating to the network, but no transaction
   can make it into the chain.

1. Once everyone got the contents, they should be able to validate it
   without issue, as it is read-only.

The community might want to publish reports of the different
validation tools so that everyone gains confidence it is the chain we
want to bring to the world.

1. When the time comes to unlock the chain for transaction, stop the node,
   remove `author-whitelist`


Emergency handoff
-----------------

If your node is unable to produce, its connectivty breaks, or you want to move
that producing node off of your Raspberry Pi, do the following:

* Call `/v1/producer/pause` on the first node.
* Ensure the node to which you are transitioning was fully sync'd and has the last blocks
* Make sure `enable-stale-production = true` on the target node.
* Call `/v1/producer/resume` on the new node or restart without `--pause-on-startup`.
* WARN: do *NOT* hand it off back to the INITIAL node without a PROPER
  RESTART OF THE FIRST NODE. Otherwise, you will hit:
  https://github.com/EOSIO/eos/issues/3442
* If you handoff before unfreezing the chain, make sure the target
  node has `author-whitelist = nobody` also, otherwise you risk
  unfreezing the chain too early.
