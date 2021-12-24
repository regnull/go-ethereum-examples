package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	var (
		nodeURL         string
		maxTransactions int
	)
	flag.StringVar(&nodeURL, "node-url", "http://127.0.0.1:7545", "URL of the node to connect to")
	flag.IntVar(&maxTransactions, "max-transactions", 100, "max transactions to list")
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
	blockNumber := head.Number()

	// Print recent transactions.
	count := 0
	for {
		block, err := client.BlockByNumber(ctx, blockNumber)
		if err != nil {
			log.Fatal(err)
		}

		done := false
		for _, tx := range block.Transactions() {
			count++
			fmt.Println("Block: ", block.Number(), ", tx: ", tx.Hash().Hex())
			if count == maxTransactions {
				done = true
				break
			}
		}
		if done {
			break
		}

		blockNumber.Sub(blockNumber, big.NewInt(1))
		if blockNumber.Cmp(big.NewInt(0)) <= 0 {
			break
		}
	}
}
