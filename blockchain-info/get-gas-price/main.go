package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

func main() {
	var (
		nodeURL string
	)
	flag.StringVar(&nodeURL, "node-url", "http://127.0.0.1:7545", "URL of the node to connect to")
	flag.Parse()

	// Connect to the node.
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		log.Fatal(err)
	}

	// Get gas price.
	ctx := context.Background()
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Convert to GWei.
	gasPriceGWei := new(big.Int)
	gasPriceGWei.Div(gasPrice, big.NewInt(params.GWei))

	fmt.Printf("Gas price: %d Wei (%d GWei)\n", gasPrice, gasPriceGWei)
}
