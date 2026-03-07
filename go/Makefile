.PHONY: build run clean help fmt lint

help:
	@echo "Tiscogs - Discogs CLI Explorer"
	@echo ""
	@echo "Available targets:"
	@echo "  build    - Build the tiscogs binary"
	@echo "  run      - Run tiscogs (requires DISCOGS_TOKEN env var)"
	@echo "  clean    - Remove build artifacts"
	@echo "  fmt      - Format code with gofmt"
	@echo "  lint     - Run golangci-lint"
	@echo "  deps     - Download dependencies"
	@echo "  help     - Show this help message"

build:
	go build -o tiscogs

run: build
	./tiscogs

clean:
	rm -f tiscogs

fmt:
	go fmt ./...

lint:
	golangci-lint run ./...

deps:
	go mod download
	go mod tidy
