VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
LDFLAGS := -s -w -X github.com/farishan-labs/user-md/checker/internal/buildinfo.Version=$(VERSION)

.PHONY: build test fmt
build:
	CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/usermd-check .

test:
	go test ./...

fmt:
	gofmt -w .

release-check:
	goreleaser check