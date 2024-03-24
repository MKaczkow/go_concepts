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
	go build add_value_to_ctx.go
.PHONY: build

run: vet
	go run add_value_to_ctx.go
.PHONY: run
