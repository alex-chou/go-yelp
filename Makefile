SHELL := /bin/bash

test:
	# Run tests.
	go test -race ./...
