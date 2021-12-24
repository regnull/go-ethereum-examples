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
	flag.StringVar(&nodeURL, "node-url", "", "URL of the node to connect to")
	flag.Parse()

	// Make sure node URL is specified.
	if nodeURL == "" {
		log.Fatal("--node-url must be specified")
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
}
