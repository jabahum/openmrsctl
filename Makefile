# =========================================================
# 🧰 Makefile for openmrsctl
# =========================================================
# Provides targets for:
#  - Building, testing, and linting
#  - Cross-platform release builds
#  - Docker image creation
#  - Tool verification and installation
# =========================================================

# --- Configuration ---

BINARY        := openmrsctl
CMD_DIR       := .
MAIN_FILE     := main.go
DIST_DIR      := dist
MODULE_PATH   := github.com/jabahum/openmrsctl

# Load environment variables if .env exists
ifneq (,$(wildcard .env))
	include .env
	export $(shell sed 's/=.*//' .env)
endif

# Version info from Git
COMMIT        := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
VERSION       := $(shell git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
DATE          := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

# Inject metadata into the binary
LDFLAGS       := -X '$(MODULE_PATH)/pkg/version.Version=$(VERSION)' \
                  -X '$(MODULE_PATH)/pkg/version.GitCommit=$(COMMIT)' \
                  -X '$(MODULE_PATH)/pkg/version.BuildDate=$(DATE)' \
                  -s -w

# Build environment
OS            := $(shell go env GOOS)
ARCH          := $(shell go env GOARCH)

# Colors
CYAN=\033[36m
GREEN=\033[32m
YELLOW=\033[33m
RED=\033[31m
RESET=\033[0m

# =========================================================
# 📦 Targets
# =========================================================
.PHONY:  build run fmt lint test tidy clean release release-all docker install verify-tools help


# =========================================================
# 🏗️ Building and Running
# =========================================================

build: ## Build the CLI for the current platform
	@echo "$(CYAN)🔨 Building $(BINARY) ($(VERSION)-$(COMMIT)) for $(OS)/$(ARCH)...$(RESET)"
	go build -ldflags "$(LDFLAGS)" -o $(BINARY) $(MAIN_FILE)
	@echo "$(GREEN)✅ Build complete: ./$(BINARY)$(RESET)"

run: build ## Run the CLI locally (pass args via ARGS="...")
	@echo "$(CYAN)▶️ Running $(BINARY) with args: $(ARGS)$(RESET)"
	@./$(BINARY) $(ARGS)

install: build ## Install binary to /usr/local/bin
	@echo "$(CYAN)📦 Installing $(BINARY) to /usr/local/bin...$(RESET)"
	sudo mv $(BINARY) /usr/local/bin/
	@echo "$(GREEN)✅ Installed successfully$(RESET)"

# =========================================================
# 💻 Development Utilities
# =========================================================

fmt: ## Format code
	@echo "$(CYAN)🖌 Formatting Go code...$(RESET)"
	go fmt ./...

lint: ## Lint code using golangci-lint
	@echo "$(CYAN)🔍 Running linter...$(RESET)"
	@$(MAKE) verify-tools
	golangci-lint run

test: ## Run all tests
	@echo "$(CYAN)🧪 Running tests with coverage...$(RESET)"
	go test ./... -v -cover

tidy: ## Clean and tidy Go modules
	@echo "$(CYAN)📦 Tidying Go modules...$(RESET)"
	go mod tidy

verify-tools: ## Verify required tools are installed
	@echo "$(CYAN)🧰 Verifying required tools...$(RESET)"
	@if ! command -v go >/dev/null 2>&1; then echo "$(RED)❌ Go is not installed$(RESET)"; exit 1; fi
	@if ! command -v git >/dev/null 2>&1; then echo "$(RED)❌ Git is not installed$(RESET)"; exit 1; fi
	@if ! command -v golangci-lint >/dev/null 2>&1; then echo "$(YELLOW)⚠️  golangci-lint not found. Install via: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest$(RESET)"; fi
	@echo "$(GREEN)✅ All required tools verified$(RESET)"

# =========================================================
# 🚀 Release and Packaging
# =========================================================

clean: ## Clean build artifacts
	@echo "$(YELLOW)🧹 Cleaning up...$(RESET)"
	go clean
	rm -rf $(BINARY) $(DIST_DIR)

release: clean ## Cross-compile for current or specified OS/ARCH
	@echo "$(CYAN)🌍 Building release for $(OS)/$(ARCH)...$(RESET)"
	mkdir -p $(DIST_DIR)
	CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) go build -a -ldflags "$(LDFLAGS)" \
		-o $(DIST_DIR)/$(BINARY)-$(OS)-$(ARCH) $(MAIN_FILE)
	@echo "$(GREEN)✅ Release artifact: $(DIST_DIR)/$(BINARY)-$(OS)-$(ARCH)$(RESET)"

release-all: clean ## Build binaries for all major platforms
	@echo "$(CYAN)🌎 Cross-compiling for all platforms...$(RESET)"
	mkdir -p $(DIST_DIR)
	@for os in linux darwin windows; do \
		for arch in amd64 arm64; do \
			echo "$(YELLOW)Building $$os/$$arch...$(RESET)"; \
			CGO_ENABLED=0 GOOS=$$os GOARCH=$$arch go build -a -ldflags "$(LDFLAGS)" \
				-o $(DIST_DIR)/$(BINARY)-$$os-$$arch $(MAIN_FILE); \
		done; \
	done
	@echo "$(GREEN)✅ All builds complete in $(DIST_DIR)/$(RESET)"

# =========================================================
# 🐳 Docker Build (Optional)
# =========================================================

docker: ## Build a Docker image for openmrsctl
	@echo "$(CYAN)🐳 Building Docker image...$(RESET)"
	docker build -t $(BINARY):$(VERSION) .

# =========================================================
# 📚 Documentation
# =========================================================

help: ## Display available targets
	@echo "$(YELLOW)Available targets:$(RESET)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
	| sort \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "  $(CYAN)%-15s$(RESET) %s\n", $$1, $$2}'
