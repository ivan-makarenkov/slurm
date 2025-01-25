package main

import (
	"flag"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func init() {
	flag.Parse()
}

// nolint
var (
	uri = flag.String("u", "amqp://guest:guest@rabbitmq:5672/", "AMQP URI")
)

func main() {
	conn, err := amqp.Dial(*uri)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected", conn.LocalAddr())
}
