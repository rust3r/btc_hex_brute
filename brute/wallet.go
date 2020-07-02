package brute

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/fairbank-io/electrum"

	"github.com/btcsuite/btcutil"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
)

type wallet struct {
	HEX                 string
	WIF                 string
	addressUncompressed string
	addressCompressed   string
	balanceConfirmedU   uint64
	balanceUnconfirmedU uint64
	balanceConfirmedC   uint64
	balanceUnconfirmedC uint64
}

func newWallet(HEX string) *wallet {
	decoded, _ := hex.DecodeString(HEX)
	priv, pub := btcec.PrivKeyFromBytes(btcec.S256(), decoded)
	wif, _ := btcutil.NewWIF(priv, &chaincfg.MainNetParams, false)
	uaddress, _ := btcutil.NewAddressPubKey(pub.SerializeUncompressed(), &chaincfg.MainNetParams)
	caddress, _ := btcutil.NewAddressPubKey(pub.SerializeCompressed(), &chaincfg.MainNetParams)
	return &wallet{HEX: HEX, WIF: wif.String(), addressUncompressed: uaddress.EncodeAddress(), addressCompressed: caddress.EncodeAddress()}
}

func (w *wallet) checkBalance(client *electrum.Client, out string) {
	balanceUncompressed, err := client.AddressBalance(w.addressCompressed)
	if err != nil {
		log.Fatal(err)
	}

	w.balanceConfirmedU = balanceUncompressed.Confirmed
	w.balanceUnconfirmedU = balanceUncompressed.Unconfirmed

	if balanceUncompressed.Confirmed != 0 && balanceUncompressed.Unconfirmed != 0 {
		saveData(w.String(), out)
	}

	balanceCompressed, err := client.AddressBalance(w.addressUncompressed)
	if err != nil {
		log.Fatal(err)
	}

	w.balanceConfirmedC = balanceCompressed.Confirmed
	w.balanceUnconfirmedC = balanceCompressed.Unconfirmed

	if balanceCompressed.Confirmed != 0 && balanceCompressed.Unconfirmed != 0 {
		saveData(w.String(), out)
	}
}

func (w *wallet) String() string {
	return fmt.Sprintf("HEX: %s\nWIF: %s\nAddress uncompressed: %s\nAddress compressed: %s\nBalance confirmed uncompressed: %d\nBalance unconfirmed uncompressed: %d\nBalance confirmed compress: %d\nBalance unconfirmed compress: %d\n",
		w.HEX, w.WIF, w.addressUncompressed, w.addressCompressed, w.balanceConfirmedU, w.balanceUnconfirmedU, w.balanceConfirmedC, w.balanceUnconfirmedC,
	)
}
