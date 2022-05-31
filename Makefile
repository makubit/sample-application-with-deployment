GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
IMG ?= makubit/simple-application:latest
ESCAPED_IMG := $(shell echo ${IMG} | sed 's/\//\\\//g')
PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
NAME ?= simple-application-with-deployment

##@ Development

fmt: ## Run go fmt against code.
	go fmt ./...

vet: ## Run go vet against code.
	go vet ./...

test: fmt vet ## Run tests.
	go test ./... -coverprofile coverage.out

##@ Build

build: fmt vet ## Build service binary.
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o build/sample ./cmd/...

docker-build: test ##Build docker image.
	docker build -t ${IMG} ./build

docker-push: ## Push docker image.
	docker push ${IMG}

clean: ## Clean all binaries and test files.
	rm -rf ${PROJECT_DIR}/build/helm
	rm -f ${PROJECT_DIR}/build/sample
	rm -f ${PROJECT_DIR}/coverage.out

##@ Deploy

helm-build:
	cp -R helm $(PROJECT_DIR)/build/helm
	sed -i '' -e 's/$${img}/${ESCAPED_IMG}/g' $(PROJECT_DIR)/build/helm/values.yaml
	helm lint $(PROJECT_DIR)/build/helm

helm-deploy: helm-build ## Deploy helm resources.
	helm install ${NAME} $(PROJECT_DIR)/build/helm
	rm -rf $(PROJECT_DIR)/build/helm

helm-delete: ## Delete helm deployment.
	helm delete ${NAME}