module github.com/josiemundi/knative-eventing-web-eventsource-server

go 1.13

require (
	github.com/cloudevents/sdk-go v0.10.1
	github.com/gorilla/websocket v1.4.1 // indirect
	gopkg.in/antage/eventsource.v1 v1.0.0-20150318155416-803f4c5af225
	knative.dev/eventing-contrib v0.11.0
)
