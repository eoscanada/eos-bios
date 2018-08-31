#!/bin/sh
rm -rf ~/eosio-wallet/./default.wallet
cleos wallet create --to-console 2>&1 | tail -n1 | tail -c+2 | head -c-2 > /tmp/pass

cleos wallet import --private-key 5KQwrPbwdL6PhXujxW37FSSQZ1JiwsST4cqQzDeyXtP79zkvFD3
cleos wallet import --private-key 5KFDkQMpWu9chAEfgQNiEDwRdJePEVJiK92jx6vvUrQA8qFBXUd
cleos wallet import --private-key 5KB513vWai23JYUVB6U8e6oN3z1jApaDDXhp33NuA3urCxgZGMR

cleos create account eosio eosio.token  EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV
cleos create account eosio eosio.bpay   EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV
cleos create account eosio eosio.msig   EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV
cleos create account eosio eosio.names  EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV
cleos create account eosio eosio.ram    EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV
cleos create account eosio eosio.ramfee EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV
cleos create account eosio eosio.saving EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV
cleos create account eosio eosio.stake  EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV
cleos create account eosio eosio.vpay   EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV

cleos set contract eosio.token ~/eos/build/contracts/eosio.token -p eosio.token
cleos push action eosio.token create '[ "eosio", "10000000000.0000 EOS"]' -p eosio.token
cleos push action eosio.token issue '[ "eosio", "1000000000.0000 EOS"]' -p eosio
cleos set contract eosio.msig ~/eos/build/contracts/eosio.msig -p eosio.msig

cleos create account eosio eosio.unregd EOS5BQS4J7f5MKdhTKEXuZQpfwwJDYsikVShaKbVJTmGbDXUvTnXw EOS5BQS4J7f5MKdhTKEXuZQpfwwJDYsikVShaKbVJTmGbDXUvTnXw
cleos set code eosio.unregd ./eosio.unregd.wasm -p eosio.unregd
cleos set abi eosio.unregd ./eosio.unregd.abi -p eosio.unregd
cleos push action eosio.unregd setmaxeos '["1.5000 EOS"]' -p eosio.unregd

cleos set contract eosio ~/dev/eos/build/contracts/eosio.system -p eosio
cleos transfer eosio eosio.unregd "500000.0000 EOS" "sum(unreg_EOS[i])" -p eosio

cleos system newaccount \
--buy-ram-kbytes 4 \
--stake-net "0.2000 EOS" \
--stake-cpu "0.2000 EOS" \
eosio eosio.regram \
EOS6qKoPuTnXK1bh3wHZ6jcuSnJT5T2Ruhkoou8fFKGwRrWqUtB8h \
EOS6qKoPuTnXK1bh3wHZ6jcuSnJT5T2Ruhkoou8fFKGwRrWqUtB8h
cleos transfer eosio eosio.regram "500000.0000 EOS" "para vos" -p eosio

cleos system newaccount \
--buy-ram-kbytes 8 \
--stake-net "2.5000 EOS" \
--stake-cpu "2.5000 EOS" \
eosio thisisatesta \
EOS6qKoPuTnXK1bh3wHZ6jcuSnJT5T2Ruhkoou8fFKGwRrWqUtB8h \
EOS6qKoPuTnXK1bh3wHZ6jcuSnJT5T2Ruhkoou8fFKGwRrWqUtB8h

#Add eosio.unregd@eosio.code to eosio.unreg@active 
tmp=$(mktemp)
cat > $tmp <<EOF
{
  "expiration": "$(date -d '+1 hour' -u +%Y-%m-%dT%H:%M:%S)", 
  "ref_block_num": 0, 
  "ref_block_prefix": 0, 
  "max_net_usage_words": 0, 
  "max_cpu_usage_ms": 0, 
  "delay_sec": 0, 
  "context_free_actions": [], 
  "actions": [
    {
      "account": "eosio", 
      "name": "updateauth", 
      "authorization": [
        {
          "actor": "eosio.unregd", 
          "permission": "active"
        }
      ], 
      "data": "9098ba5303ea305500000000a8ed32320000000080ab26a70100000001000226686e87e1a5752cf8cb85db0477eb4ecaee897044801383c24063e1b442be840100019098ba5303ea305500804a1401ea3055010000"
    }
  ], 
  "transaction_extensions": [], 
  "signatures": [], 
  "context_free_data": []
}
EOF
cleos sign -p -k 5KFDkQMpWu9chAEfgQNiEDwRdJePEVJiK92jx6vvUrQA8qFBXUd $tmp 2>&1 > /dev/null

#Set eosio.regram@[active,owner] to eosio.unregd@eosio.code
cat > $tmp <<EOF
{
  "expiration": "$(date -d '+1 hour' -u +%Y-%m-%dT%H:%M:%S)", 
  "ref_block_num": 0, 
  "ref_block_prefix": 0, 
  "max_net_usage_words": 0, 
  "max_cpu_usage_ms": 0, 
  "delay_sec": 0, 
  "context_free_actions": [], 
  "actions": [
    {
      "account": "eosio", 
      "name": "updateauth", 
      "authorization": [
        {
          "actor": "eosio.regram", 
          "permission": "active"
        }
      ], 
      "data": "20cd65ea02ea305500000000a8ed32320000000080ab26a70100000000019098ba5303ea305500804a1401ea3055010000"
    }, 
    {
      "account": "eosio", 
      "name": "updateauth", 
      "authorization": [
        {
          "actor": "eosio.regram", 
          "permission": "owner"
        }
      ], 
      "data": "20cd65ea02ea30550000000080ab26a700000000000000000100000000019098ba5303ea305500804a1401ea3055010000"
    }
  ], 
  "transaction_extensions": [], 
  "signatures": [], 
  "context_free_data": []
}
EOF
cleos sign -p -k 5KB513vWai23JYUVB6U8e6oN3z1jApaDDXhp33NuA3urCxgZGMR $tmp 2>&1 > /dev/null