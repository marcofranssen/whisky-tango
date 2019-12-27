test: build
	@go test -v ./...

coverage:
	@echo Testing with code coverage
	@go test -v -covermode=atomic -coverprofile=coverage.out ./...

coverage-out: coverage
	@echo Coverage details
	@go tool cover -func=coverage.out

coverage-html: coverage
	@go tool cover -html=coverage.out

build: download
	@echo Building binaries
	@go build -o bin/wt ./cmd/wt

download:
	@echo Download go.mod dependencies
	@go mod download
