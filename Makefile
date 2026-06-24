PREFIX ?= /usr/local
BINDIR ?= $(PREFIX)/bin

.PHONY: build test install clean

build:
	go build ./cmd/...

test:
	go test ./...

install: build
	install -d $(BINDIR)
	install $(shell go list -f '{{.Dir}}' ./cmd/... | sed 's|/[^/]*$$||' | sort -u | xargs -I{} find {} -maxdepth 1 -type f -perm +111 2>/dev/null || true) $(BINDIR) 2>/dev/null || go install ./cmd/...

clean:
	go clean ./...
	rm -rf bin/
