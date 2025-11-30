.DEFAULT_GOAL := vet 

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
	go build main.go
.PHONY: build

run-udp: vet
	go run listeners/udp/main.go
.PHONY: run-udp

run-tcp: vet
	go run listeners/tcp/main.go
.PHONY: run-tcp
