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
	// client              *electrum.Client
	tx *[]electrum.Tx
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
	w.getTransactions(client)
	w.balanceCompressed(client)
	w.balanceUncompressed(client)
	if w.balanceConfirmedU != 0 || w.balanceUnconfirmedU != 0 || w.balanceConfirmedC != 0 || w.balanceUnconfirmedC != 0 || len(*w.tx) != 0 {
		saveData(w.String(), out)
	}
}

func (w *wallet) balanceUncompressed(client *electrum.Client) {
	balanceUncompressed, err := client.AddressBalance(w.addressUncompressed)
	if err != nil {
		log.Fatal(err)
	}
	w.balanceConfirmedU = balanceUncompressed.Confirmed
	w.balanceUnconfirmedU = balanceUncompressed.Unconfirmed
}

func (w *wallet) balanceCompressed(client *electrum.Client) {
	balanceCompressed, err := client.AddressBalance(w.addressCompressed)
	if err != nil {
		log.Fatal(err)
	}
	w.balanceConfirmedC = balanceCompressed.Confirmed
	w.balanceUnconfirmedC = balanceCompressed.Unconfirmed
}

func (w *wallet) getTransactions(client *electrum.Client) {
	tx, err := client.AddressHistory(w.addressUncompressed)
	if err != nil {
		log.Fatal(err)
	}
	w.tx = tx
}

func (w *wallet) String() string {
	return fmt.Sprintf("HEX: %s\nWIF: %s\nAddress uncompressed: %s\nAddress compressed: %s\nBalance confirmed uncompressed: %d\nBalance unconfirmed uncompressed: %d\nBalance confirmed compress: %d\nBalance unconfirmed compress: %d\nTx: %v\n",
		w.HEX, w.WIF, w.addressUncompressed, w.addressCompressed, w.balanceConfirmedU, w.balanceUnconfirmedU, w.balanceConfirmedC, w.balanceUnconfirmedC, w.tx,
	)
}
