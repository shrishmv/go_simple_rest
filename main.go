package main

import (
	"fmt"
	"go_simple_rest/config"
	"go_simple_rest/server"
)

func main() {
	fmt.Println("Starting simple rest server.......")

	config.Init("dev")
	server.Init()
}
