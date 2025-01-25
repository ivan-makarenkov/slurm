package main

import (
	"flag"
	"log"

	"github.com/nats-io/nats.go"
)

func init() {
	flag.Parse()
}

// nolint
var (
	uri = flag.String("u", "nats:4222", "NATS host:port")
)

func main() {
	nc, err := nats.Connect(*uri)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	log.Println("connected to", nc.ConnectedAddr())
}
