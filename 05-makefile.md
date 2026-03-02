# 05. Makefile & Automation

In this chapter, we'll learn how to automate repetitive tasks using a **Makefile**.

[&larr; Back to [TOC](../README.md#table-of-contents)] | [&larr; [04-package-organization.md](./04-package-organization.md)] | [&rarr; [06-logging.md](./06-logging.md)]

## What is a Makefile?

A **Makefile** is a simple text file used by the `make` build tool to automate tasks. While Go has powerful built-in commands like `go build` and `go test`, a Makefile allows you to:

- **Combine multiple steps** into a single command (e.g., format, lint, and then test).
- **Standardize commands** so everyone on your team runs them the same way.
- **Save time** by not having to remember long, complex command-line arguments.

---

## The Anatomy of a Makefile

A Makefile is made up of **targets**, **dependencies**, and **commands**:

```makefile
target: dependency1 dependency2
    command to run
```

- **Target**: The name of the task you want to run (e.g., `build`).
- **Dependency**: Other targets that must run *before* this one.
- **Command**: The actual shell command to execute. **Note:** Commands MUST be indented with a **Tab** character, not spaces.

### What is `.PHONY`?

By default, `make` thinks its targets are files. If you have a folder named `build`, `make build` might not run because it thinks the "file" already exists. We use `.PHONY` to tell `make` that these are just command names, not files.

---

## A Professional Go Makefile

Here is a sample Makefile that integrates with **mise** to ensure the correct Go version is always used:

```makefile
# Define variables for reuse
BINARY_NAME=heroes

.PHONY: all build run test lint fmt tidy clean

# Default target when you just run 'make'
all: tidy fmt lint test build

# Initialize and clean up dependencies
tidy:
	mise run go mod tidy

# Format the code according to Go standards
fmt:
	mise run go fmt ./...

# Run static analysis to find potential bugs
lint:
	# We'll learn about golangci-lint in chapter 10
	mise run golangci-lint run

# Run all unit tests
test:
	mise run go test -v ./...

# Compile the application into a binary
build:
	mkdir -p bin
	mise run go build -o bin/$(BINARY_NAME) ./cmd/heroes

# Run the compiled application
run: build
	./bin/$(BINARY_NAME)

# Remove build artifacts
clean:
	rm -rf bin/
```

---

## How to Use It

Once you've created this file in your project root, you can run:

- `make build`: To just compile your app.
- `make test`: To run your tests.
- `make`: To run the `all` target (tidy, fmt, lint, test, and build).

---

## Why use `mise run` inside the Makefile?

By prefixing your commands with `mise run`, you guarantee that the Makefile uses the exact version of Go (and other tools) defined in your project. This prevents the \"it works on my machine\" problem!

## Next Step

Now that we can automate our workflow, let's look at how to properly record what our application is doing.

[06-logging.md &rarr;](./06-logging.md)

[&larr; Back to [TOC](../README.md#table-of-contents)]
