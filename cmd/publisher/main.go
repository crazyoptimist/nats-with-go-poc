package main

import (
	"time"

	"github.com/nats-io/nats.go"
)

func main() {

	url := "localhost:4222"

	nc, _ := nats.Connect(url)

	defer nc.Drain()

	nc.Publish("greet.A", []byte("hello"))
	time.Sleep(time.Second)
	nc.Publish("greet.B", []byte("hello"))
	time.Sleep(time.Second)
	nc.Publish("greet.C", []byte("hello"))
	time.Sleep(time.Second)
}
