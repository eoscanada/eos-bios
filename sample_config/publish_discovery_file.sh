#!/bin/bash

echo "Publishing our discovery file to a well-known location."

# Publish to a location where there is NO CACHING, so the community
# can quickly get updated when you change your mind or your wingmen.

# To avoid a single point of failure, be creative as to where you
# publish it. Choose a highly available location, with DDoS protection

cp my_discovery_file.yaml /keybase/public/my-keybase-user/eos-freezing/testnet-name1.yaml

# This would be made available at: https://my-keybase-user.keybase.pub/eos-freezing/testnet-name1.yaml
