include .env

MODULE = $(shell go list -m)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || echo "1.0.0")
PACKAGES := $(shell go list ./... | grep -v /vendor/)
LDFLAGS := -ldflags "-X ${MODULE}/internal/config.Version=${VERSION}"
GOARCH := $(shell go env GOARCH)
MIGRATION_PATH = internal/database/migrations


default: help
# generate help info from comments: thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help build
help: ## help information about make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version: ## display the version of the API server
	@echo $(VERSION)

.PHONY: all
all: ## execute all the targets
	@echo "Executing all targets"
	make lint
	make fmt
	make test
	make build

# ==============================================================================
# Docker
.PHONY: build-docker
build-docker: ## build the API server as a docker image
	docker build -f Dockerfile -t $(MODULE):$(VERSION) --build-arg GO_VERSION=1.22  .

.PHONY: push-docker
push-docker: ## build the API server as a docker image
	docker push  $(MODULE):$(VERSION)

.PHONY: run-docker
run-docker: ## run docker image
	docker run $(MODULE):$(VERSION)

.PHONY: stop-docker
stop-docker: ## stop docker image
	docker stop $(shell docker ps -q --filter ancestor=$(MODULE):$(VERSION))

.PHONY: clean-docker
clean-docker: ## clean docker image
	docker rmi $(MODULE):$(VERSION)

# ==============================================================================
# Tools commands
.PHONY: lint
lint: ## run golint on all Go package
	echo "Starting linters"
	golangci-lint run ./...

.PHONY: fmt
fmt: ## run gofmt on all Go package
	echo "Run fmt"
	gofmt -s -w .

# ==============================================================================
# Main

.PHONY: run
run: ## run the API server
	GOARCH=${GOARCH} go run ${LDFLAGS} ./cmd/web

.PHONY: build
build: ## build the API server
	GOARCH=${GOARCH} go build ${LDFLAGS} -o bin/web ./cmd/web

.PHONY: test
test: ## run the tests
	go test -cover ./...

# ==============================================================================
# Modules support

.PHONY: deps-reset
deps-reset: ## reset the go modules
	git checkout -- go.mod
	go mod tidy
	go mod vendor

.PHONY: deps-upgrade
deps-upgrade: ## upgrade the go modules
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -t -d -v ./...
	go mod tidy
	go mod vendor

.PHONY: tidy
tidy: ## tidy up the go modules
	go mod tidy
	go mod vendor

.PHONY: deps-cleancache
deps-cleancache: ## clean the go modules cache
	go clean -modcache

# ==============================================================================
# Database migration
.PHONY: migrate.create
migrate.create:
	migrate create -ext=sql -dir=$(MIGRATION_PATH) -seq $(NAME)

.PHONY: migrate.up
migrate.up:
	migrate -database $(DATABASE_URL) -path $(MIGRATION_PATH) up

.PHONY: migrate.down
migrate.down:
	migrate -database $(DATABASE_URL) -path $(MIGRATION_PATH) down

.PHONY: migrate.force
migrate.force:
	migrate -database $(DATABASE_URL) -path $(MIGRATION_PATH) force $(VERSION)

