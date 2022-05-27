GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

##@ Development

fmt: ## Run go fmt against code.
	go fmt ./...

vet: ## Run go vet against code.
	go vet ./...

##@ Build

build: fmt vet ## Build service binary.
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o build/sample ./cmd/...