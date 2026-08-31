.PHONY: build test

build:
	go build -o anagramsolve .

test:
	go test ./...
