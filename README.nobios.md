How to participate in an `eos-bios` launch without using `eos-bios`
-------------------------------------------------------------------

This means participating in ithe `eos-bios` consensus mechanism through the `eosio.disco` contract, yet not running `eos-bios` to either `boot` or `join`.

The `boot` of the network requires a participating to run `eos-bios`,
since all other participants are expecting injection of content based
on its computed consensus content.

The `join` operation can be done manually by simply getting the
genesis.json from the advertised, randomly chosen boot node, finding a
few p2p peers, and booting your node with that configuration in.


Publishing your participation
=============================

To participate, simply take a copy of `sample_config/my_discovery_file.yaml`, convert it to `.json`,
make the appropriate tweaks and inject it in the SEED network this way:

Enclose the resulting `json` in this struct:

```
{
 "account": "youraccountname",
 "disco": CONTENTS OF CONVERTED YAML
}
```

Then

    cleos push action -u http://a-seed-network-node eosio.disco updtdisco "`cat mydisco.json`" -p yourname


Accessing the consensus data
============================

You can also access the genesis from the boot node through this command:

    cleos get table -u ... eosio.disco eosio.disco genesis --limit 1000

You can fetch the other participants' discovery file with:

    cleos get table -u ... eosio.disco eosio.disco discovery --limit 1000

From there, you can pluck out `target_p2p_address` from *active*
participants (those who have an `updated_at` more recent than 30
minutes).

That's !
