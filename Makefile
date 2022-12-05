include .env

.EXPORT_ALL_VARIABLES:

test:
	go run ./example/example.go

.PHONY: test
