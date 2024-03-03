.DEFAULT_GOAL := run

fmt:
	go fmt ./...
.PHONY: fmt

lint: fmt
	golint ./...
.PHONY: lint

vet: fmt
	go vet ./...
.PHONY: vet

build: vet
	go build interfaces_example.go
.PHONY: build

run: vet
	go run interfaces_example.go
.PHONY: run
