USS_BINARY := bin/ussApp
USS_MIGRATE := bin/ussMigrate

.PHONY: up up-build down clean migrate up-db wait-db build_uss build_migrate gen-secret test-integration

gen-secret:
	echo "SECRET=$(openssl rand -base64 32)" > .secrets

build_uss:
	@echo "Building USS binary for Linux..."
	mkdir -p bin
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(USS_BINARY) ./cmd
	@echo "USS binary built: $(USS_BINARY)"

up-dev:
	@echo "Stopping existing stack..."
	make down
	@echo "Starting dev stack (app + db-dev)..."
	docker compose --profile dev up -d --build
	@echo "Dev stack started. Run 'make migrate' to apply migrations."

restart-dev:
	@echo "Restarting dev stack..."
	docker compose --profile dev restart

build-migrate:
	@echo "Building migrate binary..."
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/ussMigrate ./migrations/auto.go

migrate: build-migrate
	@echo "Waiting for database to be ready..."
	sleep 10
	@echo "Running migrations..."
	./bin/ussMigrate

migrate-docker: build-migrate
	@echo "Running migrations in Docker..."
	docker compose --profile dev run --rm \
		--entrypoint="/app/bin/ussMigrate" \
		app

test-integration:
	@echo "=== Cleaning up previous run ==="
	docker compose --profile test down
	docker compose --profile test up --build --abort-on-container-exit
	@echo "=== Tests finished ==="

down:
	docker compose --profile dev down

clean: down
	@echo "Removing data volumes (ALL DATA WILL BE LOST)..."
	docker compose down -v
	@echo "Cleaned. Data volumes removed."

clean-test-db:
	docker compose --profile test down -v

test-integration-clean: clean-test-db
	docker compose --profile test up --build --abort-on-container-exit