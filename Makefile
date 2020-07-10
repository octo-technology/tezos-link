# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Module
CMD_PATH=github.com/octo-technology/tezos-link/cmd

# api service
API=api
API_PATH=$(CMD_PATH)/$(API)
API_CMD=./cmd/$(API)
API_BIN=./bin/$(API)

# proxy service
PROXY=proxy
PROXY_PATH=$(CMD_PATH)/$(PROXY)
PROXY_CMD=./cmd/$(PROXY)
PROXY_BIN=./bin/$(PROXY)

# proxy service for Test Network
PROXY_CARTHAGENET=proxy-carthagenet


# snapshot lambda
SNAPSHOT=snapshot
SNAPSHOT_PATH=$(CMD_PATH)/$(SNAPSHOT)
SNAPSHOT_CMD=./cmd/$(SNAPSHOT)
SNAPSHOT_BIN=./bin/$(SNAPSHOT)

# old metrics cleaner lambda
METRICS_CLEANER=metrics
METRICS_CLEANER_PATH=$(CMD_PATH)/$(METRICS_CLEANER)
METRICS_CLEANER_CMD=./cmd/$(METRICS_CLEANER)
METRICS_CLEANER_BIN=./bin/$(METRICS_CLEANER)

.PHONY: all build build-unix build-api build-proxy build-snapshot-lambda test clean clean-app run deps docker-images docker-tag docs

all: test build
build: build-frontend
	$(GOBUILD) -o $(API_BIN) $(API_CMD) && chmod +x $(API_BIN)
	$(GOBUILD) -o $(PROXY_BIN) $(PROXY_CMD) && chmod +x $(PROXY_BIN)
	$(GOBUILD) -o $(SNAPSHOT_BIN) $(SNAPSHOT_CMD) && chmod +x $(SNAPSHOT_BIN)
	$(GOBUILD) -o $(METRICS_CLEANER_BIN) $(METRICS_CLEANER_CMD) && chmod +x $(METRICS_CLEANER_BIN)
build-frontend:
	cd web && yarn build
build-snapshot-lambda:
	CGO_ENABLED=0 GOOS=linux $(GOBUILD) -a -installsuffix cgo -o $(SNAPSHOT_BIN) $(SNAPSHOT_CMD) && chmod +x $(SNAPSHOT_BIN)
build-metrics-cleaner-lambda:
	CGO_ENABLED=0 GOOS=linux $(GOBUILD) -a -installsuffix cgo -o $(METRICS_CLEANER_BIN) $(METRICS_CLEANER_CMD) && chmod +x $(METRICS_CLEANER_BIN)
build-proxy:
	CGO_ENABLED=0 GOOS=linux $(GOBUILD) -a -installsuffix cgo -o $(PROXY_BIN) $(PROXY_CMD) && chmod +x $(PROXY_BIN)
build-api:
	CGO_ENABLED=0 GOOS=linux $(GOBUILD) -a -installsuffix cgo -o $(API_BIN) $(API_CMD) && chmod +x $(API_BIN)
build-unix: build-snapshot-lambda build-metrics-cleaner-lambda build-proxy build-api
unit-test:
	$(GOTEST) -run Unit ./... -v
integration-test:
	$(GOTEST) -run Integration ./... -v
test-ci:
	$(GOTEST) ./... -v
clean : clean-app
clean-app:
	$(GOCLEAN) $(API_PATH)
	$(GOCLEAN) $(PROXY_PATH)
	$(GOCLEAN) $(SNAPSHOT_PATH)
	$(GOCLEAN) $(METRICS_CLEANER_PATH)
	rm -f $(API_BIN)
	rm -f $(PROXY_BIN)
	rm -f $(SNAPSHOT_BIN)
	rm -f ${METRICS_CLEANER_BIN}
build-docker: build-unix
	docker-compose build
run:
	docker-compose up -d postgres node node-rolling $(API) $(PROXY) $(PROXY_CARTHAGENET)
	cd web && yarn start
run-dev:
	docker-compose up -d postgres node node-rolling $(API) $(PROXY) $(PROXY_CARTHAGENET)
	cd web && yarn start-local
down:
	docker-compose down
stop:
	docker-compose down
	if [ $$(docker ps -a | grep service) ]; then docker stop $$(docker ps -a -q); fi
	if [ $$(docker ps -a | grep service) ]; then docker rm $$(docker ps -a -q); fi
deps:
	$(GOGET) -v -d ./... && cd web && yarn install && cd -
docker-images: build-unix
	docker build -t $(API) -f build/package/$(API).Dockerfile .
	docker build -t $(PROXY) -f build/package/$(PROXY).Dockerfile .
docker-tag:
	docker tag $(API) ${REGISTRY}:$(API)-dev
	docker tag $(PROXY) ${REGISTRY}:$(PROXY)-dev
docker-push:
	echo ${PASSWORD} | docker login -u ${USERNAME} --password-stdin
	docker push ${REGISTRY}:$(API)-dev
	docker push ${REGISTRY}:$(PROXY)-dev
deploy-frontend:
	aws s3 sync web/build s3://tezoslink-front
deploy-snapshot-lambda:
	cp $(SNAPSHOT_BIN) bin/main
	zip -j bin/$(SNAPSHOT).zip bin/main
	aws s3 cp bin/$(SNAPSHOT).zip s3://tzlink-snapshot-lambda-dev/v1.0.0/$(SNAPSHOT).zip
	rm bin/$(SNAPSHOT).zip bin/main
	aws lambda update-function-code --function-name snapshot --s3-bucket tzlink-snapshot-lambda-dev --s3-key v1.0.0/$(SNAPSHOT).zip --region eu-west-1
deploy-metrics-cleaner-lambda:
	cp $(METRICS_CLEANER_BIN) bin/main
	zip -j bin/$(METRICS_CLEANER).zip bin/main
	aws s3 cp bin/$(METRICS_CLEANER).zip s3://tzlink-metrics-lambda-dev/v1.0.0/$(METRICS_CLEANER).zip
	rm bin/$(METRICS_CLEANER).zip bin/main
	aws lambda update-function-code --function-name metrics --s3-bucket tzlink-metrics-lambda-dev --s3-key v1.0.0/$(METRICS_CLEANER).zip --region eu-west-1
docs:
	if ! which swag; then go get -u github.com/swaggo/swag/cmd/swag ; fi
	swag init --generalInfo rest_controller.go --dir internal/$(API)/infrastructure/rest --output api-docs/$(API)
lint:
	vendor/golint internal/... cmd/... pkg/...
fmt:
	go fmt ./...

