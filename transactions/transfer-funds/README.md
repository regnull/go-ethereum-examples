# Transfer Funds

This example shows how to transfer funds between accounts.
The account you transfer funds FROM must be in a local keystore,
and you must have the password for this account.


The following example transfers 0.01 Ether (10000000000000000 Wei).

```
go run main.go --keystore-dir=keystore --node-url=http://11.22.33.44:8545  \
    --address-from=0xA1A3c3b7b53Fbf9602f0fC5057eCB7Ae7B2df2eD \
    --address-to=0xFB4Bc64a1849276cA72a3d4dcA3a86d2dDE2F796 \
    --password=supersecretpassword --value=10000000000000000
```