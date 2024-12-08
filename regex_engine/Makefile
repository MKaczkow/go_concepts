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

shadow: vet
	shadow ./...
.PHONY: shadow

