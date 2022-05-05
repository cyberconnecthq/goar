package example

import (
	"encoding/json"
	"testing"

	"github.com/everFinance/goar"
	"github.com/everFinance/goar/types"
	"github.com/stretchr/testify/assert"
)

type Box struct {
	Width  int
	Height int
	Color  string
	Open   bool
}

func Test_SendData(t *testing.T) {
	arNode := "https://arweave.net"
	w, err := goar.NewWalletFromPath("../arweave-keyfile-Tmi2SGLWNkz6IEr110F1VWewauOdTjN41rZO0rQ16ig.json", arNode) // your wallet private key
	assert.NoError(t, err)

	box := Box{
		Width:  10,
		Height: 20,
		Color:  "blue",
		Open:   false,
	}
	data, _ := json.Marshal(box)
	tags := []types.Tag{
		{Name: "Content-Type", Value: "application/json"},
		{Name: "App-Name", Value: "CyberConnect"},
		{Name: "Contributor", Value: "0x8ddD03b89116ba89E28Ef703fe037fF77451e38E"},
	}
	tx, err := w.SendDataSpeedUp(data, tags, 10)
	assert.NoError(t, err)
	// https://viewblock.io/arweave/tx/RrH-QG23eKwXvuMX6vuZ94l5KTiEU6J2jgWqM3gDF3M
	t.Errorf("tx hash: %s", tx.ID)
}

func Test_LoadData(t *testing.T) {
	arCli := goar.NewClient("https://arweave.net")

	arId := "r90Z_PuhD-louq6uzLTI-xWMfB5TzIti30o7QvW-6A4"
	data, err := arCli.GetTransactionData(arId)
	assert.NoError(t, err)
	t.Log(len(data))
}
