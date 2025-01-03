-include .env
export $(shell sed 's/=.*//' .env)

DB_DSN := "$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)"

build:
	go build -o webserver ./cmd/webserver

test:
	go test ./... -v

docker-build:
	docker build -t ghcr.io/${GITHUB_REPOSITORY}/usergate:latest .

docker-push:
	docker push ghcr.io/${GITHUB_REPOSITORY}/usergate:latest

docker-run:
	docker-compose up --build

migrate-up:
	goose -dir ./db/migrations mysql $(DB_DSN) up

migrate-down:
	goose -dir ./db/migrations mysql $(DB_DSN) down

migrate-status:
	goose -dir ./db/migrations mysql $(DB_DSN) status
