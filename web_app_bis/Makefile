.DEFAULT_GOAL := run

fmt:
	go fmt ./...
.PHONY: fmt

vet: fmt
	go vet ./...
.PHONY: vet

install: vet
	go install
.PHONY: install

test: install
	go test -cover ./...

build: install
	go build
.PHONY: build

run: build
	go run main.go
.PHONY: run
