# Go User API

Production-ready Go backend project built with Fiber, PostgreSQL, SQLC, Zap, and clean architecture.

## Requirements
- Go 1.24+
- Docker & Docker Compose
- golang-migrate (for migrations)
- sqlc (for query generation)

## Quick Start

1. Start database and API using Docker Compose:
   ```bash
   docker-compose up -d
   ```

2. Run migrations (locally):
   ```bash
   migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/userdb?sslmode=disable" up
   ```

3. Run the application (locally):
   ```bash
   go run cmd/server/main.go
   ```
