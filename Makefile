# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=server
BINARY_UNIX=$(BINARY_NAME)_unix

# Main package path
MAIN_PATH=./cmd/server

.PHONY: all build clean test coverage deps run dev migrate-up migrate-down docker-build docker-run

all: test build

build:
	$(GOBUILD) -o bin/$(BINARY_NAME) -v $(MAIN_PATH)

clean:
	$(GOCLEAN)
	rm -f bin/$(BINARY_NAME)
	rm -f bin/$(BINARY_UNIX)

test:
	$(GOTEST) -v ./...

coverage:
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

deps:
	$(GOMOD) download
	$(GOMOD) tidy

run:
	$(GOCMD) run $(MAIN_PATH)/main.go

dev:
	air -c .air.toml

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_UNIX) -v $(MAIN_PATH)

# Database migrations
migrate-up:
	psql -d $(DATABASE_NAME) -f migrations/001_create_users_table.sql

migrate-down:
	psql -d $(DATABASE_NAME) -c "DROP TABLE IF EXISTS users;"

# Docker commands
docker-build:
	docker build -t clean-arch-project .

docker-run:
	docker run -p 8080:8080 clean-arch-project

# Format code
fmt:
	$(GOCMD) fmt ./...

# Lint code
lint:
	golangci-lint run

# Security check
security:
	gosec ./...

# Generate swagger docs
swagger:
	swag init -g cmd/server/main.go

# Install dev tools
install-tools:
	$(GOGET) github.com/cosmtrek/air@latest
	$(GOGET) github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(GOGET) github.com/securecodewarrior/gosec/v2/cmd/gosec@latest
	$(GOGET) github.com/swaggo/swag/cmd/swag@latest