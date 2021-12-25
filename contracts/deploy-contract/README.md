# Deploy Contract

This example shows how to compile and deploy a smart contract.

## Install Solidity Complier and abigen

TODO

## Compile Contract

```
solc --optimize --bin store.sol -o build
solc --abi store.sol -o build
```

This will create files Store.abi and Store.bin in build directory.

## Create Go Wrapper

```
abigen --bin=build/Store.bin --abi=build/Store.abi --pkg=store --out=Store.go
```