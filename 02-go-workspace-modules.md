# 02. Go Workspace & Modules

In this chapter we set up your Go development environment and create your first Go project.

[← Back to [TOC](README.md#table-of-contents)] | [← [01-introduction.md](./01-introduction.md)] | [→ [03-error-handling.md](./03-error-handling.md)]

## What is a Go module?

A **Go module** is how Go manages dependencies — external libraries your code uses. Every modern Go project is a module: a directory tree with a `go.mod` file at its root.

- **`go.mod`**: the module path (usually your repo URL), the language version, and the libraries you depend on.
- **`go.sum`**: checksums of those dependencies so builds fail if a download is tampered with. Commit this file.

You do **not** need `GOPATH` mode. Since Go 1.16, modules are the default. Put the project anywhere you like.

### Initializing a module

Use `go mod init` with a unique module path, typically the GitHub (or GitLab) repository URL:

```bash
mkdir heroes-service && cd heroes-service
go mod init github.com/yourusername/heroes-service
```

That creates a `go.mod` similar to:

```
module github.com/yourusername/heroes-service

go 1.25.0
```

**Why `go 1.25.0` if we installed Go 1.26?** Starting with Go 1.26, `go mod init` writes the *previous* major release as the language version so new modules stay compatible with both currently supported toolchains. That is intentional. To require 1.26 language features, run:

```bash
go get go@1.26.0
```

You may also see a `toolchain` line. That tells the `go` command which toolchain to download automatically (`GOTOOLCHAIN=auto`, the default since Go 1.21) if the local toolchain is too old.

---

## Setting up Go with mise

Before relying on `go`, pin the version for this project.

### 1. Tell mise to use Go

```bash
mise use go@1.26
```

This creates **`mise.toml`** in the project root (the default name; older docs and some repos still use `.mise.toml`, which mise still reads). Anyone who clones the repo can run `mise install` and get the same toolchain.

A typical file looks like this:

```toml
[tools]
go = "1.26"
golangci-lint = "2"

[tasks.run]
description = "Run the application"
run = "go run ."

[tasks.test]
description = "Run all tests"
run = "go test ./..."
```

Pinning `1.26` tracks the 1.26.x patch line (security fixes). Avoid `go@latest` in a shared repo — it moves under you.

### 2. `mise exec` vs `mise run`

These are different commands. Mixing them up is the most common setup bug in this tutorial's earlier editions.

| Command | What it does |
|---------|----------------|
| `mise exec -- go run .` | Run an arbitrary command with the project's tools on `PATH` |
| `mise run test` | Run a **named task** from `[tasks.*]` in `mise.toml` |
| `go run .` | Works once mise is **activated** in your shell |

If you have not activated mise yet:

```bash
mise exec -- go version
mise exec -- go run .
```

After `eval "$(mise activate bash)"` (see Chapter 1), plain `go` is enough.

### 3. Optional: mise tasks

With the `[tasks.run]` block above:

```bash
mise run run
mise run test
```

Useful, but not required. Chapter 5 uses a Makefile for the same idea so the commands work the same on every laptop and in CI.

---

## Essential `go` commands

- **`go mod tidy`**: add missing modules, drop unused ones, refresh `go.sum`. Run this after you add or remove imports.
- **`go run .`**: compile and run in one step. Fine for day-to-day development.
- **`go build -o bin/heroes ./cmd/heroes`**: compile a binary you can ship. (We will create `cmd/heroes` in Chapter 4.)
- **`go test ./...`**: run all tests in this module.
- **`go get example.com/mod@v1.2.3`**: add or upgrade a dependency.
- **`go install example.com/cmd@latest`**: install a *binary* into `$(go env GOPATH)/bin`. Prefer pinning tools with mise or `go get -tool` (Chapter 14) so teammates share versions.

---

## Your first "Hello World"

Create `main.go` in the project root (we will move it under `cmd/` in Chapter 4):

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, your Go environment is ready!")
}
```

Run it:

```bash
# if mise is activated
go run .

# if it is not
mise exec -- go run .
```

If you see `Hello, your Go environment is ready!`, the workspace is working.

### Confirm the toolchain

```bash
go version          # should report go1.26.x
go env GOMOD        # path to this project's go.mod
go env GOPROXY      # default is https://proxy.golang.org,direct
```

### Private modules (`GOPRIVATE`)

`go` downloads public modules through [proxy.golang.org](https://proxy.golang.org) by default. That proxy **cannot see private repos**. If you import `github.com/your-org/secret-lib`, set:

```bash
go env -w GOPRIVATE=github.com/your-org/*

# optional: also skip the checksum database for those paths
go env -w GONOSUMDB=github.com/your-org/*
```

Or, for one project, put the same values in `mise.toml`:

```toml
[env]
GOPRIVATE = "github.com/your-org/*"
```

You still need Git credentials (`gh auth`, an SSH agent, or a `insteadOf` rewrite). A 404 from the public proxy almost always means `GOPRIVATE` is missing, not that the module does not exist.

Company laptops sometimes set `GOPROXY=off` or a corporate Athens/Goproxy. Ask your team; do not copy a `GOPROXY` from a blog into a personal project.

---

## A small `.gitignore`

Create `.gitignore` now so build artifacts never get committed:

```gitignore
# Binaries
/bin/
*.exe
*.test

# Coverage and profiling
coverage.out
*.prof

# Editor / OS
.idea/
.vscode/*
!.vscode/settings.json
!.vscode/extensions.json
.DS_Store
```

---

## Try this

1. Run `go env GOPATH GOROOT GOMOD`. Explain to yourself what each path is for. (`GOROOT` is the toolchain; `GOPATH` is where `go install` puts binaries; `GOMOD` is *this* module.)
2. Break `main.go` (delete a quote) and save. Confirm the Go extension underlines the error.

## Next step

Now that the environment works, let's handle errors safely.

[03-error-handling.md →](./03-error-handling.md)

[← Back to [TOC](README.md#table-of-contents)]
