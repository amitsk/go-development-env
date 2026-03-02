# 04. Makefile

[&larr; Back to [TOC](../README.md#table-of-contents)] | [&larr; [03-package-organization.md](../03-package-organization.md)] | [&rarr; [05-logging.md](../05-logging.md)]

## Why Makefile in Go?

Go has built-ins, but Makefile centralizes:
- Complex multi-step tasks
- Tool versions via mise
- Cross-platform consistency
- CI/CD friendliness

## Sample Makefile

```makefile
# Makefile
GO := $(shell which go)
GOLANGCI_LINT := $(shell which golangci-lint)
.PHONY: all build run test lint fmt vet deps clean

all: deps fmt lint test build

build:
	@mise run go build -o bin/heroes ./cmd/heroes

run: build
	@./bin/heroes

test:
	@mise run go test ./... -v -cover

lint:
	@$(GOLANGCI_LINT) run

fmt:
	@mise run gofmt -w -s . && goimports -w .

vet:
	@mise run go vet ./...

deps:
	@mise run go mod tidy
	@mise install-tools

clean:
	@rm -rf bin/ coverage.html
```

## Usage

```bash
make  # All
make test lint fmt
mise run make deps  # Ensure tools
```

Integrates mise for reproducible envs.

## Alternatives

`taskfile` or `just` for modern syntax.

## Next

[05-logging.md &rarr;](../05-logging.md)

[&larr; Back to [TOC](../README.md#table-of-contents)]
