.PHONY: help test test-integration coverage lint fmt vet clean build examples

# Default target
help:
	@echo "Reolink Camera API - Go SDK"
	@echo ""
	@echo "Available targets:"
	@echo "  make test              - Run unit tests"
	@echo "  make test-integration  - Run integration tests (requires camera)"
	@echo "  make coverage          - Run tests with coverage report"
	@echo "  make lint              - Run golangci-lint"
	@echo "  make fmt               - Format code with gofmt and goimports"
	@echo "  make vet               - Run go vet"
	@echo "  make clean             - Clean build artifacts"
	@echo "  make build             - Build all packages"
	@echo "  make examples          - Build all examples"
	@echo "  make check             - Run all checks (fmt, vet, lint, test)"

# Run unit tests
test:
	@echo "Running unit tests..."
	go test -v -race ./...

# Run integration tests (requires real camera)
test-integration:
	@echo "Running integration tests..."
	@echo "Make sure REOLINK_HOST, REOLINK_USERNAME, and REOLINK_PASSWORD are set"
	go test -v -race -tags=integration ./...

# Run tests with coverage
coverage:
	@echo "Running tests with coverage..."
	go test -v -race -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run linter
lint:
	@echo "Running golangci-lint..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	elif [ -f $(HOME)/go/bin/golangci-lint ]; then \
		$(HOME)/go/bin/golangci-lint run ./...; \
	else \
		echo "golangci-lint not found. Run 'make install-tools' first."; \
		exit 1; \
	fi

# Format code
fmt:
	@echo "Formatting code..."
	gofmt -s -w .
	@if command -v goimports >/dev/null 2>&1; then \
		goimports -w .; \
	elif [ -f $(HOME)/go/bin/goimports ]; then \
		$(HOME)/go/bin/goimports -w .; \
	else \
		echo "goimports not found. Run 'make install-tools' first."; \
		exit 1; \
	fi

# Run go vet
vet:
	@echo "Running go vet..."
	go vet ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -f coverage.out coverage.html
	go clean ./...

# Build all packages
build:
	@echo "Building..."
	go build ./...

# Build examples
examples:
	@echo "Building examples..."
	@for dir in examples/*/; do \
		if [ -f "$$dir/main.go" ]; then \
			echo "Building $$dir..."; \
			(cd "$$dir" && go build -o "$$(basename $$dir)" .); \
		fi \
	done

# Run all checks
check: fmt vet lint test
	@echo "All checks passed!"

# Install development tools
install-tools:
	@echo "Installing development tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest
	@echo "Tools installed!"

