# ms-ga-user

Gym Application - User Management Microservice.

This repository houses the core user management domain logic, built via Hexagonal Architecture, strictly conforming to API-First OpenAPI standards.

## Features
- Complete CRUD operations for core User parameters, Addresses, and precise Profiles.
- Relational mapping with GORM (PostgreSQL).
- Event streaming through pure-Go Kafka client (`segmentio/kafka-go`).
- OpenAPI specification with strictly generated generic types and Gin REST HTTP Handlers.
- JWT Authentication and Request Correlation ID Middlewares built-in.

## Architecture & Tech Stack
- **Language**: Go 1.23+
- **Framework**: Gin Gonic
- **Database**: PostgreSQL (Migrations using Flyway/pure-SQL)
- **Message Broker**: Kafka & Zookeeper
- **Code Generation**: `oapi-codegen`

## Setup & Running Locally

### Option 1: Docker Compose (Recommended)
This approach containerizes both the natively built Go API and the required backing infrastructure (PostgreSQL, Kafka).

```bash
docker-compose up -d --build
```
The application API will be seamlessly exposed at `http://localhost:8083`.

### Option 2: Run Go API Natively (Windows/Linux/Mac)
If you wish to compile and run the Go application directly on your host machine (e.g., for localized IDE step-through debugging) while keeping the database and Kafka backing instances active in Docker:

1. Spin up the infrastructure dependencies independently:
```bash
docker-compose up -d postgres zookeeper kafka
```

2. Start the application natively via Makefile bindings:
```bash
make run
# or fallback to: go run cmd/api/main.go
```
*Note: Because this project implements the pure-Go `segmentio/kafka-go` library, compiling this code requires absolutely **no strictly locked CGO or C/C++ GCC compiler dependencies**, allowing it to run flawlessly out-of-the-box on Windows!*

## API Definitions
Reference the Swagger parameters detailed directly internally inside `api/openapi.yaml` for complete, structured OpenAPI 3.0 specifications mapped straight to endpoints under `/gymapi/v1/users` and `/gymapi/v1/users/{id}/profile`.
