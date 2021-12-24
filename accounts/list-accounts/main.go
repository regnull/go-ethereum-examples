package main

import (
	"flag"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {
	var keystoreDir string
	flag.StringVar(&keystoreDir, "keystore-dir", "keystore", "keystore directory")
	flag.Parse()

	ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
	for _, acc := range ks.Accounts() {
		fmt.Println(acc.Address.Hex())
	}
}
