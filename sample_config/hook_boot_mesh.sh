#!/bin/bash -e

echo "AS THE BIOS BOOT, YOU NOW NEED TO LINK TO THE NETWORK:"
echo "* Your 'config.ini' should contains random peers. Uncomment them."
echo "* Remove the --max-transaction-time, --genesis-json and --p2p-listen-endpoint"
echo "* Restart your nodeos process"


echo "Press ENTER when that is done"
read
