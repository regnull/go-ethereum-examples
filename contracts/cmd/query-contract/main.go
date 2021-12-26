package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/regnull/go-ethereum-examples/contracts/store"
)

func main() {
	var (
		nodeURL string
		address string
	)
	flag.StringVar(&nodeURL, "node-url", "http://127.0.0.1:7545", "URL of the node to connect to")
	flag.StringVar(&address, "address", "", "contract's address")
	flag.Parse()

	if address == "" {
		log.Fatal("--address must be specified")
	}

	// Connect to the node.
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		log.Fatal(err)
	}

	instance, err := store.NewStore(common.HexToAddress(address), client)
	if err != nil {
		log.Fatal(err)
	}

	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("contract version: %s\n", version)

	key := [32]byte{}
	copy(key[:], []byte("hello"))
	result, err := instance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("items: %x\n", result)
}
