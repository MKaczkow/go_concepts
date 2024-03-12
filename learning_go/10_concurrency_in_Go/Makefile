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
	go build goprocedures_no_deadlock.go
.PHONY: build

run: vet
	go run goprocedures_no_deadlock.go
.PHONY: run
