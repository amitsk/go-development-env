* Generate a tutorial on setting up a Go development environment using the
details in this file. The audience for the tutorial are students or folks
learning programming with some experience in writing code. Create this as a
GitHub project with a README that has a table of contents and individual
chapters. Create navigation links in each document.

* Facts that must stay accurate (reviewed August 2026):
  - Current stable Go is 1.26.x. slog (`log/slog`) shipped in **Go 1.21**, not 1.26.
  - Official Tour: https://go.dev/tour/  (not tour.golang.org)
  - Effective Go: https://go.dev/doc/effective_go
  - Official layout guide: https://go.dev/doc/modules/layout
    (`cmd/` + `internal/` for servers; do not present
    github.com/golang-standards/project-layout as official)
  - mise: https://github.com/jdx/mise  https://mise.jdx.dev/
    Installer: `curl https://mise.run | sh` then activate the shell.
    Default project file is `mise.toml` (`.mise.toml` still works).
    `mise exec -- <cmd>` runs a command; `mise run <task>` runs a named task.
  - golangci-lint is on v2 (not v1 config files).
  - GitHub Actions: `actions/checkout@v5`+ and `jdx/mise-action@v3`+
  - Prefer `mise exec -- go …` in Makefiles, not `mise run go …`.

* Tutorial on setting up a Go Development environment
** Topics
*** Introduction
    **** Background. Not a Go tutorial. Provide links to the official Go docs
         and popular tutorials (https://go.dev/tour/, Effective Go, How to Write Go Code).
    **** Explain basic software engineering best practices like version control,
         linting, formatting, unit testing, CI/CD.
    **** Set the context about `mise`. Mention that `mise` will be used; link to
         https://mise.jdx.dev/ and installation instructions for Mac/Windows/Linux
         (`curl https://mise.run | sh`, brew, winget). Emphasize shell activation.
    **** Explain how VS Code + the official Go extension (`gopls`) will be used.
         Do not recommend Prettier for Go.
    **** Add a note that the tools mentioned are some of the popular modern
         choices and that there are many other excellent tools available
         (Go modules, go vet, golangci-lint, gofmt, goimports, testify, etc.)
         – the ones covered provide a solid foundation but users should pick
         what works best for their project.

*** Go workspace & modules
    **** What is a module? `go mod init`. GOPATH mode is obsolete.
    **** `go mod init` on Go 1.26 writes `go 1.25.0` by design; `go get go@1.26.0`
         if you want the current language version.
    **** `mise use go@1.26` creates `mise.toml`. `mise exec -- go build` vs
         `mise run <task>`.
    **** Working with `go install`, `go run`, `go test`, `go mod tidy`, `GOPRIVATE`.

*** Error handling (The Go Way)
    **** Explain the (value, error) return pattern.
    **** Why Go doesn't use try/catch like other languages—benefits for sophomores (explicit, composable).
    **** Beginner-friendly explanation of checking for `nil`.
    **** Standard pattern for bubbling up errors: `fmt.Errorf("…: %w", err)`
    **** `errors.Is`, `errors.As`, `errors.AsType` (Go 1.26).
    **** Comparison table: Go errors vs exceptions (Java/Python).
    **** Exercises: Refactor ignored errors.

*** Package organization
    **** Official layout first (https://go.dev/doc/modules/layout): start flat,
         then `internal/`, then `cmd/` for servers.
    **** How to structure for libraries vs applications—pros/cons.
    **** Naming conventions, keeping packages small and focused (one concern).
    **** Example: simple “heroes” service with handlers, models, store layers—full dir tree + main.go snippet.
    **** Common pitfalls: cyclic imports, god packages.
    **** Exercise: Restructure flat project into layers.

*** Makefile
    **** Why a Makefile is useful even with `go` commands. Mention mise `[tasks]`
         as an alternative (pick one).
    **** Sample `Makefile` targets: `build`, `run`, `test`, `race`, `lint`, `fmt`,
         `vet`, `vuln`, `tidy`, `clean`.
    **** Integration with `mise exec --` to ensure environment correctness.

*** Logging
    **** Why `fmt.Println` isn’t enough for production.
    **** `slog` – structured logger in the standard library since **Go 1.21**.
         ***** Basic setup, severity levels, JSON/text output, `With` / `WithGroup`.
         ***** `slog.NewMultiHandler` (Go 1.26). Attach `err` as a field.
    **** Alternatives: zap, zerolog (not "zed").
    **** Comparison table: stdlib `log` vs `slog` vs zap/zerolog.
    **** Recommendations for different project types.

*** Configuration management
    **** Why external configuration matters (12-factor app).
    **** Start with `os.Getenv`; mention caarlos0/env.
    **** `viper` – reading JSON/YAML/TOML/env files.
         ***** Example config struct, `BindEnv` (AutomaticEnv does not map
               DATABASE_URL → database_url by itself).
    **** Alternatives & when to choose them.

*** Database access
    **** Landscape: database/sql, sqlx, sqlc, ORMs (GORM, Ent).
    **** `gorm` – full-featured ORM. Always check `.Error`. AutoMigrate is for
         local/dev; use golang-migrate or goose in production.
    **** `sqlx` – lightweight extension over database/sql.
    **** `sqlc` – generate type-safe Go from SQL; common production choice.
    **** When to prefer each. Connection pooling. context on queries.

*** REST APIs
    **** `net/http` ServeMux (Go 1.22+ method + path values) first.
    **** `gin` – high-performance HTTP framework (v1.12+).
         ***** Routing, middleware, binding/validation, JSON responses.
    **** `echo` / `chi` as alternatives.
    **** Testing handlers with `httptest`.
    **** OpenAPI later (swaggo/swag) — not day one.

*** Unit testing & tools
    **** `go test` basics, table-driven tests, `t.Run`.
    **** `testify` and other helpers.
    **** Coverage, `-race`, fuzzing, benchmarks, httptest.

*** Static analysis & linting
    **** `go vet`, `go fix` (modernizers in Go 1.26), golangci-lint **v2**.
    **** `gofmt`/`goimports` for formatting. No Prettier.
    **** `govulncheck`.
    **** Running these via `Makefile` and CI.

*** CI/CD
    **** Explain CI/CD benefits.
    **** GitHub Actions; sample workflow using `jdx/mise-action` so CI reads
         `mise.toml`. Same `make` targets as local.

*** AI-Powered Development (LLMs)
    **** Introduction to LLMs in coding (Claude, GPT, Gemini, Grok).
    **** Benefits: Error handling, boilerplate, unit tests.
    **** Tools: Cursor, Copilot, terminal agents (Claude Code and peers).
    **** Best practices: Verification, context, learning with AI.
    **** Security note.

*** Advanced topics (glossed)
    **** Modules & versioning, `go get -tool` / `go tool` (Go 1.24+), `go work`.
    **** Using `context.Context` properly; graceful shutdown.
    **** Dependency injection: constructors first; fx optional; Wire is
         maintenance-mode — do not recommend for new projects.
    **** Building CLI tools with `cobra` / `urfave/cli` / stdlib `flag`.
    **** Docker: cached `go mod download`, `CGO_ENABLED=0`, non-root user.

* Stitch it all together.
    **** Explain how to create a new `mise` project and add dependencies.
    **** Create a comprehensive Makefile integrating all tools: build, test, lint, format, vet, race, vuln, clean.
    **** Step through the full lifecycle: write code, run `make`, build.
    **** Create a GitHub repo, push code.
    **** Configure GitHub Actions to run `make` targets on push/PR.

* End-to-end sample (`heroes-service/`)
    **** A real module in the repo: cmd/, internal/{api,config,models,store}.
    **** Stdlib only (net/http ServeMux, slog, encoding/json, httptest).
    **** Chapter that walks files in order and maps them back to earlier chapters.

* Observability
    **** Request IDs on slog + X-Request-ID.
    **** /healthz vs /readyz.
    **** pprof behind PPROF=1.
    **** Mention OpenTelemetry only as the next step for multi-process traces.
    **** GOPRIVATE / GONOSUMDB sidebar in the modules chapter.
