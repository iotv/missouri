package main

import (
	"fmt"
	"log"
	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
)

func echoMsg(ctx context.Context, msg *pubsub.Message) {
	msg.Ack()
	fmt.Printf("Msg: %v\n", msg.ID)
}

func main() {
	ctx := context.Background()
	projectId := "iotv-1e541"
	client, err := pubsub.NewClient(ctx, projectId)

	if err != nil {
		log.Fatalf("Client error: %v", err)
	}

	subscription := client.Subscription("transcode-raw-videos")
	exists, err := subscription.Exists(ctx)

	if err != nil {
		log.Fatalf("Exists error: %v", err)
	}

	fmt.Printf("%v", exists)

	for {
		subscription.Receive(ctx, echoMsg)
	}
}
