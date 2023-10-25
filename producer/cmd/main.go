package main

import (
	"encoding/json"
	"producer/domain"
	stream "producer/repository"
	"producer/utils"
	"time"

	"log/slog"

	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	// stream factory?
	conn, err := stream.Connect()
	defer stream.Close(conn)

	if err != nil {
		slog.Error("Error connecting to nats", "error", err)
	}

	for {
		var order domain.Order
		gofakeit.Struct(&order)

		utils.StructPrettyPrint(order)

		jsonData, _ := json.Marshal(order)
		stream.Publish(conn, string(jsonData))

		time.Sleep(time.Second)
	}
}
