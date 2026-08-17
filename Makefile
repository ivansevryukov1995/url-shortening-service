USS_BINARY := bin/ussApp
USS_MIGRATE := bin/ussMigrate

.PHONY: up up-build down clean migrate up-db wait-db build_uss build_migrate

build_uss:
	@echo "Building USS binary for Linux..."
	mkdir -p bin
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(USS_BINARY) ./cmd
	@echo "USS binary built: $(USS_BINARY)"

build_migrate:
	@echo "Building Migrate binary for Linux..."
	mkdir -p bin
	GOOS=linux CGO_ENABLED=0 go build -o $(USS_MIGRATE) ./migrations
	@echo "Migrate binary built: $(USS_MIGRATE)"

up: build_uss
	@echo "Starting full stack (data preserved)..."
	docker compose up -d
	@echo "Stack started. Check logs with: docker compose logs -f"

up-build: build_uss
	@echo "Rebuilding and starting full stack (data preserved)..."
	docker compose up -d --build
	@echo "Stack rebuilt and started."

down:
	@echo "Stopping containers (data preserved in volumes)..."
	docker compose down
	@echo "Containers stopped."

migrate: build_migrate
	docker compose run --rm -v "$(PWD)/bin:/work" --entrypoint="/work/ussMigrate" uss-service

up-db:
	@echo "Starting Postgres only..."
	docker compose up -d postgres
	@echo "Postgres started. Wait for 'ready to accept connections' in logs."

clean: down
	@echo "Removing data volumes (ALL DATA WILL BE LOST)..."
	docker compose down -v
	@echo "Cleaned. Data volumes removed."