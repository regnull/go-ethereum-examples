# Deploy Contract

This example shows how to compile and deploy a smart contract.

## Install Solidity Complier and abigen

TODO

## Compile Contract

```
solc --optimize --bin contract/store.sol -o build
solc --abi contract/store.sol -o build
```

This will create files Store.abi and Store.bin in build directory.

## Create Go Wrapper

```
abigen --bin=build/Store.bin --abi=build/Store.abi --pkg=store --out=store/Store.go
```

## Run Deployment Code

With the Go wrapper ready, you can run the deployment code under cmd/deploy-contract:

```
go run main.go --keystore-dir=keystore \
    --node-url=http://11.22.33.44:8545 \
    --address=0x13dB07972E645da1f4045b35727cFcF363Ce3994 \
    --password=supersecretpassword
```

Naturally, the account with the given address must already exist and be funded.