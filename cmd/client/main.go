package main

import (
	"flag"
	"log"
	"word-of-wisdom/client"
)

var (
	address       = flag.String("a", ":8080", "Address of the server")
	requestsCount = flag.Int("c", 1000, "Requests count")
)

func main() {
	flag.Parse()

	client := client.New(*address)

	for i := 0; i < *requestsCount; i++ {
		if wisdom, err := client.GetWisdom(); err != nil {
			log.Println("Error: ", err)
		} else {
			log.Println("Wisdom for you: ", wisdom)
		}
	}
}
