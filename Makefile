all: test

test:
	@go test -mod=vendor -cover ./...

.PHONY: all test
