.PHONY: build test clean install-deps build-all run-tests
.DEFAULT_GOAL := build

# Build both binaries
build:
	@echo "Building note-compiler..."
	@go build -o bin/note-compiler ./cmd/note-compiler
	@echo "note-compiler built successfully!"

# Build for development with verbose output
build-dev:
	@echo "Building with race detection and verbose output..."
	@go build -race -v -o bin/note-compiler ./cmd/note-compiler

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html

# Install dependencies
install-deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

# Cross-platform builds (for manual testing)
build-all:
	@echo "Building for all platforms..."
	@mkdir -p dist
	@GOOS=darwin GOARCH=amd64 go build -o dist/note-compiler-darwin-amd64 ./cmd/note-compiler
	@GOOS=darwin GOARCH=arm64 go build -o dist/note-compiler-darwin-arm64 ./cmd/note-compiler
	@GOOS=linux GOARCH=amd64 go build -o dist/note-compiler-linux-amd64 ./cmd/note-compiler
	@GOOS=windows GOARCH=amd64 go build -o dist/note-compiler-windows-amd64.exe ./cmd/note-compiler
	@echo "Binaries built in dist/"

# Run linter (if you install golangci-lint)
lint:
	@echo "Running linter..."
	@golangci-lint run

# Install golangci-lint if not present
install-lint:
	@which golangci-lint > /dev/null || (echo "Installing golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)

# Run all checks before pushing
pre-push: install-lint fmt lint test
	@echo "All pre-push checks passed ✅"

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Generate mocks (if you use mockgen)
generate:
	@echo "Generating code..."
	@go generate ./...

# Local install (builds and copies to GOPATH/bin)
install:
	@echo "Installing locally..."
	@go install ./cmd/note-compiler
	@go install ./cmd/obsidian-backup

# Show help
help:
	@echo "Available commands:"
	@echo "  build        - Build both binaries"
	@echo "  build-dev    - Build with race detection"
	@echo "  test         - Run tests"
	@echo "  test-coverage- Run tests with coverage report"
	@echo "  clean        - Clean build artifacts"
	@echo "  install-deps - Download and tidy dependencies"
	@echo "  build-all    - Cross-platform builds"
	@echo "  lint         - Run linter"
	@echo "  install-lint - Install golangci-lint"
	@echo "  fmt          - Format code"
	@echo "  pre-push     - Run all pre-push checks"
	@echo "  install      - Install binaries locally"
	@echo "  help         - Show this help" 