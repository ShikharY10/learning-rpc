package main

import (
	"os"

	"github.com/ShikharY10/learning-rpc/cmd/handler"
)

func main() {
	args := os.Args[1:]

	if args[0] == "server" {
		handler.StartServer()
	} else if args[0] == "client" {
		handler.StartClient()
	}
}
