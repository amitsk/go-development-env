# 05. Makefile & Automation

In this chapter we automate repetitive tasks with a **Makefile**.

[← Back to [TOC](README.md#table-of-contents)] | [← [04-package-organization.md](./04-package-organization.md)] | [→ [06-logging.md](./06-logging.md)]

## What is a Makefile?

A **Makefile** is a text file used by the `make` build tool. Go already has `go build` and `go test`, but a Makefile lets you:

- **Combine steps** (format, vet, lint, test, then build) into one command.
- **Standardize** so everyone on the team — and CI — runs the same thing.
- **Forget long flags** (`go test -race -count=1 ./...`).

mise can do the same job with `[tasks]` in `mise.toml`. We still show Make because it is everywhere in existing Go repos. Use one or the other; do not maintain two competing lists of commands.

On Windows, use Git Bash, WSL, or [Task](https://taskfile.dev/) instead of native `nmake`.

---

## Anatomy of a Makefile

```makefile
target: dependency1 dependency2
	command to run
```

- **Target**: the task name (`build`).
- **Dependency**: other targets that must run first.
- **Command**: a shell command. It **must** be indented with a real **Tab**, not spaces.

### What is `.PHONY`?

By default `make` thinks targets are files. If you have a folder named `build`, `make build` may no-op. `.PHONY` marks names that are always commands.

---

## A professional Go Makefile

Prefix tools with `mise exec --` so the Makefile uses the versions in `mise.toml` even if the developer forgot to activate mise:

```makefile
# Tools go through mise so CI and laptops match mise.toml.
GO             := mise exec -- go
GOLANGCI_LINT  := mise exec -- golangci-lint
GOVULNCHECK    := mise exec -- govulncheck

BINARY_NAME := heroes
CMD_PATH    := ./cmd/heroes

.PHONY: all tidy fmt vet lint test race build run clean vuln

# Default: everything a PR should pass
all: tidy fmt vet lint test build

tidy:
	$(GO) mod tidy

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...

# golangci-lint is installed in Chapter 11
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
	rm -rf bin/
```

A complete copy lives in [`Makefile.template`](./Makefile.template). The sample app uses the same targets in [`heroes-service/Makefile`](./heroes-service/Makefile) (`mise exec --` when `mise` is on `PATH`).

---

## How to use it

From the project root:

- `make build` — compile only
- `make test` — unit tests
- `make race` — tests plus the race detector (slower; use before you push)
- `make` — the `all` target (tidy, fmt, vet, lint, test, build)

---

## Why `mise exec --` inside Make?

`mise exec -- <cmd>` puts the project's Go and linters on `PATH` for that one command. That is what kills "it works on my machine":

- Laptop A has system Go 1.22.
- The repo pins `go = "1.26"` in `mise.toml`.
- `make test` still runs 1.26.

If you prefer named tasks instead of Make, put the same recipes under `[tasks]` and run `mise run test`.

## Next step

Now that we can automate the workflow, let's record what the application is doing.

[06-logging.md →](./06-logging.md)

[← Back to [TOC](README.md#table-of-contents)]
