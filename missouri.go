package main

import (
	"fmt"
	"log"
	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
)

func echoMsg(ctx context.Context, msg *pubsub.Message) {
	fmt.Println(msg.ID)
}

func main() {
	ctx := context.Background()
	projectId := "iotv-1e541"
	client, err := pubsub.NewClient(ctx, projectId)

	if err != nil {
		log.Fatalf("%v", err)
	}

	subscription := client.Subscription("transcode-raw-videos")
	subscription.Receive(ctx, echoMsg)

	fmt.Println("Hello Missouri")
}
