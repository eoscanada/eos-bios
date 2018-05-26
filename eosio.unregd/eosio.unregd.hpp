#include <functional>
#include <string>

#include <eosiolib/asset.hpp>
#include <eosiolib/eosio.hpp>
#include <eosiolib/fixed_key.hpp>

// Macro
#define TABLE(X) ::eosio::string_to_name(#X)

// Typedefs
typedef std::string ethereum_account;

// Namespaces
using eosio::const_mem_fun;
using eosio::indexed_by;
using eosio::key256;
using std::function;
using std::string;

namespace eosio {

class unregd : public contract {
 public:
  unregd(account_name contract_account)
      : eosio::contract(contract_account), accounts(contract_account, contract_account) {}

  // Actions
  void add(const ethereum_account& ethereum_account, const asset& balance);

 private:
  static const uint64_t EOS_PRECISION = 4;
  static const asset_symbol EOS_SYMBOL = S(EOS_PRECISION, EOS);

  static uint8_t hex_char_to_uint(char character) {
    const int x = character;

    return (x <= 57) ? x - 48 : (x <= 70) ? (x - 65) + 0x0a : (x - 97) + 0x0a;
  }

  static key256 compute_ethereum_account_key256(const ethereum_account& ethereum_account) {
    uint8_t ethereum_key[20];
    const char* characters = ethereum_account.c_str();

    // The ethereum account starts with 0x, let's skip those by starting at i = 2
    for (uint64_t i = 2; i < ethereum_account.length(); i += 2) {
      const uint64_t index = (i / 2) - 1;

      ethereum_key[index] = 16 * hex_char_to_uint(characters[i]) + hex_char_to_uint(characters[i + 1]);
    }

    const uint32_t* p32 = reinterpret_cast<const uint32_t*>(&ethereum_key);
    return key256::make_from_word_sequence<uint32_t>(p32[0], p32[1], p32[2], p32[3], p32[4]);
  }

  //@abi table accounts i64
  struct account {
    uint64_t id;
    ethereum_account ethereum_account;
    asset balance;

    uint64_t primary_key() const { return id; }
    key256 by_ethereum_account() const { return unregd::compute_ethereum_account_key256(ethereum_account); }

    EOSLIB_SERIALIZE(account, (id)(ethereum_account)(balance))
  };

  typedef eosio::multi_index<
      TABLE(accounts), account,
      indexed_by<N(ethereum_account), const_mem_fun<account, key256, &account::by_ethereum_account>>>
      accounts_index;

  void update_account(const ethereum_account& ethereum_account, const function<void(account&)> updater);

  accounts_index accounts;
};

}  // namespace eosio
