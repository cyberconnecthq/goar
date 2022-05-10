package main

import (
	"encoding/json"
	"fmt"

	"github.com/everFinance/goar"
	"github.com/everFinance/goar/types"
)

type Box struct {
	Width  int
	Height int
	Color  string
	Open   bool
}

func main() {
	privateKey := ""
	arNode := "https://arweave.net"
	w, err := goar.NewWallet([]byte(privateKey), arNode) // your wallet private key
	if err != nil {
		panic(err)
	}

	box := Box{
		Width:  1000,
		Height: 3000,
		Color:  "yellow",
		Open:   false,
	}
	data, _ := json.Marshal(box)

	box2 := Box{
		Width:  2000,
		Height: 4000,
		Color:  "blue",
		Open:   false,
	}
	data2, _ := json.Marshal(box2)

	target := "Fkj5J8CDLC9Jif4CzgtbiXJBnwXLSrp5AaIllleH_yY"
	tags := []types.Tag{
		{Name: "Content-Type", Value: "application/json"},
		{Name: "Application", Value: "CyberConnect"},
		{Name: "Creator", Value: "0x8ddD03b89116ba89E28Ef703fe037fF77451e38E"},
	}
	item01, err := w.CreateAndSignBundleItem(data, 1, target, "", tags)
	if err != nil {
		panic(err)
	}

	item02, err := w.CreateAndSignBundleItem(data2, 1, target, "", tags)
	if err != nil {
		panic(err)
	}

	items := []types.BundleItem{item01, item02}

	resp, err := w.Client.BatchSendItemToBundler(items, "")
	if err != nil {
		panic(err)
	}

	for _, res := range resp {
		fmt.Println(res.Id)
	}	
}
