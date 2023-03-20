COMMIT_TAG:=$(shell git describe --tags --abbrev=0 2> /dev/null || echo "no tag")
COMMIT_HASH:=$(shell git rev-parse --short HEAD)
COMMIT_MSG:=$(shell git log -1 --pretty=format:"%s")
BUILD_TIME:=$(shell date)

APP_NAME:=bosen-backend
APP_VERSION:=${COMMIT_HASH}
ENTRYPOINT?=./main.go

IMAGE_REGISTRY:=docker.io/aldoreno
IMAGE_NAME:=$(IMAGE_REGISTRY)/$(APP_NAME)
IMAGE_TAG:=$(COMMIT_HASH)

PWD:=$(shell pwd -L)
DOCKER_RUN:=docker run --rm -it -w /app -v ${PWD}:/app -v ${GOPATH}/pkg/mod/cache:/go/pkg/mod/cache golang:1.20-buster

BUILD_TAG?=development

.PHONY: all
all: help
help: ## Display help screen
	@echo "Usage:"
	@echo "	make [COMMAND]"
	@echo "	make help \n"
	@echo "Commands: \n"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: test
init: ## Create environment variables
	cp .env.example .envrc

.PHONY: clean
clean: ## Clean build bin/
	@rm -rf bin/$(APP_NAME)

.PHONY: build
build: clean ## Build golang project
	@echo "Compiling ..."
	CGO_ENABLED=1 go build \
		-v -race \
		-tags jsoniter
		-ldflags="\
			-X \"bosen/manifest.AppName=$(APP_NAME)\" \
			-X \"bosen/manifest.AppVersion=$(APP_VERSION)\" \
			-X \"bosen/manifest.BuildTime=$(BUILD_TIME)\" \
			-X \"bosen/manifest.CommitMsg=$(COMMIT_MSG)\" \
			-X \"bosen/manifest.CommitHash=$(COMMIT_HASH)\" \
			-X \"bosen/manifest.ReleaseVersion=$(BUILD_TAG)\"" \
		-o bin/$(APP_NAME) $(ENTRYPOINT)

.PHONY: run
run: ## Run golang project
	go run \
		-tags=jsoniter \
		-ldflags="\
			-X \"bosen/manifest.AppName=$(APP_NAME)\" \
			-X \"bosen/manifest.AppVersion=$(APP_VERSION)\" \
			-X \"bosen/manifest.BuildTime=$(BUILD_TIME)\" \
			-X \"bosen/manifest.CommitMsg=$(COMMIT_MSG)\" \
			-X \"bosen/manifest.CommitHash=$(COMMIT_HASH)\" \
			-X \"bosen/manifest.ReleaseVersion=$(BUILD_TAG)\"" \
		main.go

.PHONY: test-local
test-local: ## Run local golang tests
	go test -cover -race ./...

.PHONY: test
test: ## Run golang tests
	$(DOCKER_RUN) go test -cover -race ./...

.PHONY: test-report
test-report: ## Run tests with HTML coverage report
	$(DOCKER_RUN) go test -covermode=count -coverprofile coverage.out -p=1 ./... && \
	go tool cover -html=coverage.out -o coverage.html && \
	xdg-open ./coverage.html

.PHONY: test-report-func
test-report-func: ## Run tests with func report -covermode=set
	$(DOCKER_RUN) go test -covermode=set -coverprofile=coverage.out -p=1 ./... && \
	go tool cover -func=coverage.out

.PHONY: test-report-text
test-report-text:
	go test ./... -coverprofile=coverage.txt -covermode=atomic

# https://golangci-lint.run/usage/linters/
.PHONY: lint
lint: ## Lint with golangci-lint
	docker run --rm -it -v $(PWD):/app -w /app golangci/golangci-lint:v1.51-alpine \
	golangci-lint run \
		--exclude-use-default=false \
		--enable=gocyclo \
		--enable=bodyclose \
		--enable=goconst \
		--enable=sqlclosecheck \
		--enable=rowserrcheck \
		--enable=prealloc

.PHONY: fmt
fmt: ## Run go fmt
	go fmt ./...

.PHONY: vet
vet: ## Run go vet
	go vet ./...

.PHONY: up
up: ## Run docker-compose up for creating and starting containers
	docker-compose up -d

.PHONY: down
down: ## Run docker-compose down for stopping and removing containers, networks, images, and volumes
	docker-compose down --remove-orphans

.PHONY: logs
logs: ## View container log
	docker-compose logs -f app

.PHONY: docker-clean
docker-clean: ## Clean docker removes image
	docker rmi $(IMAGE_NAME):$(IMAGE_TAG)

.PHONY: docker-build
docker-build: ## Build docker image for the project
	@echo "Building image \"$(IMAGE_NAME):$(IMAGE_TAG)\" ..."
	@docker build \
		--build-arg BUILDKIT_CONTEXT_KEEP_GIT_DIR=1 \
		--target production \
		--tag $(IMAGE_NAME):$(IMAGE_TAG) \
		-f Dockerfile \
		..

.PHONY: docker-run
docker-run: docker-deps ## Run docker container for image project
	docker run \
		--rm \
		-it \
		-e APP_NAME=$(APP_NAME) \
		-e APP_VERSION=$(APP_VERSION) \
		-p 8080:8080 \
		--env-file .env \
		--network bridge \
		--name $(APP_NAME) \
		$(IMAGE_TAG)

docker-deps:
	# docker-compose up -d postgres mongodb-primary mongodb-secondary mongodb-arbiter
	# sleep 3