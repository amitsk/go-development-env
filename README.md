# Go Development Environment Tutorial

Welcome to the **Go Development Environment** tutorial! This guide helps you set up a modern, production-ready Go workspace using best practices. It's designed for students and developers with some coding experience but new to Go tooling.

## Table of Contents

- [01-introduction.md](./01-introduction.md) - Background and setup overview
- [02-go-workspace-modules.md](./02-go-workspace-modules.md) - Modules and mise
- [03-package-organization.md](./03-package-organization.md) - Idiomatic project structure
- [04-makefile.md](./04-makefile.md) - Automating workflows
- [05-logging.md](./05-logging.md) - Structured logging with slog and zed
- [06-configuration-management.md](./06-configuration-management.md) - Viper for configs
- [07-database-access.md](./07-database-access.md) - GORM and sqlx
- [08-rest-apis.md](./08-rest-apis.md) - Gin and Echo frameworks
- [09-unit-testing.md](./09-unit-testing.md) - Testing with go test and testify
- [10-static-analysis-linting.md](./10-static-analysis-linting.md) - Linting and formatting
- [11-ci-cd.md](./11-ci-cd.md) - GitHub Actions
- [12-advanced-topics.md](./12-advanced-topics.md) - Modules, DI, CLI, containers
- [13-stitching-it-together.md](./13-stitching-it-together.md) - Full project example

## Prerequisites

- [VS Code](https://github.com/microsoft/vscode)
- [mise](https://mise.jdx.dev/) (install via brew, chocolatey, or script)
- Git

## Quick Start

```bash
mise init go@latest
go mod init example.com/myproject
# Follow chapters for full setup
```

[&larr; Back to [TOC](#table-of-contents)](#table-of-contents)

---

*Note: Tools here (mise, golangci-lint, etc.) are popular choices—explore alternatives like goreleaser, air for hot reload.*
