package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	var prefix string
	fmt.Print("Enter prefix to search for: ")
	fmt.Scanln(&prefix)

	for {
		privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
		if err != nil {
			panic(err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			continue
		}

		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

		if strings.HasPrefix(address, "0x"+prefix) {
			fmt.Printf("Vanity address: %s\n", address)
			privateKeyBytes := privateKey.D.Bytes()
			fmt.Printf("Private key: 0x%s\n", hex.EncodeToString(privateKeyBytes))
			return
		}
	}
}
