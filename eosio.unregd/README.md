# Build

eosiocpp -o eosio.unregd.wast eosio.unregd.cpp
eosiocpp -g eosio.unregd.abi eosio.unregd.cpp

# Setup

```./setup.sh```

The setup script will install one contract (besides the defaults ones):
  
  `eosio.unregd` (empty)

You need to have nodeos running.

# Add test data

```./add_test_data.sh```

# Claim

```shell
python claim.py eostest11125 EOS7jUtjvK61eWM38RyHS3WFM7q41pSYMP7cpjQWWjVaaxH5J9Cb7 thisisatesta@active
```

# Dependecies

 ```pip install bitcoin --user```
 ```pip install requests --user```
 ```sudo apt-get install python-pysha3```
