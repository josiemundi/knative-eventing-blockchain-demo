package main

import (
	"context"
	"github.com/cloudevents/sdk-go"
	"github.com/gorilla/websocket"
	"github.com/kelseyhightower/envconfig"
	"log"
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

	log.Print("Connecting to sink: ", sink)
	t, err := cloudevents.NewHTTPTransport(
		cloudevents.WithTarget(sink),
	)
	if err != nil {
		log.Fatalf("failed to create transport, " + err.Error())
	}

	ce, err := cloudevents.NewClient(t,
		cloudevents.WithUUIDs(),
		cloudevents.WithTimeNow(),
	)

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
		event.SetData(message)

		if _, _, err := ce.Send(context.TODO(), event); err != nil {
			log.Printf("sending event to channel failed: %v", err)
		}
	}
}
