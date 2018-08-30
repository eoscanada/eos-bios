![Alt Text](https://i.imgur.com/6F5aHWH.png)

# Build

eosiocpp -o eosio.unregd.wast eosio.unregd.cpp

# Setup

```./setup.sh```

The setup script will install one contract (besides the defaults ones):
  
  `eosio.unregd` (empty)

You need to have nodeos running.

# Dependecies

 ```pip install bitcoin --user```
 ```pip install requests --user```
 ```sudo apt-get install python-pysha3```

# Test

```./test.sh```

The test step will:

 0. Generate a new ETH address with a random privkey.
 
 1. Add the ETH address (`eosio.unregd::add`) and transfers a random amount of 
    EOS to the `eosio.unregd` contract.
    This is to simulate a user that contributed to the ERC20 but 
    didn't register their EOS pubkey.

 2. Call `eosio.unregd::regaccount` with:

	* A signature for the message "$lib_block_num,$lib_block_prefix" generated with the ETH privkey
	* The desired EOS account name (random in this case)
  * An EOS pubkey (used to set the active/owner permission)

4.  The function will :
  
  * Verify that the destination account name is valid
  * Verify that the account does not exists
  * Extract the pubkey (compressed) from the signature using a message hash composed from the current TX block prefix/num
  * Uncompress the pubkey
  * Calculate ETH address based on the uncompressed publickey
  * Verify that the ETH address exists in eosio.unregd contract
  * Find and split the contribution into cpu/net/liquid
  * Calculate the amount of EOS to purchase 8k of RAM
  * Use the provided EOS key to build the owner/active authority to be used for the account
  * Issue to eosio.unregd the necesary EOS to buy 8K of RAM
  * Create the desired account 
  * Buy RAM for the account (8k)
  * Delegate bandwith
  * Transfer remaining if any (liquid EOS)
  * Remove information for the ETH address from the eosio.unregd DB