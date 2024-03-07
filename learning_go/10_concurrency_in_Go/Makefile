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
	go build closed_channel_example.go
.PHONY: build

run: vet
	go run closed_channel_example.go
.PHONY: run
