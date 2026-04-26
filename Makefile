# Makefile for llm-wiki-go
#
# Convenience wrapper around the Go toolchain and the wikilint CLI.
# Run `make help` to list the available targets.

GO       ?= go
BIN_DIR  := bin
WIKILINT := $(BIN_DIR)/wikilint
WIKI_DIR := ./wiki
VERSION  ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT   ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_DATE ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS  := -s -w
GOLANGCI_LINT_VERSION := v2.11.4
GOFUMPT_VERSION       := v0.9.2

.DEFAULT_GOAL := help

.PHONY: all help build build-prod build-linux-amd64 build-darwin-arm64 build-all-platforms \
        test test-race coverage coverage-html fmt fmt-check vet lint lint-go lint-wiki check deps tidy clean install-tools \
        install setup

help: ## Show this help
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z0-9_-]+:.*##/ {printf "  \033[36m%-22s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

all: build ## Build the default local/dev binary

build: ## Build fast local/dev wikilint binary into bin/
	@mkdir -p $(BIN_DIR)
	$(GO) build -o $(WIKILINT) ./cmd/wikilint

build-prod: ## Build optimized host production wikilint binary into bin/
	@mkdir -p $(BIN_DIR)
	@echo "Building wikilint $(VERSION) ($(COMMIT), $(BUILD_DATE))..."
	$(GO) build -trimpath -ldflags="$(LDFLAGS)" -o $(WIKILINT) ./cmd/wikilint

build-linux-amd64: ## Build optimized static Linux AMD64 production wikilint binary
	@mkdir -p $(BIN_DIR)
	@echo "Building wikilint $(VERSION) for Linux AMD64 ($(COMMIT), $(BUILD_DATE))..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOAMD64=v3 \
		$(GO) build -trimpath -ldflags="$(LDFLAGS)" \
		-o $(BIN_DIR)/wikilint-linux-amd64 ./cmd/wikilint

build-darwin-arm64: ## Build optimized Darwin ARM64 production wikilint binary
	@mkdir -p $(BIN_DIR)
	@echo "Building wikilint $(VERSION) for Darwin ARM64 ($(COMMIT), $(BUILD_DATE))..."
	GOOS=darwin GOARCH=arm64 \
		$(GO) build -trimpath -ldflags="$(LDFLAGS)" \
		-o $(BIN_DIR)/wikilint-darwin-arm64 ./cmd/wikilint

build-all-platforms: build-linux-amd64 build-darwin-arm64 ## Build all production platform binaries

test: ## Run the Go test suite
	$(GO) test ./...

test-race: ## Run tests with race detector
	$(GO) test -race ./...

coverage: ## Run tests with coverage summary
	$(GO) test -cover ./...

coverage-html: ## Generate HTML coverage report
	$(GO) test -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

fmt: ## Format all Go files in-place
	gofumpt -w .

fmt-check: ## Fail if gofumpt would reformat files
	@out=$$(gofumpt -l .); \
	if [ -n "$$out" ]; then \
		echo "gofumpt would reformat:"; \
		echo "$$out"; \
		exit 1; \
	fi

vet: ## Run `go vet` on all packages
	$(GO) vet ./...

lint-go: ## Run golangci-lint
	golangci-lint run ./...

lint-wiki: ## Run wikilint against ./wiki
	$(GO) run ./cmd/wikilint -wiki $(WIKI_DIR)

lint: lint-go lint-wiki ## Run all linters

check: fmt-check vet lint test ## Run the full local quality gate

deps: ## Download Go module dependencies
	$(GO) mod download

tidy: ## Run `go mod tidy`
	$(GO) mod tidy

clean: ## Remove build artifacts
	rm -rf $(BIN_DIR) coverage.out coverage.html

install-tools: ## Install pinned developer tools
	$(GO) install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)
	$(GO) install mvdan.cc/gofumpt@$(GOFUMPT_VERSION)

install: ## Install wikilint into $(GOBIN) or $(GOPATH)/bin
	$(GO) install ./cmd/wikilint

setup: ## Create raw/ and wiki/ directory structure
	@mkdir -p raw wiki/entities wiki/topics wiki/sources
	@test -f wiki/index.md || printf '# Wiki Index\n\n## Entities\n\n## Topics\n\n## Log\n\n- [Change Log](log.md)\n' > wiki/index.md
	@test -f wiki/log.md  || printf '# Change Log\n' > wiki/log.md
	@echo "Directory structure ready."
