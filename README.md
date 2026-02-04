# StockWatch

StockWatch is a microservice-based system for monitoring stock prices, evaluating alerts, and notifying users. The repository contains backend services (written in Go), a React frontend, database initialization SQL, and helper tooling for running everything with Docker.

- Microservices hosted on AWS ECS
- Database is AWS RDS PostgreSQL
- Front-end is stored in AWS S3

## Repository layout

- `AlertsPoller/` - polls prices from external sources and stores them.
- `AlertsEvaluator/` - evaluates saved prices against user alerts and triggers notifications.
- `EmailSender/` - sends email notifications when alerts fire.
- `WatchListAPI/` - API service that manages watchlists, alerts, and serves the frontend data.
- `StockwatchUI/` - React + Vite frontend.
- `Database/` - SQL files for initializing the database: `initDB.sql`, `initTables.sql`.
- `compose.yaml` - Docker Compose file to bring up the full stack locally.

## Architecture

The system components communicate using RabbitMQ (message queue) and a shared relational database. High-level flow:

1. `AlertsPoller` fetches price data and stores it in the database.
2. `AlertsEvaluator` consumes price updates, checks alert conditions, and publishes alert events.
3. `EmailSender` subscribes to alert events and sends emails to users.
4. `WatchListAPI` provides REST endpoints for managing watchlists and alerts and is consumed by the frontend.
5. `StockwatchUI` is a React app that allows users to log in and manage watchlists and alerts.

## Prerequisites

- Docker & Docker Compose
- Go (for running services locally without Docker)
- RabbitMQ for event brokering
- Node.js + npm (for frontend development)

## Quick start (recommended)

Bring up the entire stack using Docker Compose (root of repository):

```bash
docker compose -f compose.yaml up --build
```

This will build and start the services, RabbitMQ, and any database container defined in `compose.yaml`.

## Running services individually (development)

Database:
- Apply `Database/initDB.sql` and `Database/initTables.sql` to initialize the DB schema and seed data (adjust to your DB engine).

Go services (run from repo root or each service folder):

```bash
cd AlertsPoller
go run Main.go

cd ../AlertsEvaluator
go run Main.go

cd ../WatchListAPI
go run Server.go

cd ../EmailSender
go run Main.go
```

Adjust service configuration via environment variables or the Compose file as needed (DB connection, RabbitMQ URL, credentials).

Frontend (development):

```bash
cd StockwatchUI
npm install
npm run dev
```

The frontend uses Vite; dev server will be available at the URL printed by Vite (typically `http://localhost:5173`).

## Important files and configuration

- `compose.yaml` — orchestration for local development.
- `Database/initDB.sql`, `Database/initTables.sql` — DB initialization scripts.
- `StockwatchUI/package.json` — frontend scripts (`dev`, `build`, `preview`).
- Look in each service folder for `Dockerfile`, `go.mod`, and `Main.go`/`Server.go`.

## Environment variables
Check the global `.env_example` to see common variables used by services (check each service's source for exact names):

Environment file passed in `compose.yaml` testing, but should be included when creating the task on AWS ECS.

Front-end in `StockwatchUI` also has its own `.env_example`

- For local Go services, run them in a terminal to see stdout logs.

## Development notes

- Each microservice is a small Go module; use `go build` / `go run` in the service directory.
- The frontend is a Vite React app; use `npm run dev` for hot-reload during development and `npm run build` to produce static assets.
