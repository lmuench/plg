package main

import (
	"log"

	"github.com/lmuench/plg/rpc/server"

	"github.com/lmuench/plg/plg"
)

type Greeter interface {
	Greet(name string)
}

func main() {
	plg := plg.NewPLG()
	ch := make(chan error)
	go server.Run(plg, ch)

	err := <-ch
	if err != nil {
		log.Fatal(err)
	}
	symb, ok := plg.GetSymbol("Greeter")
	if !ok {
		log.Fatal("registry does not contain a Greeter service")
	}
	greeter, ok := symb.(Greeter)
	if ok {
		greeter.Greet("world")
	}
}
