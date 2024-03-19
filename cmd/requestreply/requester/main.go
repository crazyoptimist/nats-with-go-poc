package main

import (
	"fmt"
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
		fmt.Println("connection to nats server failed: ", err)
		os.Exit(1)
	}
	defer nc.Drain()

	rep, err := nc.Request("greet.joe", nil, time.Second)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(string(rep.Data))
}
