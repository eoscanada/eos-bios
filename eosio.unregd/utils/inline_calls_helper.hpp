#pragma once

#include "authority.hpp"

namespace eosio {

   // Helper struct for inline calls.
   // Only the prototype of the functions are used for 
   // the serialization of the action's payload.
   struct call {
      struct token {
         void issue( account_name to, asset quantity, string memo );
         void transfer( account_name from,
                        account_name to,
                        asset        quantity,
                        string       memo );
      };

      struct eosio {
         void newaccount(account_name creator, account_name name, 
                           authority owner, authority active);
         void delegatebw( account_name from, account_name receiver,
                          asset stake_net_quantity, asset stake_cpu_quantity, bool transfer );
         void buyram( account_name buyer, account_name receiver, asset tokens );
      };
   };   
}
