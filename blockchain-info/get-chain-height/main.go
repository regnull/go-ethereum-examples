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
}
