# Load environment variables from .env file
include .env
export $(shell sed 's/=.*//' .env)

# Define database DSN dynamically using environment variables
DB_DSN := "$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)"

build:
	go build -o webserver ./cmd/webserver

test:
	go test ./... -v

migrate-up:
	goose -dir ./db/migrations mysql $(DB_DSN) up

migrate-down:
	goose -dir ./db/migrations mysql $(DB_DSN) down

migrate-status:
	goose -dir ./db/migrations mysql $(DB_DSN) status
