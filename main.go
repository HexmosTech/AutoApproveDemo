package main

import (
	"fmt"

	"github.com/HexmosTech/AutoApproveDemo/internal"
)

func main() {
	fmt.Println(internal.Greeting())
	fmt.Println("version:", internal.Version)
	cfg := internal.DefaultConfig()
	fmt.Println("listening on", cfg.Addr)

	sessions := internal.NewSessionStore()
	sessions.Put(internal.NewSession("u1", "tok-1"))

	metrics := internal.NewCounters()
	metrics.Incr("startup")
}
