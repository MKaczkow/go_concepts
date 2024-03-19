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
	go build process_and_gather.go
.PHONY: build

run: vet
	go run process_and_gather.go
.PHONY: run
