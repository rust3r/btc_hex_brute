package main

import (
	"btc_hex-brute/brute"
)

func main() {
	config := brute.NewConfig()
	brute.Run(config)
}
