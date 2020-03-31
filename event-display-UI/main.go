package main

import (
	"context"
	"log"
	"net/http"
	"time"

	eventsource "gopkg.in/antage/eventsource.v1"

	"fmt"

	cloudevents "github.com/cloudevents/sdk-go"
	"knative.dev/eventing-contrib/pkg/kncloudevents"
)

var es eventsource.EventSource
var id int

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func display(event cloudevents.Event) {
	fmt.Printf("☁️  cloudevents.Event\n%s", event.String())
	fmt.Println(es)
	evtData, _ := event.DataBytes()
	es.SendEventMessage(string(evtData[:]), "", "")
	id++
}

func main() {
	es = eventsource.New(
		eventsource.DefaultSettings(),
		func(req *http.Request) [][]byte {
			return [][]byte{
				[]byte("X-Accel-Buffering: no"),
				[]byte("Access-Control-Allow-Origin: *"),
			}
		},
	)
	defer es.Close()
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.Handle("/events", es)
	http.HandleFunc("/test", handler)

	port := "9080"

	c, err := kncloudevents.NewDefaultClient()
	if err != nil {
		log.Fatal("Failed to create client, ", err)
	}

	fmt.Println("Starting")
	go func() {
		fmt.Println("Starting to Listen and Serve")
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	}()
	time.Sleep(4000 * time.Millisecond)
	fmt.Println("Middle")
	log.Fatal(c.StartReceiver(context.Background(), display))
	fmt.Println("End")

}
