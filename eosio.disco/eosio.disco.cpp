#include "eosio.disco/eosio.disco.hpp"

namespace eosio {

  void disco::updtdisco(const uint64_t account, const discovery_file& content) {
    require_auth(account);

    auto disco_itr = discovery_tbl.find(account);
    if (disco_itr == discovery_tbl.end()) {
      discovery_tbl.emplace(_self, [&](auto& row) {
          row.id = account;
          row.content = content;
          row.updated_at = now();
        });
    } else {
      discovery_tbl.modify(disco_itr, _self, [&](auto& row) {
          row.content = content;
          row.updated_at = now();
        });
    }
  }

  void disco::deldisco(const uint64_t account) {
    require_auth(account);

    auto disco_itr = discovery_tbl.find(account);
    eosio_assert(disco_itr != discovery_tbl.end(), "entry not found");
    discovery_tbl.erase(disco_itr);
  }

  void disco::updtgenesis(const account_name account, const string genesis_json, const vector<string> initial_p2p_addresses) {
    require_auth(account);

    auto genesis_itr = genesis_tbl.find(account);
    if (genesis_itr == genesis_tbl.end()) {
      genesis_tbl.emplace(_self, [&](auto& row) {
          row.id = account;
          row.genesis_json = genesis_json;
          row.initial_p2p_addresses = initial_p2p_addresses;
        });
    } else {
      genesis_tbl.modify(genesis_itr, _self, [&](auto& row) {
          row.genesis_json = genesis_json;
          row.initial_p2p_addresses = initial_p2p_addresses;
        });
    }
  }

  void disco::delgenesis(const account_name account) {
    require_auth(account);

    auto genesis_itr = genesis_tbl.find(account);
    eosio_assert(genesis_itr != genesis_tbl.end(), "entry not found");

    genesis_tbl.erase(genesis_itr);
  }

} // namespace eosio

EOSIO_ABI(eosio::disco, (updtgenesis)(updtdisco)(deldisco)(delgenesis))


// GENESIS DATA: genesis_json, initial_p2p_list
// Alex a créé une chain, `seed.eoscanada.com`
// Y'a le contract ci-haut dedans.
//       eosio.disco  w/ update(account_name, target_ipfs_ref)
//
// discovery_file we add:
//   seed_network_chain_id: ad123091823091823091820391820398120938102983019283091283
//
/*

alex$ eos-bios boot --seed-network-endpoint=none or ""
No seed network specified, skipping network discovery
Network:
BIOS Boot: eoscanadacom
---
Booting target network
Producing genesis

    GENESIS DATA

No seed network, not publishing genesis to some seed chain
Injecting bootsequence... done
- No seed network, not meshing with anyone.
Done, boot node up.




FOR POST-LAUNCH:
alex$ eos-bios join 'genesis-data' // inclus genesis_json et initial_p2p_list




alex$ export EOS_BIOS_SEED_NETWORK_ENDPOINT=stage0.eoscanada.com:80
alex$ eos-bios invite eosnation PUB_EOSNATION3123123123123123123123123123
Creating account "eosnation" on the seed network... with pubkey = PUB_EOSNATION...
Transfering 1000 EOS...
Done


denis$ echo "PVT_EOSNATION212121212121" > privkey-cryptolions
# implicit `my_discovery_file.yaml`
denis$ export EOS_BIOS_SEED_NETWORK_ENDPOINT=stage0.eoscanada.com:80
denis$ eos-bios publish discovery --key=privkey-cryptolions
Validating `my_discovery_file.yaml: valid
Loading private key... success.
Signing transaction `eosio.disco::update(eosnation, &Discovery{})`
Submitting transaction to seed network "stage0.eoscanada.com:80"... done



eosrio$ eos-bios discover --seed-network-endpoint myownseed1.eosrio.com:80
Reading `eosio.disco` tables on seed network
Found 82 entries, reading them all
- mamahead: invalid file, excluding from graph
- papahead: invalid file, excluding from graph
Downloading 62 contract, bootseq and snapshot files.
Reading `my_discovery_file.yaml`
- 4 peers: [eosnation @ 0.40, eosblocksmith @ 0.90, eosredhead @ 0.21, eosmyfriend @ 0.00]
  eosnation:
  - 2 peers: [mamahead, papahead]
    mamahead:
    - 1 peer: [eosnation]
    papahead:
    - 1 peer: [eoscanadacom]
Discovered network:
BIOS BOOT:   eosantartica
ABP 1:       eosrioeosrio
ABP 2:       eoscanadacom
ABP 3:       eosblocksmit
...



-----------------------------

eosrio$ export EOS_BIOS_SEED_NETWORK_ENDPOINT=...
eosrio$ export EOS_BIOS_LOCAL_HTTP_ENDPOINT=...
eosrio$ eos-bios orchestrate
... reading network ...
Discovered network:
...
Waiting on block 123456... block arrived! Randomizing network.
Re-reading network:
- Consensus achieved on "boot_sequence"
- Consensus achieved on "snapshot"
- Consensus achieved on contract "eosio.bios"
- Consensus achieved on contract "eosio.system"
- Consensus achieved on contract "eosio.msig"
- Consensus achieved on contract "eosio.token"
BIOS Boot node is: eosantartica (chain_id will be 000000000000000000stage3)
Polling seed chain for new genesis from "eosantartica"... got genesis
Spinning our node with genesis and mesh:
- 12.12.12.12:9876
- 12.23.23.23:9876
- stage2.eoscanada.com:9876
Waiting for validation, polling our http endpoint...
- Validating actions...



eosmama$ export EOS_BIOS_SEED_NETWORK_ENDPOINT=...
eosmama$ export EOS_BIOS_LOCAL_HTTP_ENDPOINT=...
eosmama$ eos-bios orchestrate
... reading network...
Error: these network nodes are not targeting the same chain_id as you:
- eosnation  (chain_id 00000000000000000000000000000000stage1)
- eosnewyork (chain_id 0000000000000000000000000000000000stage2)
- eospapa    (chain_id 0000000000000000000000000000000000000stage3)
EXIT 1


---------------------------------

eosrio$ export EOS_BIOS_SEED_NETWORK_ENDPOINT=...
eosrio$ export EOS_BIOS_LOCAL_HTTP_ENDPOINT=...
eosrio$ eos-bios orchestrate
... reading network ...
Discovered network:
BIOS Boot:   A                 123456
ABP 1:       B                 123223
Waiting on block 123456... block arrived! Randomizing network.
Re-reading network:
- Consensus not achieved on "boot_sequence" by top 15 candidates:
  - eosnation        /ipfs/Qm123123123123123            Updated 5 blocks ago
  - eoscanada        /ipfs/Qm123123123123123            Updated 10 blocks ago
  - eosbob           /ipfs/Qmacdefefedfefeff            Updated 25000 blocks ago
- Consensus not achieved on "snapshot" by top 15 candidates:
  - eosnation        /ipfs/Qm123123123123123            Updated 5 blocks ago
  - eoscanada        /ipfs/Qm123123123123123            Updated 10 blocks ago
  - eosbob           /ipfs/Qmacdefefedfefeff            Updated 25000 blocks ago
Re-reading network in 10 seconds..
- Consensus achieved on "boot_sequence"
- Consensus not achieved on "snapshot" by top 15 candidates:
  - eosnation        /ipfs/Qm123123123123123            Updated 5 blocks ago
  - eoscanada        /ipfs/Qm123123123123123            Updated 0 blocks ago
  - eosbob           /ipfs/Qmacdefefedfefeff            Updated 2 blocks ago
Re-reading network in 10 seconds..
- Consensus achieved on "boot_sequence"
- Consensus achieved on "snapshot"
---
Booting network... etc


...

We're the BIOS Boot node
Generating genesis data:

    GENESISDATA

Publishing to the seed chain.
Booting our node
Injecting bootsequence... done
Meshing boot node with network:
- 123.123.123.123
- 234.234.234.234
- 145.145.145.145
Done, boot node up.


 */
