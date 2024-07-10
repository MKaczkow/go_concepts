.DEFAULT_GOAL := run 

fmt:
	go fmt ./...
.PHONY: fmt

lint: fmt
	golint ./...
.PHONY: lint

build: fmt
	go build main.go
.PHONY: build

run: fmt
	go run texttest_fixture.go ${ARGS}
.PHONY: run

test: fmt
	go test ./...
.PHONY: test

coverage: fmt
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out
.PHONY: coverage
