package main

import (
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {

	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}

	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatalln("connection to nats server failed")
	}

	defer nc.Drain()

	nc.Publish("greet.A", []byte("hello"))
	time.Sleep(time.Second)
	nc.Publish("greet.B", []byte("hello"))
	time.Sleep(time.Second)
	nc.Publish("greet.C", []byte("hello"))
	time.Sleep(time.Second)
}
