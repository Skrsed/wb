package main

import (
	"encoding/json"
	"fmt"
	"time"
	"wb/producer/domain"
	stream "wb/producer/repository"
	"wb/producer/utils"

	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	// stream factory?
	conn, err := stream.Connect()
	defer stream.Close(conn)

	if err != nil {
		fmt.Println(err)
	}

	for {
		var order domain.Order
		gofakeit.Struct(&order)

		utils.StructPrettyPrint(order)

		jsonData, _ := json.Marshal(order)
		stream.Publish(conn, string(jsonData))
		//slog.Info("Message is published")

		time.Sleep(time.Second * 10)
	}
}
