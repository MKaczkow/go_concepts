.DEFAULT_GOAL := build

run:
	go run main.go
.PHONY: run

test-verbose:
	go test -v ./... 
.PHONY: test-verbose

test:
	go test ./...
.PHONY: test

build:
	go build -o monkey
.PHONY: build
