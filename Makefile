VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := $(shell basename "$(PWD)")

# Go related variables.
GOFILES := $(wildcard *.go)
STIME := $(shell date +%s)

.PHONY: all build coverage clean

## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
install: go-get

## build-public-http: start without docker
build:
	@echo "  >  Building Program..."
	go build -ldflags="-s -w" -o bin/${PROJECTNAME} main.go;
	@echo "Process took $$(($$(date +%s)-$(STIME))) seconds"

## start-http: start without docker
start-http: build
	@echo "  >  Starting Program..."
	./bin/${PROJECTNAME} http-start
	@echo "Process took $$(($$(date +%s)-$(STIME))) seconds"

generate: openapi_http

## lint: lint program
lint:
	@echo "  >  Linting Program..."
	golangci-lint run --issues-exit-code 0 --timeout 10m

coverage: ## Generate global code coverage report
	chmod u+x coverage.sh
	./coverage.sh;

coverhtml: ## Generate global code coverage report in HTML
	chmod u+x coverage.sh
	./coverage.sh html;

test: ## Generate global code coverage report in HTML
	chmod u+x test.sh
	./test.sh;

.PHONY: openapi_http
openapi_http:
	oapi-codegen -package gen -generate "types" api/v1/openapi/gamecurreny-http-api.yaml > internal/handler/gen/openapi_types.gen.go
	oapi-codegen -package=gen -generate=chi-server api/v1/openapi/gamecurreny-http-api.yaml > internal/handler/gen/openapi.gen.go

goose:
	go run github.com/pressly/goose/cmd/goose --dir db/migrations mysql "$(DSN)" $(filter-out $@,$(MAKECMDGOALS))

mock:
	mockgen --source=internal/domain/currencies/repositories/currency.go --destination=internal/domain/currencies/repositories/mocks/currency.go --package mocks
	mockgen --source=internal/domain/conversions/repositories/conversion.go --destination=internal/domain/conversions/repositories/mocks/conversion.go --package mocks
%:
	@:

