.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY: fmt

lint: fmt
	golint ./...
.PHONY: lint

vet: fmt
	go vet ./...
.PHONY: vet

shadow: vet
	shadow ./...
.PHONY: shadow

build: vet
	go build for_range_loop.go
.PHONY: build

run: vet
	go run for_range_loop.go
.PHONY: run


