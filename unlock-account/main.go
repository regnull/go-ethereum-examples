package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {
	var (
		keystoreDir    string
		accountAddress string
		password       string
	)
	flag.StringVar(&keystoreDir, "keystore-dir", "keystore", "keystore directory")
	flag.StringVar(&accountAddress, "address", "", "account address")
	flag.StringVar(&password, "password", "", "account password")
	flag.Parse()

	if accountAddress == "" {
		log.Fatal("--address must be specified")
	}

	if password == "" {
		log.Fatal("--password must be specified")
	}

	ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
	var acc *accounts.Account
	for _, a := range ks.Accounts() {
		if a.Address.Hex() == accountAddress {
			acc = &a
			break
		}
	}
	if acc == nil {
		log.Fatal("account not found")
	}

	err := ks.Unlock(*acc, password)
	if err != nil {
		log.Fatal("Failed to unlock account: ", err)
	}
	fmt.Println("Account unlocked successfully")
}
