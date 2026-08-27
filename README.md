<a id='anchor'></a>
# URL Shortening Service

A lightweight [RESTful API](https://roadmap.sh/projects/url-shortening-service) written in **Go** that allows users to shorten long URLs. The service supports creating short links, redirecting to original URLs, and tracking access statistics. Built with **GORM**, **Docker**, and a clear separation of concerns for maintainability.

## Features

- **Authentication & Authorization**: User registration and login via JWT (or session-based) with protected endpoints.
- **URL Shortening**: Create short links from long URLs with automatic hash generation.
- **Redirection**: Access a short link (`/{hash}`) to be redirected to the original URL.
- **Statistics**: View access statistics for your links (requires auth).
- **CRUD for Links**: Create, read, update, and delete your own links (protected by auth).
- **Database Migrations**: Versioned migrations using the `ussMigrate` binary.
- **Dockerized**: Easy setup with Docker Compose for dev and test environments.
-  **Integration Tests**: Dedicated test profile for end-to-end validation of the service.

## Tech Stack

- **Language**: Go 1.26+
- **Web Framework**: `net/http` + custom middleware stack
- **ORM**: GORM with PostgreSQL driver
- **Database**: PostgreSQL (via Docker)
- **Containerization**: Docker & Docker Compose
- **Dependency Injection**: Custom DI container (`pkg/di`)
- **Configuration**: Environment variables + optional `.env` (via `godotenv`)
- **Build Automation**: Makefile with targets for build, migrate, test, and deploy

## Requirements

- Go 1.26.3 or higher
- Docker & Docker Compose
- Make (optional but recommended)
- OpenSSL (for generating secrets)

## How to Run

### Option 1: Local Development (App Local, DB in Docker)

This approach runs the Go app locally while using Docker only for the database. Ideal for quick iteration and debugging.

1. Clone the repository:
   ```bash
   git clone https://github.com/ivansevryukov1995/url-shortening-service.git
   cd url-shortening-service
   ```
2. Create a .env file based on .env.example:
    ```bash
    DATABASE_URL=postgresql://user:pass@localhost:5433/link_db?sslmode=disable
    APP_HOST=0.0.0.0
    APP_PORT=8081
    ```
    and .secrets    
    ```bash
    SECRET=your-secret-key
    ```
3. Start the database:
    ```bash
    make up-dev
    ```
4. Run migrations:
    ```bash
    make migrate
    ```
5. Start the application:
    ```bash
    go run ./cmd/app
    ```

### Option 2: Full Docker Stack (App + DB)
Run the entire stack (application and database) inside Docker containers.

1. Build and start the full stack:
    ```bash
    make up-dev
    ```

2. (Optional) Run migrations inside Docker: 
    ```bash
    make migrate-docker
    ```
3. Verify the service:
* Register a user: ```bash curl -X POST http://localhost:8081/auth/register -H "Content-Type: application/json" -d '{"email":"test@example.com","password":"password123"}'```
* Get token: ```bash curl -X POST http://localhost:8081/auth/login -H "Content-Type: application/json" -d '{"email":"test@example.com","password":"password123"}'```

## API Endpoints

All protected endpoints require a valid authentication token in the `Authorization: Bearer <token>` header.

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| `POST` | `/auth/register` | Register a new user. | ❌ No |
| `POST` | `/auth/login` | Login and get an access token. | ❌ No |
| `POST` | `/link` | Create a new short link. | ✅ Yes |
| `GET` | `/{hash}` | Redirect to the original URL (HTTP 302). | ❌ No |
| `PATCH` | `/link/{id}` | Update an existing link (e.g., change URL). | ✅ Yes |
| `DELETE` | `/link/{id}` | Delete a link. | ✅ Yes |
| `GET` | `/link` | List all links for the authenticated user. | ✅ Yes |
| `GET` | `/stat` | Get statistics for all links of the authenticated user. | ✅ Yes |


### Example: Create a Short Link
1. Register and login to get a token (as shown above).
2. Create a short link:
```bash 
curl -X POST http://localhost:8081/link \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{"url":"https://www.golang.org"}'
```
Response:
```bash
{
  "id": 1,
  "hash": "aBc12",
  "originalUrl": "https://www.golang.org",
  "createdAt": "2024-06-01T12:00:00Z"
}
```
## Configuration
All configuration is driven by environment variables:

* DATABASE_URL: PostgreSQL connection string.
* APP_HOST: Host to bind the server (default: 0.0.0.0).
* APP_PORT: Port to expose the API (default: 8081).
* SECRET: Secret key for signing tokens or other security features.
Use .env for local development; in production or CI, inject these variables directly into the container or runtime.

Generate a new secret with:
```bash
make gen-secret
```
## Makefile Targets
The Makefile provides convenient targets for common operations:
| Target | Description |
| --- | --- |
| `make up-dev` | Start the dev stack (Go app runs locally, PostgreSQL in Docker). |
| `make migrate` | Build and run database migrations locally using `ussMigrate`. |
| `make migrate-docker` | Run migrations inside a Docker container (uses embedded binary). |
| `make test-integration` | Spin up a fresh test stack and run integration tests; tears down automatically. |
| `make down` | Stop the development stack (containers remain, data is preserved). |
| `make clean` | ⚠️ Stop stack **and remove all data volumes** (all database data will be lost). Use with caution. |
| `make gen-secret` | Generate a new random secret key and save it to `.secrets`. |

[Up](#anchor)