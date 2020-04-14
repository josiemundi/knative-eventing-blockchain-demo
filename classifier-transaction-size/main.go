package main

import (
	"context"
	"encoding/json"
	"log"

	"fmt"

	cloudevents "github.com/cloudevents/sdk-go"
	"knative.dev/eventing-contrib/pkg/kncloudevents"
)

func display(event cloudevents.Event, response *cloudevents.EventResponse) {
	evtData, _ := event.DataBytes()

	var transact Transaction
	err := json.Unmarshal(evtData, &transact)
	totalValue := 0
	if err == nil {
		// fmt.Printf("INPUTS:\n")
		for _, v := range transact.X.Inputs {
			// fmt.Println(k, v)
			totalValue += v.PrevOut.Value
		}
		// fmt.Println("")
	}
	classified := "small"
	if totalValue > 100000000 {
		classified = "large"
	}

	newEvent := cloudevents.NewEvent()
	newEvent.SetSource(fmt.Sprintf("https://knative.dev/josiemundi/classifier-transaction-size"))
	newEvent.SetType("dev.knative.eventing.josiemundi.transaction.classifer")
	newEvent.SetID("1234")
	newEvent.SetData(HiFromKnative{Msg:classified})
	response.RespondWith(200, &newEvent)

	//fmt.Println(classified)
	log.Printf("Responded with event %v", newEvent)
}


func main() {

	c, err := kncloudevents.NewDefaultClient()
	if err != nil {
		log.Fatal("Failed to create client, ", err)
	}

	log.Fatal(c.StartReceiver(context.Background(), display))

}
