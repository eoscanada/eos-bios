Sample configuration for `eos-bios` launch
------------------------------------------

This directory contains all the files necessary for a launch, be it
booting a network or joining an existing network.  You can also use it
for local development.

Table of contents:

* `config.yaml` is a **local** configuration file, read by `eos-bios` to know about its environment.
*
* `base_config.ini`, the base configuration you want to provide to your `nodeos` instance. It is consume by the sample hooks, and shouldn't include any `private_key`, `enable-stale-production` or `producer-name` fields.
