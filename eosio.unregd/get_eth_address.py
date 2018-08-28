import sys
from sha3 import keccak_256
from bitcoin import *

priv = sys.argv[1]
pub = privtopub(priv)

# sys.stderr.write("PUB----\n")
# sys.stderr.write(pub[2:]+"\n")

addy = keccak_256(pub[2:].decode('hex')).digest()[12:].encode('hex')

print '0x'+addy

# sys.stderr.write("ADDY----\n")
# sys.stderr.write(addy+"\n")
