all: check test race

.PHONY: check
check:
	go build ./...

.PHONY: test
test:
	go test ./...

.PHONY: race
race:
	go test -race ./...

.PHONY: fmt
fmt:
	gofmt -s -w ./...

.PHONY: monkey
monkey:
	go build -o monkey ./cmd/interpreter
