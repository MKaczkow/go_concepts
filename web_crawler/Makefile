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
	go build main.go
.PHONY: build

run: vet
	go run main.go
.PHONY: run

run-no-vet:
	go run main.go
.PHONY: run-no-vet

test: vet
	go test ./...
.PHONY: test

test-verbose: vet
	go test -v ./...
.PHONY: test-verbose

test-coverage: vet
	go test -cover ./...
.PHONY: test-coverage