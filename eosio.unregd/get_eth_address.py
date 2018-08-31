import sys
from sha3 import keccak_256
from bitcoin import privtopub

priv = sys.argv[1]
pub = privtopub(priv)
addy = keccak_256(pub[2:].decode('hex')).digest()[12:].encode('hex')

print '0x'+addy