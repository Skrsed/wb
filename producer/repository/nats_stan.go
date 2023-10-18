package main // ?????

import (
	"github.com/nats-io/nats.go"
	stan "github.com/nats-io/stan.go"
)

func Connect() (*stan.conn, err) {
	clusterID := "test-cluster"
	clientID := "wb/publisher"

	return stan.Connect(clusterID, clientID, stan.NatsURL("nats://localhost:4222"))
}

func Publish(sc *stan.conn, message []byte) {
	err := sc.Publish("foo", []byte("Hello World"))

	if err != nil {
		// TODO: log error
	}
}