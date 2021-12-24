package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	var nodeURL string
	flag.StringVar(&nodeURL, "node-url", "http://127.0.0.1:7545", "URL of the node to connect to")
	flag.Parse()

	// Connect to the node.
	ctx := context.Background()
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		log.Fatal(err)
	}

	// Get chain height.
	head, err := client.BlockByNumber(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Chain height:", head.Number())

	block, err := client.BlockByNumber(ctx, head.Number())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Block number: ", block.Number().Uint64())
	fmt.Println("Time: ", block.Time())
	fmt.Println("Difficulty: ", block.Difficulty().Uint64())
	fmt.Println("Hash: ", block.Hash().Hex())
	fmt.Println("Transactions: ", len(block.Transactions()))
}
