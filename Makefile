GO := go

.PHONY: test test-all test-domain test-domain-user test-usecase test-v test-race test-cover

test: test-all

test-all:
	$(GO) test ./...

test-domain:
	$(GO) test ./internal/domain/...

test-domain-user:
	$(GO) test ./internal/domain/user

test-usecase:
	$(GO) test ./internal/usecase/...

test-v:
	$(GO) test -v ./...

test-race:
	$(GO) test -race ./...

test-cover:
	$(GO) test -cover ./...
