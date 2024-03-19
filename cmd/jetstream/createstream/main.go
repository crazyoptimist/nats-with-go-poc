package main

import (
	"context"
	"encoding/json"
	"fmt"
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

	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatalln("nats connection failed: ", err)
	}

	defer nc.Drain()

	// Build a new JetStream instance
	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatalln("jetstream initialization failed: ", err)
	}

	const streamName = "EVENTS"

	// Remove the stream if it exists already
	_ = js.DeleteStream(context.Background(), streamName)

	// Stream naming convention is SCREAMING CASE
	cfg := jetstream.StreamConfig{
		Name:     streamName,
		Subjects: []string{"events.>"},
	}

	// Persistant file storage is by default, but just to be explicit
	cfg.Storage = jetstream.FileStorage

	// JetStream API uses context for timeout and cancellation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// stream, err := js.CreateStream(ctx, cfg)
	if _, err := js.CreateStream(ctx, cfg); err != nil {
		log.Fatalln("stream creation failed: ", err)
	}

	cfg.MaxAge = time.Hour * 24 * 7
	js.UpdateStream(ctx, cfg)
	fmt.Println("created a stream")

	// DEMO for publish, async publish, 3 types of retension policy
	/*
		js.Publish(ctx, "events.page_loaded", nil)
		js.Publish(ctx, "events.mouse_clicked", nil)
		js.Publish(ctx, "events.mouse_clicked", nil)
		js.Publish(ctx, "events.page_loaded", nil)
		js.Publish(ctx, "events.mouse_clicked", nil)
		js.Publish(ctx, "events.input_focused", nil)
		fmt.Println(ctx, "published 6 messages")

		js.PublishAsync("events.input_changed", nil)
		js.PublishAsync("events.input_blurred", nil)
		js.PublishAsync("events.key_pressed", nil)
		js.PublishAsync("events.input_focused", nil)
		js.PublishAsync("events.input_changed", nil)
		js.PublishAsync("events.input_blurred", nil)

		select {
		case <-js.PublishAsyncComplete():
			fmt.Println("asynchronously published 6 messages")
		case <-time.After(time.Second):
			log.Fatal("async publish took too long")
		}

		printStreamState(ctx, stream)

		cfg.MaxMsgs = 10
		js.UpdateStream(ctx, cfg)
		fmt.Println("set max messages to 10")

		printStreamState(ctx, stream)

		cfg.MaxBytes = 300
		js.UpdateStream(ctx, cfg)
		fmt.Println("set max bytes to 300")

		printStreamState(ctx, stream)

		cfg.MaxAge = time.Second
		js.UpdateStream(ctx, cfg)
		fmt.Println("set max age to one second")

		printStreamState(ctx, stream)

		fmt.Println("sleeping one second...")
		time.Sleep(time.Second)

		printStreamState(ctx, stream)
	*/
}

func printStreamState(ctx context.Context, stream jetstream.Stream) {
	info, _ := stream.Info(ctx)
	b, _ := json.MarshalIndent(info.State, "", " ")
	fmt.Println("inspecting stream info")
	fmt.Println(string(b))
}
