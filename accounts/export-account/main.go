package main

import (
	"flag"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	var (
		keystoreDir    string
		accountAddress string
		password       string
		exportPassword string
	)
	flag.StringVar(&keystoreDir, "keystore-dir", "keystore", "keystore directory")
	flag.StringVar(&accountAddress, "address", "", "account address")
	flag.StringVar(&password, "password", "", "account password")
	flag.StringVar(&exportPassword, "export-password", "", "export password")
	flag.Parse()

	if accountAddress == "" {
		log.Fatal("--address must be specified")
	}

	if password == "" {
		log.Fatal("--password must be specified")
	}

	if exportPassword == "" {
		log.Fatal("--export-password must be specified")
	}

	ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.Find(accounts.Account{Address: common.HexToAddress(accountAddress)})
	if err != nil {
		log.Fatal(err)
	}

	b, err := ks.Export(account, password, exportPassword)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(b)
}
