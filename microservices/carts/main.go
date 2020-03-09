package main

import (
	"fmt"
	"os"

	"github.com/laetadevelopment/shopping-cart-demo/microservices/carts"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
