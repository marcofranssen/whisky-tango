build: download
	@echo Building binaries
	@go build -o bin/wt ./cmd/wt

download:
	@echo Download go.mod dependencies
	@go mod download
