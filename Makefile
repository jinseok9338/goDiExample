# Database configuration
DB_HOST=localhost
DB_PORT=5432
DB_NAME=postgres
DB_USER=postgres
DB_PASSWORD=mysecretpassword

# Construct the database URL
DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

# Migration directory
MIGRATION_DIR=./db/migration

# golang-migrate binary (assuming it's installed and in your PATH)
MIGRATE=migrate

# Default target
.DEFAULT_GOAL := help

.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make migrate-up   - Run all up migrations"
	@echo "  make migrate-down - Rollback all migrations"
	@echo "  make migrate-create name=migration_name - Create a new migration"
	@echo "  make swagger-gen - Generate swagger docs"
	@echo "  make wire-gen - Generate Wire generation"

.PHONY: migrate-up
migrate-up:
	$(MIGRATE) -path $(MIGRATION_DIR) -database "$(DB_URL)" up

.PHONY: migrate-down
migrate-down:
	$(MIGRATE) -path $(MIGRATION_DIR) -database "$(DB_URL)" down

.PHONY: migrate-create
migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Error: 'name' parameter is required. Usage: make migrate-create name=migration_name"; \
		exit 1; \
	fi
	$(MIGRATE) create -ext sql -dir $(MIGRATION_DIR) -seq $(name)

.PHONY: swagger-gen
swagger-gen:
	go run github.com/swaggo/swag/cmd/swag@latest init -g ./main.go -o ./docs

.PHONY: wire-gen
wire-gen:
	go run github.com/google/wire/cmd/wire@latest