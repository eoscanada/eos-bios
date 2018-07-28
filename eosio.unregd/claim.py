import os
import sys
import json
import time
import requests as req
from datetime import datetime, timedelta
from hashlib import sha256
from bitcoin import ecdsa_raw_sign, encode_privkey
from tempfile import mktemp
from subprocess import Popen, PIPE

def url_for(url):
  return 'http://127.0.0.1:8888{0}'.format(url)

def ref_block(block_id):
  block_num    = block_id[0:8]
  block_prefix = block_id[16:16+8]
  
  ref_block_num     = int(block_num,16)
  ref_block_prefix  = int("".join(reversed([block_prefix[i:i+2] for i in range(0, len(block_prefix), 2)])),16)

  return ref_block_num, ref_block_prefix

if len(sys.argv) < 3:
  print "claim.py ETHPRIV EOSACCOUNT"
  print "    ETHPRIV   : Ethereum private key. (can be in Wif or hex format)"
  print "    EOSACCOUNT: Desired EOS account name"
  sys.exit(1)

block_id = req.get(url_for('/v1/chain/get_info')).json()['last_irreversible_block_id']

ref_block_num, ref_block_prefix = ref_block( block_id )

priv = sys.argv[1]
eos_account = sys.argv[2]
eos_pub = sys.argv[3]

msg = '%d,%d'%(ref_block_num, ref_block_prefix)

msghash = sha256(msg).digest()
v, r, s = ecdsa_raw_sign(msghash, encode_privkey(priv,'hex').decode('hex'))
signature = '00%x%x%x' % (v,r,s)

binargs = req.post(url_for('/v1/chain/abi_json_to_bin'),json.dumps({
  "code"   : "eosio.unregd",
  "action" : "regaccount",
  "args"   : {
    "signature"  : signature,
    "account"    : eos_account,
    "eos_pubkey" : eos_pub
  }
})).json()['binargs']

tx_json = """
{
  "expiration": "%s",
  "ref_block_num": %d,
  "ref_block_prefix": %d,
  "max_net_usage_words": 0,
  "max_cpu_usage_ms": 0,
  "delay_sec": 0,
  "context_free_actions": [],
  "actions": [{
      "account": "eosio.unregd",
      "name": "regaccount",
      "authorization": [{
          "actor": "%s",
          "permission": "active"
        }
      ],
      "data": %s
    }
  ],
  "transaction_extensions": [],
  "signatures": [],
  "context_free_data": []
}
""" % (
  (datetime.utcnow() + timedelta(minutes=3)).strftime("%Y-%m-%dT%T"),
  ref_block_num, 
  ref_block_prefix,
  "thisisatesta",
  binargs
)

tmpf = mktemp()
with open(tmpf,"w") as f:
  f.write(tx_json)

# print tmpf
# sys.exit(0)

with open(os.devnull, 'w') as devnull:
  cmd = ["cleos","sign","-p","-k","5KQwrPbwdL6PhXujxW37FSSQZ1JiwsST4cqQzDeyXtP79zkvFD3",tmpf]
  p = Popen(cmd, stdout=PIPE, stderr=devnull)
  output, err = p.communicate("")

if p.returncode:
  sys.exit(1)

with open(tmpf,"w") as f:
  f.write(output)

print tmpf
sys.exit(0)