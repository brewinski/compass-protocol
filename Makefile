# Makefile for Compass Protocol MCP Server

# Binary name
BINARY_NAME=compass-protocol

# Version (can be overridden)
VERSION?=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")

# Build flags
LDFLAGS=-ldflags "-X main.Version=$(VERSION)"

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt

.PHONY: all build clean test test-coverage run deps fmt lint help

all: test build ## Run tests and build the binary

build: ## Build the binary
	@echo "Building $(BINARY_NAME) version $(VERSION)..."
	@$(GOBUILD) $(LDFLAGS) -o $(BINARY_NAME) -v

clean: ## Remove build artifacts
	@echo "Cleaning..."
	@$(GOCLEAN)
	@rm -f $(BINARY_NAME)
	@rm -rf dist/
	@rm -f coverage.txt coverage.html

test: ## Run tests
	@echo "Running tests..."
	@$(GOTEST) -v ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@$(GOTEST) -v -coverprofile=coverage.txt -covermode=atomic ./...
	@$(GOCMD) tool cover -html=coverage.txt -o coverage.html
	@echo "Coverage report generated: coverage.html"

run: build ## Build and run the server
	@echo "Running $(BINARY_NAME)..."
	@./$(BINARY_NAME)

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@$(GOMOD) download
	@$(GOMOD) tidy

fmt: ## Format Go code
	@echo "Formatting code..."
	@$(GOFMT) ./...

lint: ## Run golangci-lint (requires golangci-lint to be installed)
	@echo "Running linter..."
	@which golangci-lint > /dev/null || (echo "golangci-lint not found. Install it from https://golangci-lint.run/usage/install/" && exit 1)
	@golangci-lint run

install: ## Install the binary to GOPATH/bin
	@echo "Installing $(BINARY_NAME)..."
	@$(GOCMD) install $(LDFLAGS)

help: ## Display this help message
	@echo "Compass Protocol MCP Server - Makefile commands:"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo ""
