package main

import (
	"fmt"

	"github.com/beldmian/Rombick/cmd/config"
	"github.com/beldmian/Rombick/cmd/server"
)

func init() {
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}
}

func main() {
	server := server.New()
	if err := server.Start(); err != nil {
		panic(fmt.Errorf("Server error: %s", err))
	}
}
