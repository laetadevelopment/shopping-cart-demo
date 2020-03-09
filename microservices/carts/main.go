package main

import (
	"fmt"
	"os"

	"source.cloud.google.com/laeta-shopping-cart-demo/ShoppingCartDemo"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
