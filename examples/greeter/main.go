package main

import (
	"github.com/lmuench/plg/rpc/server"

	"github.com/lmuench/plg/plg"
)

func main() {
	plg := plg.NewPLG()
	server.Run(plg)
}
