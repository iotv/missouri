package main

import (
	"cloud.google.com/go/pubsub"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"log"
	"os"
	"os/exec"
	"strings"
)

type TranscodeRawVideoMessage struct {
	Name string `json:"name"`
}

type RawVideoCommand struct {
	InputFile string
	OutputDir string
}

func handleMsg(ctx context.Context, msg *pubsub.Message) {
	msg.Ack()
	fmt.Printf("Msg: %v\n", msg.ID)
	fmt.Printf("Data: %s\n", string(msg.Data))

	var trvMsg TranscodeRawVideoMessage
	err := json.Unmarshal(msg.Data, &trvMsg)
	if err != nil {
		log.Fatalf("Unmarshal error: %v", err)
	}

	rvc := RawVideoCommand{
		InputFile: strings.TrimPrefix(trvMsg.Name, "raw-videos/"),
		OutputDir: strings.Split(strings.TrimPrefix(trvMsg.Name, "raw-videos/"), "/")[1],
	}

	// FIXME: this is fairly unsafe. If the path can be manipulated this could jump to the container fs
	cmdArgs := fmt.Sprintf("-nostdin -i /data/raw-videos/%s -hls_time 2 -hls_list_size 0 -f hls -hls_segment_filename /data/videos/%s/%%16d /data/videos/%s/index.m3u8", rvc.InputFile, rvc.OutputDir, rvc.OutputDir)
	fmt.Println(cmdArgs)
	os.MkdirAll("/data/videos/" +rvc.OutputDir, 0777)
	cmd := exec.Command("ffmpeg", strings.Split(cmdArgs, " ")...)
	err = cmd.Run()

	if err != nil {
		fmt.Printf("Command Error: %v", err)
	}
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
	fmt.Printf("Subscription Exists: %v\n", exists)

	for {
		subscription.Receive(ctx, handleMsg)
	}
}
