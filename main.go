package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ENS struct {
	Address string  `json:"address"`
	Balance float32 `json:"balance"`
}

func main() {
	name := "vitalik.eth"
	len := len(os.Args)
	if len > 1 {
		name = os.Args[1]
	}
	res, err := http.Get("https://dakshk.xyz/api/ens?name=" + name)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	response, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var ens ENS
	err = json.Unmarshal(response, &ens)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Address: %s\n", ens.Address)
	fmt.Printf("Balance: %.4f ETH\n", ens.Balance)
}
