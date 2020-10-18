SHELL := /bin/bash

SRC = $(shell find . -type f -name '*.go' -not -iname '*.pb.*' -not -path '**/mocks/*')

lint:
	@golangci-lint run -v -c .golangci.yml ./...

test:
	@go test -cover ./...

coverage:
	@go test -coverpkg=./internal/... -coverprofile=coverage.out ./internal/... > /dev/null
	@sed -i '/mock.go/d' coverage.out
	@sed -i '\#/mocks#d' coverage.out
	@go tool cover -func coverage.out

build:
	@echo "  >  Building binary..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o people-analytics cmd/app/main.go

clean:
	@go clean

fmt:
	@gofmt -s -l -w $(SRC)

goimports:
	@goimports -w -local github.com/golangci/golangci-lint $(SRC)

format: fmt goimports

.PHONY: all test