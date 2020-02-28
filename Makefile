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

# proxy service
PROXY=proxy
PROXY_PATH=$(CMD_PATH)/$(PROXY)
PROXY_CMD=./cmd/$(PROXY)
PROXY_BIN=./bin/$(PROXY)

.PHONY: all build build-unix test clean clean-app run deps docker-images docker-tag docs

all: test build
build:
	$(GOBUILD) -o $(BACKEND_BIN) $(BACKEND_CMD) && chmod +x $(BACKEND_BIN)
	$(GOBUILD) -o $(PROXY_BIN) $(PROXY_CMD) && chmod +x $(PROXY_BIN)
build-unix:
	CGO_ENABLED=0 GOOS=linux $(GOBUILD) -a -installsuffix cgo -o $(BACKEND_BIN) $(BACKEND_CMD) && chmod +x $(BACKEND_BIN)
	CGO_ENABLED=0 GOOS=linux $(GOBUILD) -a -installsuffix cgo -o $(PROXY_BIN) $(PROXY_CMD) && chmod +x $(PROXY_BIN)
unit-test:
	set -o pipefail && $(GOTEST) -run Unit ./... -v | grep -v -E 'GET|POST|PUT'
integration-test:
	set -o pipefail && $(GOTEST) -run Integration ./... -v | grep -v -E 'GET|POST|PUT'
test-ci:
	$(GOTEST) ./... -v
clean : clean-app
clean-app:
	$(GOCLEAN) $(BACKEND_PATH)
	$(GOCLEAN) $(PROXY_PATH)
	rm -f $(BACKEND_BIN)
	rm -f $(PROXY_BIN)
build-docker: build-unix
	docker-compose build
run:
	docker-compose up -d postgres node $(BACKEND) $(PROXY)
stop:
	docker-compose down
	if [ $$(docker ps -a | grep service) ]; then docker stop $$(docker ps -a -q); fi
	if [ $$(docker ps -a | grep service) ]; then docker rm $$(docker ps -a -q); fi
deps:
	$(GOGET) -v -d ./...
docker-images: build-unix
	docker build -t $(BACKEND) -f build/package/$(BACKEND).Dockerfile .
	docker build -t $(PROXY) -f build/package/$(PROXY).Dockerfile .
docker-tag:
	docker tag $(BACKEND) ${REGISTRY}:$(BACKEND)-dev
	docker tag $(PROXY) ${REGISTRY}:$(PROXY)-dev
docs:
	if ! which swag; then go get -u github.com/swaggo/swag/cmd/swag ; fi
	swag init --generalInfo rest_controller.go --dir internal/$(BACKEND)/infrastructure/rest --output api/$(BACKEND)
lint:
	vendor/golint internal/... cmd/...
fmt:
	go fmt ./...
