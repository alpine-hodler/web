package main

import (
	"fmt"

	"github.com/alpine-hodler/sdk/pkg/coinbase"
)

func main() {
	// initialize the client connection
	client := coinbase.NewClientEnv("/path/to/.env")

	// then use it for something
	accounts, err := client.Accounts()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", accounts)
}
