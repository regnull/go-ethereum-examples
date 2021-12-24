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
		nodeURL  string
		txHash   string
		blockNum int64
	)
	flag.StringVar(&nodeURL, "node-url", "http://127.0.0.1:7545", "URL of the node to connect to")
	flag.StringVar(&txHash, "tx", "", "transaction hash")
	flag.Int64Var(&blockNum, "block", -1, "block")
	flag.Parse()

	if txHash == "" {
		log.Fatal("--tx must be specified")
	}

	if blockNum < 0 {
		log.Fatal("--block must be specified")
	}

	// Connect to the node.
	ctx := context.Background()
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		log.Fatal(err)
	}

	block, err := client.BlockByNumber(ctx, big.NewInt(blockNum))
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		if tx.Hash().Hex() == txHash {
			fmt.Printf("hash: %s\n", tx.Hash().Hex())
			fmt.Printf("value: %d\n", tx.Value())
			fmt.Printf("gas: %d\n", tx.Gas())
			fmt.Printf("gas price: %d\n", tx.GasPrice())
			fmt.Printf("nonce: %d\n", tx.Nonce())
			fmt.Printf("data: %v\n", tx.Data())
			fmt.Printf("to: %s\n", tx.To().Hex())
			fmt.Printf("cost: %d\n", tx.Cost())
			break
		}
	}

}
