#!/bin/sh
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;35m'
NC='\033[0m' # No Color

PRIV_HEX=$(openssl rand -hex 32)
ADDY=$(python get_eth_address.py $PRIV_HEX)
AMOUNT=$(python -c 'import random; print "%.4f" % (random.random()*1000)')

echo "* Generated ETH privkey $BLUE$PRIV_HEX$NC"
echo "* Adding $GREEN$ADDY$NC to eosio.unregd db with $GREEN$AMOUNT EOS$NC"
cleos push action eosio.unregd add '["'$ADDY'","'$AMOUNT' EOS"]' -p eosio.unregd

if [ $? -eq 0 ]; then
  echo $GREEN OK $NC
else
  echo $RED ERROR $NC
fi