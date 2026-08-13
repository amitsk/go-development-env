# 01. Introduction

Welcome to the Go Development Environment tutorial. This guide is for students and developers who can already write some code but are new to the professional tooling and workflows used in the Go ecosystem.

[← Back to [README](README.md)]

## Goal of this tutorial

The goal is to move beyond "writing code that runs" and start building **production-ready** applications. We focus on the environment, tools, and processes that help professional engineers keep code high quality.

**This is not a tutorial on how to program in Go.** If you are new to the language, start here:

- [A Tour of Go](https://go.dev/tour/) — interactive introduction to syntax and features
- [Effective Go](https://go.dev/doc/effective_go) — idiomatic style
- [How to Write Go Code](https://go.dev/doc/code) — packages, modules, and the `go` command
- [Go by Example](https://gobyexample.com/) — short annotated programs

## Why professional tooling matters

In a professional setting we care about more than the logic of the code:

- **Version control (Git)** — track changes, collaborate, and undo mistakes safely.
- **Formatting** — one official style (`gofmt`) so reviews focus on behavior, not braces.
- **Linting & vetting** — catch ignored errors, dead code, and other bugs before you run the program.
- **Unit testing** — small tests so behavior stays correct as you change things.
- **CI/CD** — automatically format-check, test, and build on every push.

## Tools we will use

We've selected a modern stack that is popular in industry in 2026.

### 1. Mise (version and tool management)

Managing different versions of Go (or other languages) by hand is a headache. [**mise**](https://mise.jdx.dev/) (formerly `rtx`) lets you pin the Go version — and other tools like `golangci-lint` — in a `mise.toml` file that lives with the project.

**Installation** (see the [official install guide](https://mise.jdx.dev/installing-mise.html) for more options):

```bash
# Recommended on macOS and Linux
curl https://mise.run | sh

# Then activate in your shell (bash shown; use zsh or fish if that is your shell)
echo 'eval "$(~/.local/bin/mise activate bash)"' >> ~/.bashrc
source ~/.bashrc
```

- **macOS (Homebrew):** `brew install mise`
- **Windows:** `winget install jdx.mise`
- **Linux packages:** apt/yum/nix — see [installing mise](https://mise.jdx.dev/installing-mise.html)

Activation is important. Without it, `go` on your `PATH` may be a system package that is years old. After activation, `which go` should point at a mise-managed binary.

### 2. VS Code and `gopls`

We recommend [**Visual Studio Code**](https://code.visualstudio.com/) with the [official Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go). The extension uses **`gopls`**, the official Go language server, for completion, jump-to-definition, and format-on-save.

The first time you open a `.go` file, the extension will offer to install extra tools (`gopls`, `goimports`, `staticcheck`, …). Accept that prompt.

You do **not** need Prettier for Go. Use `gofmt` / `goimports` through the Go extension.

### 3. A note on alternatives

The tools here — `mise`, `golangci-lint`, `testify` — are solid, modern choices. The Go ecosystem has many others. Once you are comfortable, you might explore:

- **Task runners:** [Task](https://taskfile.dev/) (`Taskfile.yml`) instead of Make, or mise's own `[tasks]`
- **Version managers:** `asdf` (mise is a compatible successor), or the official [go.dev/dl](https://go.dev/dl/) installers plus the `toolchain` line in `go.mod`
- **Linters:** `staticcheck` on its own, or `go vet` + `govulncheck` if you want a smaller set
- **Editors:** GoLand, Neovim with `gopls`, Zed, or Cursor (VS Code–compatible)

Pick what your team already uses. Consistency beats novelty.

## Next step

Now that we have the "why" and "what", let's set up a Go workspace.

[02-go-workspace-modules.md →](./02-go-workspace-modules.md)

[← Back to [TOC](README.md#table-of-contents)]
