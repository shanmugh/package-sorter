.PHONY: all
all: test run

.PHONY: test
test:
	go test -v ./...

.PHONY: help
help: build
	./package-sorter -h

.PHONY: run
run: build
	./package-sorter -i ./input.json

.PHONY: build
build:
	go build -o package-sorter main.go

.PHONY: lint
lint:
	golangci-lint run

.PHONY: deps
deps:
	go mod tidy
	go mod verify
