USS_BINARY := ussApp

.PHONY: up up_build down build_uss

up:
	@echo "Starting Docker images..."
	docker compose up -d
	@echo "Docker images started!"

up_build: build_uss
	@echo "Stopping docker images (if running)..."
	docker compose down
	@echo "Building (when required) and starting docker images..."
	docker compose up --build -d
	@echo "Docker images built and started!"

down:
	@echo "Stopping docker compose..."
	docker compose down
	@echo "Done!"

build_uss:
	@echo "Building USS binary for Linux..."
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(USS_BINARY) ./cmd
	@echo "Done!"

run:
	go run cmd/main.go
migrate:
	go run migrations/auto.go