package brute

import (
	"crypto/tls"
	"fmt"
	"log"

	"github.com/fairbank-io/electrum"
)

// Run ...
func Run(config *Config) {
	client, err := electrum.New(
		&electrum.Options{
			Address:   config.Host,
			TLS:       &tls.Config{InsecureSkipVerify: true},
			Protocol:  electrum.Protocol11,
			KeepAlive: true,
		},
	)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	hex := make(chan string)

	for threads := 0; threads < config.Threads; threads++ {
		go check(hex, client, config.Output)

	}

	generate(config.HEX, hex)
}

func generate(hex string, chHEX chan<- string) {
	newHex := hex
	for {
		newHex = generateNewHEX(newHex)
		chHEX <- newHex
	}
}

func check(hex <-chan string, client *electrum.Client, out string) {
	for {
		wallet := newWallet(<-hex)
		wallet.checkBalance(client, out)
		fmt.Println(wallet.String())
	}
}
