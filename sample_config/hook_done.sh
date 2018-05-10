#!/bin/bash -e

# `done` hook
# $1 = operation ("join", "boot" or "orchestrate")

echo "All done, anything else you want to do ? Call regproducer?"
echo ""
echo "Don't forget to view your logs with:"
echo "   docker logs -f nodeos-bios"
echo ""
echo "Kill your node with:"
echo "   docker kill nodeos-bios"
echo ""
