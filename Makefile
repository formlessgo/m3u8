APP_NAME := M3U8
BIN_NAME := m3u8
APP_VERSION = 0.1.0

SHELL := env /bin/bash

help:
	@echo "Usage: make [command]"
	@echo ""
	@echo "    build: build the application"
	@echo ""

build:
	go build -ldflags "-s -w" -o ./dist/$(BIN_NAME) ./main.go
	@echo "Build success"

build-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o ./dist/$(BIN_NAME)-darwin-amd64 ./main.go

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./dist/$(BIN_NAME)-linux-amd64 ./main.go

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./dist/$(BIN_NAME)-windows-amd64.exe ./main.go

dev:
	go build -o ./dist/$(BIN_NAME) ./main.go
	./dist/$(BIN_NAME)

fmt:
	go fmt ./...

lint:
	go vet ./...

release: build-darwin build-linux build-windows
	cd dist; \
	zip ./$(BIN_NAME)-darwin-amd64.zip $(BIN_NAME)-darwin-amd64; \
	zip ./$(BIN_NAME)-linux-amd64.zip $(BIN_NAME)-linux-amd64; \
	zip ./$(BIN_NAME)-windows-amd64.zip $(BIN_NAME)-windows-amd64.exe
	@echo "Release success"

start:
	@./dist/$(BIN_NAME)
