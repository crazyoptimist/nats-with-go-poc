package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
)

func main() {
	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}

	nc, err := nats.Connect(url)
	if err != nil {
		fmt.Println("connection to nats server failed: ", err)
		os.Exit(1)
	}
	defer nc.Drain()

	sub, err := nc.Subscribe("greet.*", func(msg *nats.Msg) {
		// Parse out the second token in the subject (everything after greet.)
		// and use it as part of the response message.
		name := msg.Subject[6:]
		msg.Respond([]byte("hello, " + name))
	})

	if err != nil {
		fmt.Println("subscribe failed: ", err)
	}

	defer sub.Unsubscribe()

	// Keep the program running until interrupted
	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)
	<-quitChan
}
