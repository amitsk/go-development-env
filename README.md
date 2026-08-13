# Professional Go Development Environment

Welcome to the **Go Development Environment** tutorial.

Building professional software is about more than just writing code. It is about creating a sustainable, maintainable, and reliable system. This guide walks you through setting up a modern, production-ready Go workspace from scratch.

**Last reviewed:** August 2026, against [Go 1.26](https://go.dev/doc/go1.26) (current stable: 1.26.x) and [mise](https://mise.jdx.dev/).

### Who is this for?

- **Students** who know the basics of programming and want to see how professionals work.
- **Developers** new to Go who want to set up their environment correctly from day one.
- **Curious minds** who want to understand the modern Go ecosystem.

This is **not** a language tutorial. If you are new to Go syntax, start with [A Tour of Go](https://go.dev/tour/) first.

---

## Prerequisites

You will need:

- A terminal (bash, zsh, or PowerShell)
- [Git](https://git-scm.com/)
- A code editor — we use [Visual Studio Code](https://code.visualstudio.com/) plus the [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)
- About 30–60 minutes for the setup chapters, then you can work through the rest at your own pace

**Sophomore tip:** Follow along in a new folder and type every command. Reading alone will not stick.

---

## How this tutorial is organized

| Part | Chapters | What you will do |
|------|----------|------------------|
| **Getting started** | 1–2 | Install tools, pin a Go version, create a module |
| **Writing professional Go** | 3–11 | Errors, layout, automation, logging, config, data, APIs, tests, lint |
| **Shipping and beyond** | 12–17 | CI, AI assistants, advanced topics, checklist, a real service, observability |

Chapters 1–2 are sequential. After that you can skip around, but later chapters assume you have a working `mise` + module setup.

---

## Table of Contents

### Getting started

1. [**Introduction**](./01-introduction.md) — Why professional tooling matters, and how to install mise and VS Code.
2. [**Go Workspace & Modules**](./02-go-workspace-modules.md) — Pin Go with `mise`, initialize a module, run your first program.

### Writing professional Go

3. [**Error Handling (The Go Way)**](./03-error-handling.md) — Errors as values, wrapping, `errors.Is` / `errors.As`.
4. [**Package Organization**](./04-package-organization.md) — Official layout guidance vs a typical service layout.
5. [**Makefile & Automation**](./05-makefile.md) — Standardize format, lint, test, and build.
6. [**Structured Logging**](./06-logging.md) — `log/slog` (stdlib since Go 1.21).
7. [**Configuration Management**](./07-configuration-management.md) — Environment-safe settings with Viper (and simpler alternatives).
8. [**Database Access**](./08-database-access.md) — `database/sql`, sqlx, sqlc, and GORM.
9. [**Building REST APIs**](./09-rest-apis.md) — `net/http` and Gin.
10. [**Unit Testing**](./10-unit-testing.md) — `go test`, table-driven tests, race detector, fuzzing.
11. [**Static Analysis & Linting**](./11-static-analysis-linting.md) — `gofmt`, `go vet`, golangci-lint v2, `govulncheck`.

### Shipping and beyond

12. [**CI/CD**](./12-ci-cd.md) — GitHub Actions with the same tools as your laptop.
13. [**AI-Powered Development**](./13-llms.md) — Using assistants without skipping review.
14. [**Advanced Topics**](./14-advanced-topics.md) — Context, DI, `go tool`, Docker.
15. [**Stitching It All Together**](./15-stitching-it-together.md) — Your new-project checklist.
16. [**The Heroes Service (End to End)**](./16-end-to-end.md) — One runnable module that uses the earlier chapters.
17. [**Observability**](./17-observability.md) — Request IDs, `/healthz`, pprof; when to add OpenTelemetry.

The sample module lives in [`heroes-service/`](./heroes-service/). After Chapter 2 you can `cd heroes-service && go test ./...`.

---

## Quick Setup

Install [mise](https://mise.jdx.dev/), activate it in your shell, then pin a Go version:

```bash
curl https://mise.run | sh
echo 'eval "$(~/.local/bin/mise activate bash)"' >> ~/.bashrc
source ~/.bashrc

mise use --global go@1.26
go version
```

Use `zsh` or `fish` in the `activate` line if that is your shell. See [mise getting started](https://mise.jdx.dev/getting-started.html) for other installers (`brew`, `winget`, apt).

**VS Code extensions** (install via `Ctrl+Shift+X` / `Cmd+Shift+X`):

- [Go](https://marketplace.visualstudio.com/items?itemName=golang.Go) (`golang.Go`) — required. This installs `gopls`, the official language server.
- [GitLens](https://marketplace.visualstudio.com/items?itemName=eamodio.gitlens) — optional, for Git history.
- [Error Lens](https://marketplace.visualstudio.com/items?itemName=usernamehw.errorlens) — optional, inline diagnostics.

You do **not** need Prettier for Go. `gofmt` / `goimports` (via the Go extension) is the formatter.

Workspace settings that make the Go extension pleasant:

```json
{
  "go.useLanguageServer": true,
  "editor.formatOnSave": true,
  "[go]": {
    "editor.defaultFormatter": "golang.go",
    "editor.codeActionsOnSave": {
      "source.organizeImports": "explicit"
    }
  }
}
```

---

## Templates in this repo

- [`heroes-service/`](./heroes-service/) — a complete, stdlib-only API you can run and test.
- [`Makefile.template`](./Makefile.template) — copy into a new project as `Makefile`.
- [`go.mod.template`](./go.mod.template) — starting `go.mod` for Go 1.26.
- [`mise.toml.template`](./mise.toml.template) — pin Go and golangci-lint for the project.
- [`.github/workflows/ci.yml`](./.github/workflows/ci.yml) — CI for `heroes-service/`.

---

## Getting Started

[**Start Chapter 1 →**](./01-introduction.md)
