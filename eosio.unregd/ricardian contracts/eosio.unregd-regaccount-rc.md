## ACTION NAME : regaccount

### Description

The intent of the `regaccount` action is to create an EOS account using the stored information {{ Ethereum address }} and token balance from the `eosio.unregd` contract, after verifying the submitted Ethereum {{ signature }}. This is intended to be used only once for each Ethereum address stored in the `eosio.unregd` contract.

As an authorized party, I {{ signer }} wish to create an account {{ account }} on the EOS chain with ID: aca376f206b8fc25a6ed44dbdc66547c36c6c33e3a119ffbeaef943642f0e906, accessible with EOS public key {{ eos_pubkey_str }} by submitting cryptographic proof {{ signature }} corresponding to the {{ Ethereum address }}.

As signer, I stipulate that if I am not the beneficial owner of these tokens, I have been authorized to take this action by the party submitting the cryptographic proof {{signature}}.

In case of dispute, all cases should be brought to the EOS Core Arbitration Forum at https://eoscorearbitration.io/.
