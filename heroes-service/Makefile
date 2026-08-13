GO            := go
GOLANGCI_LINT := golangci-lint
GOVULNCHECK   := govulncheck

# Prefer mise when it is installed so local and CI match mise.toml.
ifneq (,$(shell command -v mise 2>/dev/null))
GO            := mise exec -- go
GOLANGCI_LINT := mise exec -- golangci-lint
GOVULNCHECK   := mise exec -- govulncheck
endif

BINARY_NAME := heroes
CMD_PATH    := ./cmd/heroes

.PHONY: all tidy fmt vet lint test race vuln build run clean

all: tidy fmt vet test build

tidy:
	$(GO) mod tidy

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...

lint:
	$(GOLANGCI_LINT) run

test:
	$(GO) test ./...

race:
	$(GO) test -race -count=1 ./...

vuln:
	$(GOVULNCHECK) ./...

build:
	mkdir -p bin
	$(GO) build -o bin/$(BINARY_NAME) $(CMD_PATH)

run: build
	./bin/$(BINARY_NAME)

clean:
	rm -rf bin/ coverage.out
