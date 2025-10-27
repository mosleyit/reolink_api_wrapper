# Reolink API Wrapper - Makefile
# Comprehensive build, test, and development targets

.PHONY: help
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Variables
GO := go
GOFLAGS := -v
GOTEST := $(GO) test
GOBUILD := $(GO) build
GOCLEAN := $(GO) clean
GOMOD := $(GO) mod
GOFMT := gofmt
GOLINT := golangci-lint

# Directories
EXAMPLES_DIR := examples
BASIC_EXAMPLE := $(EXAMPLES_DIR)/basic
DEBUG_EXAMPLE := $(EXAMPLES_DIR)/debug_test
HARDWARE_EXAMPLE := $(EXAMPLES_DIR)/hardware_test
SCRIPTS_DIR := scripts

# Binary names
BASIC_BIN := $(BASIC_EXAMPLE)/basic
DEBUG_BIN := $(DEBUG_EXAMPLE)/debug_test
HARDWARE_BIN := $(HARDWARE_EXAMPLE)/hardware_test

# Coverage
COVERAGE_FILE := coverage.txt
COVERAGE_HTML := coverage.html

# Colors for output
COLOR_RESET := \033[0m
COLOR_BOLD := \033[1m
COLOR_GREEN := \033[32m
COLOR_YELLOW := \033[33m
COLOR_BLUE := \033[34m

##@ Development

.PHONY: all
all: clean deps fmt lint test build ## Run all checks and build everything

.PHONY: deps
deps: ## Download and verify dependencies
	@echo "$(COLOR_BLUE)Downloading dependencies...$(COLOR_RESET)"
	$(GOMOD) download
	$(GOMOD) verify
	$(GOMOD) tidy

.PHONY: fmt
fmt: ## Format all Go code
	@echo "$(COLOR_BLUE)Formatting code...$(COLOR_RESET)"
	$(GOFMT) -s -w .
	@if command -v goimports >/dev/null 2>&1; then \
		echo "Running goimports..."; \
		goimports -w .; \
	fi

.PHONY: fmt-check
fmt-check: ## Check if code is formatted
	@echo "$(COLOR_BLUE)Checking code formatting...$(COLOR_RESET)"
	@UNFORMATTED=$$($(GOFMT) -l .); \
	if [ -n "$$UNFORMATTED" ]; then \
		echo "$(COLOR_YELLOW)The following files are not formatted:$(COLOR_RESET)"; \
		echo "$$UNFORMATTED"; \
		exit 1; \
	fi
	@echo "$(COLOR_GREEN)✓ All files are properly formatted$(COLOR_RESET)"

.PHONY: lint
lint: ## Run linter
	@echo "$(COLOR_BLUE)Running linter...$(COLOR_RESET)"
	$(GOLINT) run ./...
	@echo "$(COLOR_GREEN)✓ Linting passed$(COLOR_RESET)"

.PHONY: lint-fix
lint-fix: ## Run linter with auto-fix
	@echo "$(COLOR_BLUE)Running linter with auto-fix...$(COLOR_RESET)"
	$(GOLINT) run --fix ./...

##@ Testing

.PHONY: test
test: ## Run all tests
	@echo "$(COLOR_BLUE)Running tests...$(COLOR_RESET)"
	$(GOTEST) $(GOFLAGS) ./...
	@echo "$(COLOR_GREEN)✓ All tests passed$(COLOR_RESET)"

.PHONY: test-short
test-short: ## Run tests with -short flag
	@echo "$(COLOR_BLUE)Running short tests...$(COLOR_RESET)"
	$(GOTEST) -short ./...

.PHONY: test-verbose
test-verbose: ## Run tests with verbose output
	@echo "$(COLOR_BLUE)Running tests (verbose)...$(COLOR_RESET)"
	$(GOTEST) -v ./...

.PHONY: test-race
test-race: ## Run tests with race detector
	@echo "$(COLOR_BLUE)Running tests with race detector...$(COLOR_RESET)"
	$(GOTEST) -race ./...

.PHONY: test-coverage
test-coverage: ## Run tests with coverage
	@echo "$(COLOR_BLUE)Running tests with coverage...$(COLOR_RESET)"
	$(GOTEST) -coverprofile=$(COVERAGE_FILE) -covermode=atomic ./...
	@echo "$(COLOR_GREEN)✓ Coverage report generated: $(COVERAGE_FILE)$(COLOR_RESET)"

.PHONY: coverage-html
coverage-html: test-coverage ## Generate HTML coverage report
	@echo "$(COLOR_BLUE)Generating HTML coverage report...$(COLOR_RESET)"
	$(GO) tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	@echo "$(COLOR_GREEN)✓ HTML coverage report: $(COVERAGE_HTML)$(COLOR_RESET)"
	@if command -v open >/dev/null 2>&1; then \
		open $(COVERAGE_HTML); \
	elif command -v xdg-open >/dev/null 2>&1; then \
		xdg-open $(COVERAGE_HTML); \
	fi

.PHONY: coverage-report
coverage-report: test-coverage ## Show coverage report in terminal
	@echo "$(COLOR_BLUE)Coverage report:$(COLOR_RESET)"
	$(GO) tool cover -func=$(COVERAGE_FILE)

.PHONY: bench
bench: ## Run benchmarks
	@echo "$(COLOR_BLUE)Running benchmarks...$(COLOR_RESET)"
	$(GOTEST) -bench=. -benchmem ./...

##@ Building

.PHONY: build
build: build-examples ## Build all examples

.PHONY: build-examples
build-examples: build-basic build-debug build-hardware ## Build all example binaries

.PHONY: build-basic
build-basic: ## Build basic example
	@echo "$(COLOR_BLUE)Building basic example...$(COLOR_RESET)"
	cd $(BASIC_EXAMPLE) && $(GOBUILD) $(GOFLAGS) -o basic main.go
	@echo "$(COLOR_GREEN)✓ Built: $(BASIC_BIN)$(COLOR_RESET)"

.PHONY: build-debug
build-debug: ## Build debug_test example
	@echo "$(COLOR_BLUE)Building debug_test example...$(COLOR_RESET)"
	cd $(DEBUG_EXAMPLE) && $(GOBUILD) $(GOFLAGS) -o debug_test main.go
	@echo "$(COLOR_GREEN)✓ Built: $(DEBUG_BIN)$(COLOR_RESET)"

.PHONY: build-hardware
build-hardware: ## Build hardware_test example
	@echo "$(COLOR_BLUE)Building hardware_test example...$(COLOR_RESET)"
	cd $(HARDWARE_EXAMPLE) && $(GOBUILD) $(GOFLAGS) -o hardware_test main.go
	@echo "$(COLOR_GREEN)✓ Built: $(HARDWARE_BIN)$(COLOR_RESET)"

##@ Running

.PHONY: run-basic
run-basic: ## Run basic example (requires REOLINK_* env vars)
	@echo "$(COLOR_BLUE)Running basic example...$(COLOR_RESET)"
	@if [ -z "$$REOLINK_HOST" ]; then \
		echo "$(COLOR_YELLOW)Warning: REOLINK_HOST not set$(COLOR_RESET)"; \
		echo "Set environment variables: REOLINK_HOST, REOLINK_USERNAME, REOLINK_PASSWORD"; \
		exit 1; \
	fi
	cd $(BASIC_EXAMPLE) && $(GO) run main.go

.PHONY: run-debug
run-debug: ## Run debug_test example (requires REOLINK_* env vars)
	@echo "$(COLOR_BLUE)Running debug_test example...$(COLOR_RESET)"
	@if [ -z "$$REOLINK_HOST" ]; then \
		echo "$(COLOR_YELLOW)Warning: REOLINK_HOST not set$(COLOR_RESET)"; \
		echo "Set environment variables: REOLINK_HOST, REOLINK_USERNAME, REOLINK_PASSWORD"; \
		exit 1; \
	fi
	cd $(DEBUG_EXAMPLE) && $(GO) run main.go

.PHONY: run-hardware
run-hardware: ## Run hardware_test example (requires REOLINK_* env vars)
	@echo "$(COLOR_BLUE)Running hardware_test example...$(COLOR_RESET)"
	@if [ -z "$$REOLINK_HOST" ]; then \
		echo "$(COLOR_YELLOW)Warning: REOLINK_HOST not set$(COLOR_RESET)"; \
		echo "Set environment variables: REOLINK_HOST, REOLINK_USERNAME, REOLINK_PASSWORD"; \
		exit 1; \
	fi
	cd $(HARDWARE_EXAMPLE) && $(GO) run main.go

##@ Cleaning

.PHONY: clean
clean: clean-binaries clean-coverage ## Clean all generated files

.PHONY: clean-binaries
clean-binaries: ## Remove built binaries
	@echo "$(COLOR_BLUE)Cleaning binaries...$(COLOR_RESET)"
	@rm -f $(BASIC_BIN) $(DEBUG_BIN) $(HARDWARE_BIN)
	$(GOCLEAN)
	@echo "$(COLOR_GREEN)✓ Binaries cleaned$(COLOR_RESET)"

.PHONY: clean-coverage
clean-coverage: ## Remove coverage files
	@echo "$(COLOR_BLUE)Cleaning coverage files...$(COLOR_RESET)"
	@rm -f $(COVERAGE_FILE) $(COVERAGE_HTML)
	@echo "$(COLOR_GREEN)✓ Coverage files cleaned$(COLOR_RESET)"

.PHONY: clean-all
clean-all: clean ## Clean everything including dependencies
	@echo "$(COLOR_BLUE)Cleaning all...$(COLOR_RESET)"
	$(GOMOD) clean -cache
	@echo "$(COLOR_GREEN)✓ Everything cleaned$(COLOR_RESET)"

##@ Git Hooks

.PHONY: install-hooks
install-hooks: ## Install git pre-commit hooks
	@echo "$(COLOR_BLUE)Installing git hooks...$(COLOR_RESET)"
	@bash $(SCRIPTS_DIR)/install-hooks.sh
	@echo "$(COLOR_GREEN)✓ Git hooks installed$(COLOR_RESET)"

##@ Documentation

.PHONY: docs
docs: ## Generate documentation
	@echo "$(COLOR_BLUE)Generating documentation...$(COLOR_RESET)"
	@if command -v gomarkdoc >/dev/null 2>&1; then \
		gomarkdoc --output docs/SDK_API.md ./...; \
		echo "$(COLOR_GREEN)✓ Documentation generated: docs/SDK_API.md$(COLOR_RESET)"; \
	else \
		echo "$(COLOR_YELLOW)gomarkdoc not installed. Install with:$(COLOR_RESET)"; \
		echo "  go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest"; \
	fi

.PHONY: docs-serve
docs-serve: ## Serve documentation locally
	@echo "$(COLOR_BLUE)Serving documentation at http://localhost:8000$(COLOR_RESET)"
	@cd docs && python3 -m http.server 8000

##@ Verification

.PHONY: verify
verify: fmt-check lint test ## Run all verification checks (CI-friendly)

.PHONY: pre-commit
pre-commit: fmt lint test-short ## Run pre-commit checks (fast)

.PHONY: ci
ci: deps verify test-coverage ## Run full CI pipeline

##@ Tools

.PHONY: install-tools
install-tools: ## Install development tools
	@echo "$(COLOR_BLUE)Installing development tools...$(COLOR_RESET)"
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(GO) install golang.org/x/tools/cmd/goimports@latest
	$(GO) install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
	@echo "$(COLOR_GREEN)✓ Tools installed$(COLOR_RESET)"

.PHONY: version
version: ## Show version information
	@echo "$(COLOR_BOLD)Reolink API Wrapper$(COLOR_RESET)"
	@grep -E '^const Version' version.go || echo "Version: unknown"
	@echo ""
	@echo "Go version: $$($(GO) version)"
	@echo "Module: $$(grep '^module' go.mod | awk '{print $$2}')"

