package brute

import (
	"crypto/tls"
	"fmt"
	"log"
	"sync"

	"github.com/fairbank-io/electrum"
)

var wg sync.WaitGroup

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

	// For check server
	// balance, err := client.AddressBalance("3Qj8iqqPxVE14BPWRefGGLJQgaW4xDWW2C")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(balance)

	keys := generate(config.HEX)
	for i := 0; i < config.Threads; i++ {
		wg.Add(1)
		go check(keys, client, config.Output)
	}
	wg.Wait()
}

func generate(hex string) <-chan string {
	out := make(chan string)
	newHex := hex

	go func() {
		for i := 0; i < 20000000; i++ {
			newHex = generateNewHEX(newHex)
			out <- newHex
		}
		close(out)
	}()
	return out
}

func check(keys <-chan string, client *electrum.Client, out string) {
	for key := range keys {
		wallet := newWallet(key)
		wallet.checkBalance(client, out)
		fmt.Println(wallet.String())
	}
	defer wg.Done()
}
