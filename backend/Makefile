API_BIN=bin/user-api
API_SRC=cmd/user-api/main.go

ifneq (,$(wildcard .env))
    include .env
    export
endif

build: $(API_BIN) 

$(API_BIN): $(API_SRC)
	@echo "Building User API..."
	@mkdir -p bin
	@go build -o $@ $<


run-api: $(API_BIN)
    @echo "Running User API..."
    @./$(API_BIN)

test:
	@echo "Running tests..."
	@go test ./... -cover

lint:
	@echo "Running linter..."
	@golangci-lint run ./...

clean:
	@echo "Cleaning up..."
	@rm -rf bin

run-api-env:
	@echo "Running User API with .env..."
	@env $(shell cat .env | xargs) ./$(API_BIN)

run-api-local: $(API_BIN)
    @echo "Running User API (local dev, bypass telegram auth)..."
    @NETWORK=tcp4 ADDR=:3000 TG_BOT_TOKEN=dummy MONGODB_URI=mongodb://localhost:27017/myappdb MONGODB_NAME=myappdb MONGODB_COLLECTION_NAME=users BYPASS_TELEGRAM_AUTH=true ./$(API_BIN)

dev:
    @echo "Starting dev API with air (hot-reload)"
    @AIR_CONFIG=.air.toml air || echo "Install air: go install github.com/cosmtrek/air@latest"

mongo-up:
    @echo "Starting local MongoDB (Docker)"
    @docker run --name tma-mongo -p 27017:27017 -v $(PWD)/../devops/mongo-data:/data/db -d mongo:7 || true

mongo-down:
    @echo "Stopping local MongoDB"
    @docker rm -f tma-mongo || true


.PHONY: build run-api  test lint clean run-api-env 
