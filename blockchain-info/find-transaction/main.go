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
		nodeURL string
		txHash  string
	)
	flag.StringVar(&nodeURL, "node-url", "http://127.0.0.1:7545", "URL of the node to connect to")
	flag.StringVar(&txHash, "tx", "", "transaction hash")
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
	fmt.Println("Chain height:", head.Number())
	blockNumber := head.Number()

	for {
		block, err := client.BlockByNumber(ctx, blockNumber)
		if err != nil {
			log.Fatal(err)
		}

		found := false
		for _, tx := range block.Transactions() {
			if tx.Hash().Hex() == txHash {
				fmt.Println("block: ", block.Number())
				fmt.Println("tx hash: ", tx.Hash().Hex())
				fmt.Println("tx value: ", tx.Value().String())
				fmt.Println("gas: ", tx.Gas())
				fmt.Println("gas price: ", tx.GasPrice().Uint64())
				fmt.Println("nonce: ", tx.Nonce())
				fmt.Println("data: ", tx.Data())
				fmt.Println("receiver: ", tx.To().Hex())
				found = true
			}
			//fmt.Println("Block: ", block.Number(), ", tx: ", tx.Hash().Hex())
		}
		if found {
			break
		}

		blockNumber.Sub(blockNumber, big.NewInt(1))

		// fmt.Println("Block number: ", block.Number().Uint64())
		// fmt.Println("Time: ", block.Time())
		// fmt.Println("Difficulty: ", block.Difficulty().Uint64())
		// fmt.Println("Hash: ", block.Hash().Hex())
		// fmt.Println("Transactions: ", len(block.Transactions()))
	}
}
