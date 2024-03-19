package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {

	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}

	nc, _ := nats.Connect(url)
	defer nc.Drain()

	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatalln("jetstream initialization failed: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	js.Publish(ctx, "events.user_clicked_login", nil)
	js.Publish(ctx, "events.user_login_failed", nil)
	js.Publish(ctx, "events.user_rate_limited", nil)
}
