# Action - `{{ regaccount }}`

## Description

The intent of the `{{ regaccount }}` action is to register an EOS account using the stored information (Ethereum address/balance) after verifying the submitted ETH signature.

As an authorized party I {{ signer }} wish to create an account {{ account }} accessible with EOS public key {{ eos_pubkey_str }} by submitting cryptographic proof {{ signature }} corresponding to the Ethereum address.

As signer I stipulate that, if I am not the beneficial owner of these tokens, I have proof that I’ve been authorized to take this action by their beneficial owner(s).
