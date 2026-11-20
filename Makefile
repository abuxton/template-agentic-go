SHELL := /bin/sh

APP := agent
BIN_DIR := bin

.PHONY: all build test test-race vet fmt lint run clean

all: test build

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP) ./cmd/agent

test:
	go test ./...

test-race:
	go test -race ./...

vet:
	go vet ./...

fmt:
	test -z "$$(gofmt -l .)"

lint:
	golangci-lint run --timeout 5m

run:
	go run ./cmd/agent "$${PROMPT:?set PROMPT='...'}"

clean:
	rm -rf $(BIN_DIR)
