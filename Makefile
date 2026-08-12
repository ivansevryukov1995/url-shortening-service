docker-up:
	docker compose up -d
run:
	go run cmd/main.go
migrate:
	go run migrations/auto.go