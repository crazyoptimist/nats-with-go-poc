package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
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

	streamName := "EVENTS"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := js.Stream(ctx, streamName)
	if err != nil {
		log.Fatalln("getting existing stream failed: ", err)
	}

	consumerName := "processor"

	consumer, err := stream.Consumer(ctx, consumerName)
	if err != nil && errors.Is(err, jetstream.ErrConsumerNotFound) {
		consumer, _ = stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
			Durable: consumerName,
		})
	}

	consumeCtx, err := consumer.Consume(func(msg jetstream.Msg) {
		msg.Ack()
		fmt.Println("received msg on ", msg.Subject())
	})
	if err != nil {
		log.Fatalln("consume failed: ", err)
	}
	defer consumeCtx.Drain()

	// Keep the program running until interrupted
	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)
	<-quitChan
}
