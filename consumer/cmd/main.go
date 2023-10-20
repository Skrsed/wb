package main

import (
	adapters "consumer/adaptes"
	"consumer/model"
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
	stan "github.com/nats-io/stan.go"
)

func loadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		slog.Error("Error loading .env file")
	}
}

func initDb() *adapters.DB {
	// Init database
	ctx := context.Background()
	db, err := adapters.NewDB(ctx)

	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}

	slog.Info("Connection was established")

	return db
}

func onMessage(m *stan.Msg) {
	fmt.Println(m)
}

func main() {
	clusterID := "nats_streaming"
	clientID := "wb_consumer"

	loadEnv()
	db := initDb()
	err := addPayment(db, model.Payment{
		Transaction:  "testdata",
		RequestId:    "testdata",
		Currency:     "testdata",
		Provider:     "testdata",
		Amount:       12345,
		PaymentDt:    12345,
		Bank:         "testdata",
		DeliveryCost: 9876,
		GoodsTotal:   9876,
		CustomFee:    9876,
	})
	defer db.Close()

	fmt.Println(payment, err)

	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		fmt.Println(err)
	}

	defer sc.Close()

	sub, _ := sc.Subscribe("foo", onMessage, stan.StartWithLastReceived()) // stan.DeliverAllAvailable()
	defer sub.Close()

	//slog.Info("waiting for messages...")

	for {
		time.Sleep(time.Second * 5)
	}

	// server listen
}
