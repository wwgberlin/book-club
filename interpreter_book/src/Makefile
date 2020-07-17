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
