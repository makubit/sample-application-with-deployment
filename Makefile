GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
IMG ?= makubit/sample-application:latest

##@ Development

fmt: ## Run go fmt against code.
	go fmt ./...

vet: ## Run go vet against code.
	go vet ./...

test: fmt vet ## Run tests.
	go test ./... -coverprofile cover.out

##@ Build

build: fmt vet ## Build service binary.
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o build/sample ./cmd/...

docker-build: test ##Build docker image.
	docker build -t ${IMG} ./build

docker-push: ## Push docker image.
	docker push ${IMG}