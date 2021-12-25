# Contracts

This example shows how to compile and deploy a smart contract.

## Install Solidity Complier and abigen

### Install Solidity compiler

```
sudo add-apt-repository ppa:ethereum/ethereum
sudo apt-get update
sudo apt-get install solc
```

### Install abigen

Clone go-etherium GitHub repo:

```
git clone git@github.com:ethereum/go-ethereum.git
```

Compile dev tools:

```
make
make devtools
```

This will install abigen to /usr/local/bin/abigen. Do 'which abigen' to verify.

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

## Deploy the Contract

With the Go wrapper ready, you can run the deployment code under cmd/deploy-contract:

```
go run main.go --keystore-dir=keystore \
    --node-url=http://11.22.33.44:8545 \
    --address=0x13dB07972E645da1f4045b35727cFcF363Ce3994 \
    --password=supersecretpassword
```

Naturally, the account with the given address must already exist and be funded.

You will be given the contract address and the transaction hash. Copy those because you 
will need them in future.

