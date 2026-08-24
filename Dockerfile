# Stage 1: Builder
FROM golang:1.26-alpine AS builder
RUN apk add --no-cache git ca-certificates
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/ussApp ./cmd
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/ussMigrate ./migrations/auto.go

# Stage 2: Test
FROM builder AS test
CMD ["go", "test", "-v", "./cmd/..."]

# Stage 3: Prod
FROM alpine:3.21 AS prod
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/ussApp ./ussApp
COPY --from=builder /app/bin/ussMigrate /app/bin/ussMigrate

EXPOSE 8081
CMD ["./ussApp"]
