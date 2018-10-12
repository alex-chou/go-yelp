SHELL := /bin/bash

test:
	# Run tests.
	go test -race ./...

test_coverage:
	# Run tests and generate coverage profile
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
