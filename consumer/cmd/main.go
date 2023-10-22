package main

//"consumer/itnernal/adapter/postgres/pgRepository"

import (
	cacheRepository "consumer/internal/adapter/cache"
	httpHandler "consumer/internal/adapter/http"
	pgRepository "consumer/internal/adapter/postgres"
	stanHandler "consumer/internal/adapter/stan"

	"consumer/internal/core/service"
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// @title           WB Order consumer service API
// @version         1.0
// @description     This is a sample server for consuming orders over nats.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    t.me/@k_zelenin
// @contact.email  nice.speed.boy@yandex.ru

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  NoAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          bit.ly/wb-golang-task

func loadEnv() {
	// Init env
	err := godotenv.Load("../.env")
	if err != nil {
		slog.Error("Error loading .env file")
	}
}

func initDb() *pgRepository.DB {
	// Init database
	ctx := context.Background()
	db, err := pgRepository.NewDBConnection(ctx, &pgRepository.Credentials{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Host:     "0.0.0.0", //os.Getenv("POSTGRES_HOST"), only inside container
		Port:     os.Getenv("POSTGRES_PORT"),
		DB:       os.Getenv("POSTGRES_DB"),
	})

	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}

	slog.Info("Connection to db was established")

	return db
}

func initStan() *stanHandler.Stan {
	// init nats-streaming
	stanInstance, err := stanHandler.NewStanConnection(&stanHandler.Credentials{
		ClusterID: os.Getenv("NATS_STAN_CLUSTER_ID"),
		ClientID:  os.Getenv("NATS_STAN_CLIENT_ID"),
		StorePort: os.Getenv("NATS_STAN_STORE_PORT"),
	})

	if err != nil {
		slog.Error("Error initializing nats-streaming connection", "error", err)
		os.Exit(1)
	}

	stanInstance.Subscribe()

	slog.Info("Connection to stan was established")

	return stanInstance
}

func initHttp(orderService *service.OrderService) {
	orderHandler := httpHandler.NewOrderHandler(orderService)
	router, err := httpHandler.NewRouter(orderHandler)

	if err != nil {
		slog.Error("Error initializing router", "error", err)
		os.Exit(1)
	}

	// Start server
	httpCredentials := &httpHandler.Credentials{
		Host: os.Getenv("CONSUMER_HOST"),
		Port: os.Getenv("CONSUMER_PORT"),
	}

	slog.Info("Starting the HTTP server on %s:%s", httpCredentials)
	err = router.Serve(httpCredentials)

	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}

func main() {
	loadEnv()
	db := initDb()
	defer db.Close()

	ns := initStan()
	defer ns.Close()

	cache := cacheRepository.NewOrderCacheRepository()
	pg := pgRepository.NewPostgresRepository(db)

	ordSvc, err := service.NewOrderService(pg, cache)
	if err != nil {
		slog.Error("Error while loading dependencies", "error", err)
		os.Exit(1)
	}

	initHttp(ordSvc)

	slog.Info("waiting for messages...")

	for {
		time.Sleep(time.Second * 5)
	}

	// server listen
}
