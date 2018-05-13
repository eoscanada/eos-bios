//
// Discovery protocol contract for `eos-bios` network launches.
//
#pragma once

#include <eosiolib/eosio.hpp>
#include <eosiolib/action.hpp>
#include <eosiolib/public_key.hpp>
#include <eosiolib/types.hpp>
#include <eosiolib/system.h>

using std::string;
using std::vector;

namespace eosiosystem {
  // Copied from eosio.system/native.hpp
  using eosio::permission_level;
  using eosio::public_key;

  struct permission_level_weight {
    permission_level  permission;
    weight_type       weight;

    EOSLIB_SERIALIZE( permission_level_weight, (permission)(weight) )
  };

  struct key_weight {
    public_key   key;
    weight_type  weight;

    EOSLIB_SERIALIZE( key_weight, (key)(weight) )
  };

  struct wait_weight {
    uint32_t wait_sec;
    uint16_t weight;

    EOSLIB_SERIALIZE( wait_weight, (wait_sec)(weight) )
  };

  struct authority {
    uint32_t                              threshold;
    std::vector<key_weight>               keys;
    std::vector<permission_level_weight>  accounts;
    std::vector<wait_weight>              waits;

    EOSLIB_SERIALIZE( authority, (threshold)(keys)(accounts)(waits) )
  };
}


namespace eosio {

  //@abi type
  struct peer_link {
    account_name account;
    string comment;
    uint8_t weight; // betwee 0 and 100
  };

  //@abi type
  struct content_ref {
    string name;
    string ref;
    string comment;
  };

  //@abi type
  struct authority_map {
    eosiosystem::authority owner;
    eosiosystem::authority active;

    EOSLIB_SERIALIZE(authority_map, (owner)(active))
  };

  //@abi type
  struct discovery_file {
    checksum256 seed_network_chain_id;
    account_name seed_network_account_name;
    string seed_network_http_address;
    vector<peer_link> seed_network_peers;
    uint64_t seed_network_launch_block;

    vector<string> urls;
    int16_t gmt_offset; // -0700, +0000

    uint8_t target_network_is_test; //
    checksum256 target_chain_id;
    string target_p2p_address;
    string target_http_address;
    account_name target_account_name;
    public_key target_appointed_block_producer_signing_key;
    authority_map target_initial_authority;

    vector<content_ref> target_contents;

    EOSLIB_SERIALIZE(discovery_file, (seed_network_chain_id)(seed_network_account_name)(seed_network_http_address)(seed_network_peers)(seed_network_launch_block)(urls)(gmt_offset)(target_network_is_test)(target_chain_id)(target_p2p_address)(target_http_address)(target_account_name)(target_appointed_block_producer_signing_key)(target_initial_authority)(target_contents))
  };

  class disco : public contract {
  public:
    disco(account_name self) : contract(self),
                               discovery_tbl(self, self),
                               genesis_tbl(self, self) {}

    //@abi action
    void updtdisco(const account_name account, const discovery_file& disco);
    //@abi action
    void deldisco(const account_name account);
    //@abi action
    void updtgenesis(const account_name account, const string genesis_json, const vector<string> initial_p2p_addresses);
    //@abi action
    void delgenesis(const account_name account);

  private:
    //@abi table genesis i64
    struct genesis_row {
      account_name id;

      string genesis_json;
      vector<string> initial_p2p_addresses;

      uint64_t primary_key() const { return id; }

      EOSLIB_SERIALIZE(genesis_row, (id)(genesis_json)(initial_p2p_addresses))
    };

    //@abi table discovery i64
    struct discovery_row {
      account_name id;
      discovery_file content;
      time updated_at;

      uint64_t primary_key() const { return id; }

      EOSLIB_SERIALIZE(discovery_row, (id)(content)(updated_at))
    };

    typedef eosio::multi_index<N(discovery), discovery_row> discovery_index;
    typedef eosio::multi_index<N(genesis), genesis_row> genesis_index;

    discovery_index discovery_tbl;
    genesis_index genesis_tbl;

  };
} /* namespace eosio */
