SHELL := /bin/bash

.PHONY: example

test:
	# Run tests.
	go test -race ./...

test_coverage:
	# Run tests and generate coverage profile
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

example:
	# Run example with provided API key.
	API_KEY=$(api_key) go run example/main.go
