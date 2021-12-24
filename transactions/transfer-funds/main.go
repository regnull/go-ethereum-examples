package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	var (
		keystoreDir string
		nodeURL     string
		addressFrom string
		addressTo   string
		password    string
		value       string
	)
	flag.StringVar(&keystoreDir, "keystore-dir", "keystore", "keystore directory")
	flag.StringVar(&nodeURL, "node-url", "http://127.0.0.1:7545", "URL of the node to connect to")
	flag.StringVar(&addressFrom, "address-from", "", "account address of the transactions originator")
	flag.StringVar(&addressTo, "address-to", "", "account address of the transaction receiver")
	flag.StringVar(&password, "password", "", "account password")
	flag.StringVar(&value, "value", "", "value in wei")
	flag.Parse()

	if addressFrom == "" {
		log.Fatal("--address-from must be specified")
	}

	if addressTo == "" {
		log.Fatal("--address-to must be specified")
	}

	if password == "" {
		log.Fatal("--password must be specified")
	}

	if value == "" {
		log.Fatal("--value must be specified")
	}

	// Connect to the node.
	ctx := context.Background()
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		log.Fatal(err)
	}

	// Open the keystore, find the account, and unlock it.
	ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
	accountFrom, err := ks.Find(accounts.Account{Address: common.HexToAddress(addressFrom)})
	if err != nil {
		log.Fatal(err)
	}
	err = ks.Unlock(accountFrom, password)
	if err != nil {
		log.Fatal(err)
	}

	// Get nonce.
	nonce, err := client.PendingNonceAt(ctx, accountFrom.Address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("got nonce: %d\n", nonce)

	// Recommended gas limit.
	gasLimit := uint64(21000)

	// Get gas price.
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("gas price: %d\n", gasPrice)

	// Send to address.
	toAddress := common.HexToAddress(addressTo)

	// Parse value.
	valueNum := new(big.Int)
	_, ok := valueNum.SetString(value, 10)
	if !ok {
		log.Fatal("failed to parse value")
	}

	// Create and sign the transaction.
	tx := types.NewTransaction(nonce, toAddress, valueNum, gasLimit, gasPrice, nil)

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("chain ID: %d\n", chainID)

	signedTx, err := ks.SignTx(accountFrom, tx, chainID)
	if err != nil {
		log.Fatal(err)
	}

	// Send transaction.
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex())
}
