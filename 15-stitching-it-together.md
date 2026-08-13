# 15. Stitching It All Together

A reusable workflow for the next Go project you start.

[← Back to [TOC](README.md#table-of-contents)] | [← [14-advanced-topics.md](./14-advanced-topics.md)] | [→ [16-end-to-end.md](./16-end-to-end.md)]

Prefer cloning the ideas in [`heroes-service/`](./heroes-service/) over starting from an empty folder. Chapter 16 walks that module file by file.

## The new-project checklist

### 1. Initialize the environment

- [ ] `mkdir my-app && cd my-app && git init`
- [ ] Pin Go: `mise use go@1.26`
- [ ] Pin the linter: `mise use golangci-lint@2`
- [ ] `go mod init github.com/yourusername/my-app`
- [ ] If you want the *language* version to be 1.26 (not the `go mod init` default of 1.25): `go get go@1.26.0`
- [ ] Copy [`.gitignore` ideas from Chapter 2](./02-go-workspace-modules.md#a-small-gitignore)

### 2. Set up the structure

Start with `cmd/my-app/main.go`. Add `internal/` packages only when the second file needs a home.

```bash
mkdir -p cmd/my-app internal
```

- [ ] Copy [`Makefile.template`](./Makefile.template) to `Makefile`
- [ ] Copy [`mise.toml.template`](./mise.toml.template) if you did not run `mise use`
- [ ] Add `README.md` and `config.example.yaml`

### 3. Add tools and libraries as you need them

Do not `go get` a stack you are not using yet.

```bash
# When you actually build an HTTP API
go get github.com/gin-gonic/gin@v1.12.0

# When env/file config is more than os.Getenv
go get github.com/spf13/viper

# When you want shorter test assertions
go get github.com/stretchr/testify

# Vulnerability scanner as a module tool (Go 1.24+)
go get -tool golang.org/x/vuln/cmd/govulncheck
```

### 4. Wire GitHub Actions

- [ ] Commit `.github/workflows/ci.yml` from [Chapter 12](./12-ci-cd.md)
- [ ] Push and wait for the first green check

---

## Daily workflow

1. **Write code** under `internal/`; keep `main` thin.
2. **Write tests** next to the code (`foo_test.go`).
3. **Run `make`** (or `make race` before a PR).
4. **Fix** lint and test failures before you switch tasks.
5. **Commit** small, reviewed diffs. If an assistant wrote the diff, you still own it.
6. **Check CI** on the pull request.

---

## The big picture

| Piece | Job |
|-------|-----|
| **mise** | Same Go and linters on every machine and in CI |
| **Go modules** | Reproducible dependencies (`go.mod` / `go.sum`) |
| **`cmd/` + `internal/`** | Thin binaries, private business logic |
| **Makefile** | One command for format, vet, lint, test, build |
| **`slog` + config** | Observable, environment-specific processes |
| **Tests + lint + `govulncheck`** | Behavior, style, and known vulns |
| **GitHub Actions** | The same checks, unattended |

---

## Example projects

- [**heroes-service/**](./heroes-service/) — the tutorial's own stdlib API (this repo).
- [**Gin CRUD Demo**](https://github.com/amitsk/gin-crud-demo) — REST API with Gin, GORM, and structured logging.
- [**gRPC View Summary**](https://github.com/amitsk/grpc-view-summary) — a small gRPC service with environment-based configuration.

---

## What now?

1. Run [`heroes-service/`](./heroes-service/) and work through [Chapter 16](./16-end-to-end.md).
2. Add request-id / health / pprof literacy in [Chapter 17](./17-observability.md).
3. Then build something of your own (to-do API, bookmarks, weather client) with the checklist above. Add a database and CI before you add a framework.

**Congratulations on completing the core tutorial.** The last two chapters are the field trip.

[16-end-to-end.md →](./16-end-to-end.md)

[← Back to [TOC](README.md#table-of-contents)]
