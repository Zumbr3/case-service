DATABASE_URL ?= pgx5://postgres:postgres@localhost:5433/case_service?sslmode=disable

.PHONY: migrate-up migrate-down migrate-version

## Applies all pending migrations
migrate-up:
	DATABASE_URL="$(DATABASE_URL)" go run ./cmd/migrate up

## Rolls back the last migration
migrate-down:
	DATABASE_URL="$(DATABASE_URL)" go run ./cmd/migrate down

## Prints the current migration version
migrate-version:
	DATABASE_URL="$(DATABASE_URL)" go run ./cmd/migrate version
