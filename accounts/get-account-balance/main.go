package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	var (
		nodeURL        string
		accountAddress string
	)
	flag.StringVar(&nodeURL, "node-url", "http://127.0.0.1:7545", "URL of the node to connect to")
	flag.StringVar(&accountAddress, "address", "", "account address")
	flag.Parse()

	// Verify the address argument.
	if accountAddress == "" {
		log.Fatal("--address must be specified")
	}
	account := common.HexToAddress(accountAddress)

	// Connect to the node.
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		log.Fatal(err)
	}

	// Get balance.
	ctx := context.Background()
	balance, err := client.BalanceAt(ctx, account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Balance: ", balance)
}
