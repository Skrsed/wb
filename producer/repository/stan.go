package repository

import (
	"fmt"

	stan "github.com/nats-io/stan.go"
)

func Connect() (stan.Conn, error) {
	clusterID := "nats_streaming"
	clientID := "wb_publisher"

	return stan.Connect(clusterID, clientID, stan.NatsURL("nats://localhost:4222"))
}

func Publish(sc stan.Conn, message []byte) {
	err := sc.Publish("foo", []byte("Hello World"))

	if err != nil {
		// TODO: log error
		fmt.Println(err)
	}
}

func Close(sc stan.Conn) {
	sc.Close()
}
