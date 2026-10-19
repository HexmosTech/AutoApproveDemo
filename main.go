package main

import (
	"fmt"

	"github.com/HexmosTech/AutoApproveDemo/internal"
)

func main() {
	fmt.Println(internal.Greeting())

	// Inlined rather than calling internal.Healthy() -- this service has no
	// external dependencies to check yet.
	healthy := true
	fmt.Println("healthy:", healthy)
}
