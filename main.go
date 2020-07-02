package main

import (
	"btc_hex_brute/brute"
)

func main() {
	config := brute.NewConfig()
	brute.Run(config)
}
