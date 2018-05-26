#include "eosio.unregd.hpp"

using eosio::unregd;

EOSIO_ABI(eosio::unregd, (add))

/**
 * Add a mapping betwen an ethereum_account and an initial EOS token balance.
 */
void unregd::add(const ethereum_account& ethereum_account, const asset& balance) {
  require_auth(_self);

  auto symbol = balance.symbol;
  eosio_assert(symbol.is_valid() && symbol == EOS_SYMBOL, "balance must be EOS token");

  eosio_assert(ethereum_account.length() == 42, "Ethereum account should have exactly 42 characters");

  update_account(ethereum_account, [&](auto& account) {
    account.ethereum_account = ethereum_account;
    account.balance = balance;
  });
}

void unregd::update_account(const ethereum_account& ethereum_account, const function<void(account&)> updater) {
  auto index = accounts.template get_index<N(ethereum_account)>();

  auto itr = index.find(compute_ethereum_account_key256(ethereum_account));
  if (itr == index.end()) {
    accounts.emplace(_self, [&](auto& account) {
      account.id = accounts.available_primary_key();
      updater(account);
    });
  } else {
    index.modify(itr, _self, [&](auto& account) { updater(account); });
  }
}
