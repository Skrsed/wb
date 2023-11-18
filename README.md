## WB L0 LEARNING APP

### Run environment:
	docker-compose up -d
### Apply migrations:
	migrate -path ./internal/adapter/postgres/migrations -database postgres://user:password@0.0.0.0:5432/wb?sslmode=disable up
### Start consumer:
	cd consumer
	go run ./cmd/main.go
### Start producer:
	cd producer
	go run ./cmd/main.go
