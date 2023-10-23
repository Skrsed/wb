package pgRepository

import (
	stan "github.com/nats-io/stan.go"
	"golang.org/x/exp/slog"
)

func Connect() (stan.Conn, error) {
	clusterID := "nats_streaming"
	clientID := "wb_publisher"

	return stan.Connect(clusterID, clientID, stan.NatsURL("nats://localhost:4222"))
}

func Publish(sc stan.Conn, message string) {
	// TODO: .env
	err := sc.Publish("wb_orders", []byte(message))

	if err != nil {
		// TODO: log error
		slog.Error("Error publishing message", "error", err)
	}
}

func Close(sc stan.Conn) {
	sc.Close()
}
