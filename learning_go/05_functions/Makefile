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

shadow: vet
	shadow ./...
.PHONY: shadow

build: vet
	go build simple_cat.go
.PHONY: build

run: vet
	go run simple_cat.go
.PHONY: run


