package main

import (
	"fmt"
	"os"
)

func main() {
	name := "vitalik.eth"
	len := len(os.Args)
	if len > 1 {
		name = os.Args[1]
	}
	fmt.Println(name)
}
