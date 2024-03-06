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
	go build rand_and_crand.go
.PHONY: build

run: vet
	go run rand_and_crand.go
.PHONY: run
