# Makefile for llm-wiki-go
#
# Convenience wrapper around the Go toolchain and the wikilint CLI.
# Run `make help` to list the available targets.

GO       ?= go
GOFMT    ?= gofmt
BIN_DIR  := bin
WIKILINT := $(BIN_DIR)/wikilint
WIKI_DIR := ./wiki

.DEFAULT_GOAL := help

.PHONY: help build test fmt fmt-check vet lint check tidy clean install

help: ## Show this help
	@awk 'BEGIN {FS = ":.*?## "; printf "Usage: make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?## / { printf "  \033[36m%-11s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

build: ## Build the wikilint binary into bin/
	@mkdir -p $(BIN_DIR)
	$(GO) build -o $(WIKILINT) ./cmd/wikilint

test: ## Run the Go test suite
	$(GO) test ./...

fmt: ## Format all Go files in-place
	$(GOFMT) -w .

fmt-check: ## Fail if any Go file needs gofmt
	@out=$$($(GOFMT) -l .); \
	if [ -n "$$out" ]; then \
		echo "gofmt would reformat:"; \
		echo "$$out"; \
		exit 1; \
	fi

vet: ## Run `go vet` on all packages
	$(GO) vet ./...

lint: ## Run wikilint against ./wiki
	$(GO) run ./cmd/wikilint -wiki $(WIKI_DIR)

check: fmt-check vet test lint ## Run the full pre-commit verification chain

tidy: ## Run `go mod tidy`
	$(GO) mod tidy

clean: ## Remove build artifacts
	rm -rf $(BIN_DIR)

install: ## Install wikilint into $(GOBIN) or $(GOPATH)/bin
	$(GO) install ./cmd/wikilint
