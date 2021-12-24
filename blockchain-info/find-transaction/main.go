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
		nodeURL   string
		txHash    string
		maxBlocks int
	)
	flag.StringVar(&nodeURL, "node-url", "http://127.0.0.1:7545", "URL of the node to connect to")
	flag.StringVar(&txHash, "tx", "", "transaction hash")
	flag.IntVar(&maxBlocks, "max-blocks", 100, "max blocks to search")
	flag.Parse()

	if txHash == "" {
		log.Fatal("--tx must be specified")
	}

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
	blockNumber := head.Number()

	count := 0
	for {
		count++
		block, err := client.BlockByNumber(ctx, blockNumber)
		if err != nil {
			log.Fatal(err)
		}

		found := false
		for _, tx := range block.Transactions() {
			if tx.Hash().Hex() == txHash {
				fmt.Printf("found in block %d\n", block.Number())
				found = true
			}
		}
		if found {
			break
		}

		blockNumber.Sub(blockNumber, big.NewInt(1))
		if blockNumber.Cmp(big.NewInt(0)) <= 0 {
			fmt.Println("genesis block reached")
		}
		if count == maxBlocks {
			fmt.Printf("transaction not found after searching %d blocks\n", maxBlocks)
			break
		}
	}
}
