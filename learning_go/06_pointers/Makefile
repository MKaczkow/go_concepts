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
	go build create_instance.go
.PHONY: build

run: vet
	go run create_instance.go
.PHONY: run

perf-test: vet
	go test -bench=.
.PHONY: perf-test
