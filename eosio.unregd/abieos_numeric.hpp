// copyright defined in abieos/LICENSE.txt

#include <algorithm>
#include <array>
#include <stdexcept>
#include <stdint.h>
#include <string>
#include <string_view>
//#include "ripemd160.hpp"

namespace abieos {

const char base58_chars[] = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz";

bool map_initialized = false;
std::array<int8_t, 256> base58_map{{0}};
auto get_base58_map() {
    if(!map_initialized) {
      for (unsigned i = 0; i < base58_map.size(); ++i)
          base58_map[i] = -1;
      for (unsigned i = 0; i < sizeof(base58_chars); ++i)
          base58_map[base58_chars[i]] = i;
      map_initialized = true;
    }
    return base58_map;
}

template <size_t size>
std::array<uint8_t, size> base58_to_binary(std::string_view s) {
    std::array<uint8_t, size> result{{0}};
    for (auto& src_digit : s) {
        int carry = get_base58_map()[src_digit];
        if (carry < 0)
            eosio_assert(0, "invalid base-58 value");
        for (auto& result_byte : result) {
            int x = result_byte * 58 + carry;
            result_byte = x;
            carry = x >> 8;
        }
        if (carry)
            eosio_assert(0, "base-58 value is out of range");
    }
    std::reverse(result.begin(), result.end());
    return result;
}

enum class key_type : uint8_t {
    k1 = 0,
    r1 = 1,
};

struct public_key {
    key_type type{};
    std::array<char, 33> data{};
};


template <typename Key, int suffix_size>
Key string_to_key(std::string_view s, key_type type, const char (&suffix)[suffix_size]) {
    static const auto size = std::tuple_size<decltype(Key::data)>::value;
    auto whole = base58_to_binary<size + 4>(s);
    Key result{type};
    memcpy(result.data.data(), whole.data(), result.data.size());
    return result;
}


public_key string_to_public_key(std::string_view s) {
    if (s.size() >= 3 && s.substr(0, 3) == "EOS") {
        auto whole = base58_to_binary<37>(s.substr(3));
        public_key key{key_type::k1};
        static_assert(whole.size() == key.data.size() + 4, "Error: whole.size() != key.data.size() + 4");
        memcpy(key.data.data(), whole.data(), key.data.size());
        return key;
    } else if (s.size() >= 7 && s.substr(0, 7) == "PUB_R1_") {
        return string_to_key<public_key>(s.substr(7), key_type::r1, "R1");
    } else {
        eosio_assert(0, "unrecognized public key format");
    }
    return public_key{};
}

} // namespace abieos

