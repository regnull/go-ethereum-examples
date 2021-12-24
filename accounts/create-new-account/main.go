package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {
	var (
		keystoreDir string
		password    string
	)

	flag.StringVar(&keystoreDir, "keystore-dir", "keystore", "keystore directory")
	flag.StringVar(&password, "password", "supersecretpassword", "password")
	flag.Parse()

	ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
	newAcc, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("New account address: ", newAcc.Address.Hex())
}
