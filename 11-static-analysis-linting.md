# 11. Static Analysis & Linting

In this chapter we catch mistakes *without running the program*.

[← Back to [TOC](README.md#table-of-contents)] | [← [10-unit-testing.md](./10-unit-testing.md)] | [→ [12-ci-cd.md](./12-ci-cd.md)]

## What is static analysis?

Two related jobs:

1. **Formatting** — the code looks like everyone else's Go (`gofmt` / `goimports`).
2. **Linting / vetting** — unused variables, ignored errors, buggy patterns, known vulnerabilities.

---

## 1. Formatting with `gofmt` and `goimports`

Go does not have a style debate. `gofmt` is the style.

```bash
go fmt ./...
```

**`goimports`** is `gofmt` plus keeping the import block tidy. The VS Code Go extension runs it on save if you enabled `source.organizeImports` (see the README).

Some teams use [gofumpt](https://github.com/mvdan/gofumpt), a stricter `gofmt`. Only adopt it if the whole team agrees.

---

## 2. `go vet`

`go vet` ships with the toolchain and finds common bugs (`fmt` verb mismatches, suspicious copies of locks, and so on):

```bash
go vet ./...
```

It is fast. Always run it. The Makefile in Chapter 5 already has a `vet` target.

### `go fix` (revamped in Go 1.26)

Go 1.26 rebuilt **`go fix`** as a *modernizer*: it rewrites code to newer, equivalent idioms. Safe to try on a branch:

```bash
go fix ./...
git diff
```

Review the diff. Modernizers should not change behavior; if one does, [file an issue](https://go.dev/issue/new).

---

## 3. `golangci-lint` v2

[**golangci-lint**](https://golangci-lint.run/) runs many linters in one process. Current line is **v2** (2.12.x as of mid-2026). v1 configs are not drop-in compatible — if you copy an old `.golangci.yml` from the internet, migrate it (`golangci-lint migrate`) or start fresh.

Install it with mise so the version is pinned next to Go:

```bash
mise use golangci-lint@2
golangci-lint version
```

```bash
golangci-lint run
```

Useful default linters include:

- **errcheck** — ignored errors
- **staticcheck** — deprecated APIs, pointless code, real bugs
- **unused** — dead functions and fields
- **govet** — the same checks as `go vet`

A minimal config (v2 format; generate one with `golangci-lint config`) is enough to start. [`heroes-service/.golangci.yml`](./heroes-service/.golangci.yml) uses the v2 `default: standard` set.

---

## 4. Vulnerability scanning

[**govulncheck**](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck) reports known vulnerabilities in *code you actually call*, not just modules you happen to download.

```bash
mise use govulncheck@latest
mise exec -- govulncheck ./...
```

Or, with Go 1.24+'s tool directive (Chapter 14):

```bash
go get -tool golang.org/x/vuln/cmd/govulncheck
go tool govulncheck ./...
```

Run this in CI. Treat findings like failing tests.

---

## Automation

You should not have to remember these commands.

1. **Editor** — format (and organize imports) on save.
2. **Makefile** — `make` runs `fmt`, `vet`, `lint`, and `test` (Chapter 5).
3. **CI** — the same `make` targets on every push (Chapter 12).

## Next step

Let's run the same checks on GitHub.

[12-ci-cd.md →](./12-ci-cd.md)

[← Back to [TOC](README.md#table-of-contents)]
