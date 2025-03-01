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


.PHONY: build run-api  test lint clean run-api-env 
