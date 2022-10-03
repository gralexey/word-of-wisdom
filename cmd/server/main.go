package main

import (
	"flag"
	"word-of-wisdom/server"
)

var (
	port = flag.Uint("p", 8080, "Port for incoming connections")
)

func main() {
	flag.Parse()
	server.Start(*port)
}
