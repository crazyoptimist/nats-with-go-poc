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

	sub, err := nc.SubscribeSync("greet.*")
	if err != nil {
		log.Fatalln("subscription failed")
	}

	for {
		msg, _ := sub.NextMsg(10 * time.Millisecond)
		if msg == nil {
			continue
		}
		log.Printf("subject: %q\n msg data: %q\n", msg.Subject, string(msg.Data))
	}
}
