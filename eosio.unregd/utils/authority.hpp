#pragma once

namespace eosio {

struct permission_level_weight {
   permission_level  permission;
   weight_type       weight;

   // explicit serialization macro is not necessary, used here only to improve compilation time
   EOSLIB_SERIALIZE( permission_level_weight, (permission)(weight) )
};

struct key_weight {
   eosio::public_key key;
   weight_type       weight;

   // explicit serialization macro is not necessary, used here only to improve compilation time
   EOSLIB_SERIALIZE( key_weight, (key)(weight) )
};

struct wait_weight {
   uint32_t     wait_sec;
   weight_type  weight;

   // explicit serialization macro is not necessary, used here only to improve compilation time
   EOSLIB_SERIALIZE( wait_weight, (wait_sec)(weight) )
};

struct authority {
   uint32_t                         threshold;
   vector<key_weight>               keys;
   vector<permission_level_weight>  accounts;
   vector<wait_weight>              waits;

   // explicit serialization macro is not necessary, used here only to improve compilation time
   EOSLIB_SERIALIZE( authority, (threshold)(keys)(accounts)(waits) )
};

}