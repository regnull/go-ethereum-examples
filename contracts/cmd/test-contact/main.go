package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/regnull/go-ethereum-examples/contracts/store"
)

func main() {
	// Generate private key.
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)

	// Create a simulated blockchain.
	alloc := make(core.GenesisAlloc)
	// Balance should be high enough to cover the transaction costs.
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, 10000000)

	// Deploy contract.
	gasPrice, err := blockchain.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice
	address, tx, instance, err := store.DeployStore(auth, blockchain, "v12.34")
	blockchain.Commit()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("contract address: %s\n", address.String())
	fmt.Printf("deploy tx: %s\n", tx.Hash().Hex())

	// Set item.
	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("hello"))
	copy(value[:], []byte("world"))

	tx, err = instance.SetItem(auth, key, value)
	if err != nil {
		log.Fatal(err)
	}
	blockchain.Commit()
	fmt.Printf("set tx: %s\n", tx.Hash().Hex())

	// Query an item.
	result, err := instance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("item: %x\n", result)
}
