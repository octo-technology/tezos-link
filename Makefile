# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Module
CMD_PATH=github.com/octo-technology/tezos-link/cmd

# backend service
BACKEND=backend
BACKEND_PATH=$(CMD_PATH)/$(BACKEND)
BACKEND_CMD=./cmd/$(BACKEND)
BACKEND_BIN=./bin/$(BACKEND)

.PHONY: all build build-unix test clean clean-app run deps docker-images docker-tag docs

all: test build
build:
	$(GOBUILD) -o $(BACKEND_BIN) $(BACKEND_CMD)
build-unix:
	CGO_ENABLED=0 GOOS=linux $(GOBUILD) -a -installsuffix cgo -o $(BACKEND_BIN) $(BACKEND_CMD) && chmod +x $(BACKEND_BIN)
test:
	set -o pipefail && $(GOTEST) ./... -v | grep -v -E 'GET|POST|PUT'
test-ci:
	$(GOTEST) ./... -v
clean : clean-cedrus
clean-app:
	$(GOCLEAN) $(BACKEND_PATH)
	rm -f $(BACKEND_BIN)
build-docker: build-unix
	docker-compose build
run:
	docker-compose up -d postgres $(BACKEND)
stop:
	docker-compose down
	if [ $$(docker ps -a | grep service) ]; then docker stop $$(docker ps -a -q); fi
	if [ $$(docker ps -a | grep service) ]; then docker rm $$(docker ps -a -q); fi
deps:
	$(GOGET) -v -d ./...
docker-images: build-unix
	docker build -t $(BACKEND) -f build/package/$(BACKEND).Dockerfile .
docker-tag:
	docker tag $(BACKEND) ${REGISTRY}:$(BACKEND)-dev
docs:
	if ! which swag; then go get -u github.com/swaggo/swag/cmd/swag ; fi
	swag init --generalInfo rest_controller.go --dir internal/$(BACKEND)/infrastructure/rest --output api-docs/$(BACKEND)
