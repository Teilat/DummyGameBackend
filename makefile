BINARY_NAME := $(shell git config --get remote.origin.url | awk -F/ '{print $$5}' | awk -F. '{print tolower($$1)}')
BINARY_VERSION := $(shell git describe --always)
BINARY_BUILD_DATE := $(shell date +%d.%m.%Y)
WIN_BINARY_NAME := $(BINARY_NAME).exe
BUILD_FOLDER := .build
DOCKER_REPO := "teilat"

PRINTF_FORMAT := "\033[35m%-18s\033[33m %s\033[0m\n"

ABS_PATH := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))
ifeq ($(shell go env GOHOSTOS), windows)
	ABS_PATH = $(CURDIR)
endif

.PHONY: all build windows linux vendor gen-webapi gen-ssl clean docker-build lint

all: build

build: windows linux ## Default: build for windows and linux

gen-swagger:
	cd ./webapi && swag init --parseInternal -g webapi.go

windows: vendor ## Build artifacts for windows
	@printf $(PRINTF_FORMAT) BINARY_NAME: $(WIN_BINARY_NAME)
	@printf $(PRINTF_FORMAT) BINARY_BUILD_DATE: $(BINARY_BUILD_DATE)
	mkdir -p $(BUILD_FOLDER)/windows
	env GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc  go build -ldflags "-s -w -X $(BINARY_NAME).Version=$(BINARY_VERSION) -X $(BINARY_NAME).BuildDate=$(BINARY_BUILD_DATE)" -o $(BUILD_FOLDER)/windows/$(WIN_BINARY_NAME) .

linux: vendor ## Build artifacts for linux
	@printf $(PRINTF_FORMAT) BINARY_NAME: $(BINARY_NAME)
	@printf $(PRINTF_FORMAT) BINARY_BUILD_DATE: $(BINARY_BUILD_DATE)
	mkdir -p $(BUILD_FOLDER)/linux
	env GOOS=linux GOARCH=amd64  go build -ldflags "-s -w -X $(BINARY_NAME).Version=$(BINARY_VERSION) -X $(BINARY_NAME).BuildDate=$(BINARY_BUILD_DATE)" -o $(BUILD_FOLDER)/linux/$(BINARY_NAME) .

docker-build: linux ## Build artifacts for linux
	docker build -t $(BINARY_NAME) .

docker-tag: docker-build ## Generate container `{version}` tag
	@echo 'create tag latest'
	docker tag $(BINARY_NAME) $(DOCKER_REPO)/$(BINARY_NAME):latest

docker-publish: docker-tag ## Build the container without caching
	@echo 'publish latest to $(DOCKER_REPO)'
	docker push $(DOCKER_REPO)/$(BINARY_NAME):latest

docker-up: ## Build the container without caching
	docker compose up --detach --force-recreate

vendor: ## Get dependencies according to go.sum
	env GO111MODULE=auto go mod tidy
	env GO111MODULE=auto go mod vendor

clean: ## Remove vendor and artifacts
	rm -rf vendor
	rm -rf $(BUILD_FOLDER)

lint: vendor
	docker run --rm -t -v "$(ABS_PATH)":/app -w /app --entrypoint=golangci-lint golangci/golangci-lint:v1.50.1 run -v
