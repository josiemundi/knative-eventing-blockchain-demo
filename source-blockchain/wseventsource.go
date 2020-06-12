package main

import (
	"context"
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/gorilla/websocket"
	"github.com/kelseyhightower/envconfig"
)

const source = "wss://ws.blockchain.info/inv"

var sink string

type Config struct {
	Sink string `envconfig:"SINK"`
}

func main() {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatalf("failed to process env var: %s", err)
	}

	if config.Sink != "" {
		sink = config.Sink
	}

	ce, err := cloudevents.NewDefaultClient()
	if err != nil {
		log.Fatalf("Failed to create a http cloudevent client: %s", err.Error())
	}

	ctx := cloudevents.ContextWithTarget(context.Background(), sink)

	if err != nil {
		log.Fatalf("unable to create cloudevent client: " + err.Error())
	}

	log.Print("Connecting to source: ", source)
	ws, _, err := websocket.DefaultDialer.Dial(source, nil)
	if err != nil {
		log.Fatal("error connecting:", err)
	}

	err = ws.WriteMessage(websocket.TextMessage, []byte("{\"op\":\"unconfirmed_sub\"}"))
	if err != nil {
		log.Fatal("failed to send subscribe message", err)
	}

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("error while reading message:", err)
			return
		}
		log.Print(string(message))

		event := cloudevents.NewEvent()
		event.SetType("websocket-event")
		event.SetSource(source)
		_ = event.SetData(cloudevents.ApplicationJSON, message)

		if result := ce.Send(ctx, event); !cloudevents.IsACK(result) {
			log.Printf("sending event to channel failed: %v", result)
		}
	}
}
