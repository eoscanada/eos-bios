#include "eosio.unregd.hpp"
#include <eosiolib/crypto.h>
using eosio::unregd;

EOSIO_ABI(eosio::unregd, (add)(regaccount))

/**
 * Add a mapping between an ethereum_address and an initial EOS token balance.
 */
void unregd::add(const ethereum_address& ethereum_address, const asset& balance) {
  require_auth(_self);

  auto symbol = balance.symbol;
  eosio_assert(symbol.is_valid() && symbol == CORE_SYMBOL, "balance must be EOS token");

  eosio_assert(ethereum_address.length() == 42, "Ethereum address should have exactly 42 characters");

  update_address(ethereum_address, [&](auto& address) {
    address.ethereum_address = ethereum_address;
    address.balance = balance;
  });
}

/**
 * Register an EOS account using the stored information (address/balance) verifying an ETH signature
 */
void unregd::regaccount(const bytes& signature, const string& account, const string& eos_pubkey_str) {

  eosio_assert(signature.size() == 66, "Invalid signature");
  eosio_assert(account.size() == 12, "Invalid account length");

  // Verify that the destination account name is valid
  for(const auto& c : account) {
    if(!((c >= 'a' && c <= 'z') || (c >= '1' && c <= '5')))
      eosio_assert(false, "Invalid account name");
  }

  auto naccount = string_to_name(account.c_str());

  // Verify that the account does not exists
  eosio_assert(!is_account(naccount), "Account already exists");

  // Rebuild signed message based on current TX block num/prefix, pubkey and name
  const abieos::public_key eos_pubkey = abieos::string_to_public_key(eos_pubkey_str);

  char tmpmsg[128];
  sprintf(tmpmsg, "%u,%u,%s,%s", tapos_block_num(), tapos_block_prefix(),
    eos_pubkey_str.c_str(), account.c_str());

  //Add prefix and length of signed message
  char message[128];
  sprintf(message, "%s%s%d%s", "\x19", "Ethereum Signed Message:\n", strlen(tmpmsg), tmpmsg);

  //Calculate sha3 hash of message
  sha3_ctx shactx;
  checksum256 msghash;
  rhash_keccak_256_init(&shactx);
  rhash_keccak_update(&shactx, (const uint8_t*)message, strlen(message));
  rhash_keccak_final(&shactx, msghash.hash);

  // Recover compressed pubkey from signature
  uint8_t pubkey[64];
  uint8_t compressed_pubkey[34];
  auto res = recover_key(
    &msghash,
    signature.data(),
    signature.size(),
    (char*)compressed_pubkey,
    34
  );

  eosio_assert(res == 34, "Recover key failed");

  // Decompress pubkey
  uECC_decompress(compressed_pubkey+1, pubkey, uECC_secp256k1());

  // Calculate ETH address based on decompressed pubkey
  checksum256 pubkeyhash;
  rhash_keccak_256_init(&shactx);
  rhash_keccak_update(&shactx, pubkey, 64);
  rhash_keccak_final(&shactx, pubkeyhash.hash);

  uint8_t eth_address[20];
  memcpy(eth_address, pubkeyhash.hash + 12, 20);

  // Verify that the ETH address exists in the "addresses" eosio.unregd contract table
  addresses_index addresses(_self, _self);
  auto index = addresses.template get_index<N(ethereum_address)>();

  auto itr = index.find(compute_ethereum_address_key256(eth_address));
  eosio_assert(itr != index.end(), "Address not found");

  // Split contribution balance into cpu/net/liquid
  auto balances = split_snapshot_abp(itr->balance);
  eosio_assert(balances.size() == 3, "Unable to split snapshot");
  eosio_assert(itr->balance == balances[0] + balances[1] + balances[2], "internal error");

  // Calculate the amount of EOS to purchase 8k of RAM
  auto amount_to_purchase_8kb_of_RAM = buyrambytes(8*1024);

  // Build authority with the pubkey passed as parameter
  const auto auth = authority{
    1, {{{(uint8_t)eos_pubkey.type, eos_pubkey.data} , 1}}, {}, {}
  };

  // Issue to eosio.unregd the necesary EOS to buy 8K of RAM
  INLINE_ACTION_SENDER(call::token, issue)( N(eosio.token), {{N(eosio),N(active)}},
    {N(eosio), amount_to_purchase_8kb_of_RAM,""});

  // Create account with the same key for owner/active
  INLINE_ACTION_SENDER(call::eosio, newaccount)( N(eosio), {{N(eosio),N(active)}},
    {N(eosio), naccount, auth, auth});

  // Buy RAM for this account (8k)
  INLINE_ACTION_SENDER(call::eosio, buyram)( N(eosio), {{N(eosio),N(active)}},
    {N(eosio), naccount, amount_to_purchase_8kb_of_RAM});

  // Delegate bandwith
  INLINE_ACTION_SENDER(call::eosio, delegatebw)( N(eosio), {{N(eosio.unregd),N(active)}},
    {N(eosio.unregd), naccount, balances[0], balances[1], 1});

  // Transfer remaining if any (liquid EOS)
  if( balances[2] != asset(0) ) {
    INLINE_ACTION_SENDER(call::token, transfer)( N(eosio.token), {{N(eosio.unregd),N(active)}},
    {N(eosio.unregd), naccount, balances[2], ""});
  }

  // Remove information for the ETH address from the eosio.unregd DB
  index.erase(itr);
}

void unregd::update_address(const ethereum_address& ethereum_address, const function<void(address&)> updater) {
  auto index = addresses.template get_index<N(ethereum_address)>();

  auto itr = index.find(compute_ethereum_address_key256(ethereum_address));
  if (itr == index.end()) {
    addresses.emplace(_self, [&](auto& address) {
      address.id = addresses.available_primary_key();
      updater(address);
    });
  } else {
    index.modify(itr, _self, [&](auto& address) { updater(address); });
  }
}
