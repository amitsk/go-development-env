* Generate a tutorial on setting up a Go development environment using the
details in this file. The audience for the tutorial are students or folks
learning programming with some experience in writing code. Create this as a
GitHub project with a README that has a table of contents and individual
chapters. Create navigation links in each document.

* Tutorial on setting up a Go Development environment
** Topics
*** Introduction
    **** Background. Not a Go tutorial. Provide links to the official Go docs
         and popular tutorials (tour.golang.org, effectivego.dev).
    **** Explain basic software engineering best practices like version control,
         linting, formatting, unit testing, CI/CD.
    **** Set the context about `mise`. Mention that `mise` will be used; link to
         https://github.com/mitchellh/mise and installation instructions for
         Mac/Windows/Linux.
    **** Explain how VS Code will be used. Point to VS Code GitHub link.
    **** Add a note that the tools mentioned are some of the popular modern
         choices and that there are many other excellent tools available
         (Go modules, govet, golangci-lint, gofmt, goimports, testify, etc.)
         – the ones covered provide a solid foundation but users should pick
         what works best for their project.

*** Go workspace & modules
    **** What is a module? `go mod init`/`go env GOPATH`/`mise init`.
    **** `mise` project creation and dependency management (`mise add`,
         `mise run go build`).
    **** Working with `go install`, `go run`, `go test`.

*** Error handling (The Go Way)
    **** Explain the \"tuple\" return pattern (value, error).
    **** Why Go doesn't use try/catch like other languages—benefits for sophomores (explicit, composable).
    **** Beginner-friendly explanation of checking for `nil`.
    **** Standard pattern for bubbling up errors: if err != nil { return err }
    **** `errors.Is`, `errors.As`, `fmt.Errorf` wrapping.
    **** Comparison table: Go errors vs exceptions (Java/Python).
    **** Exercises: Refactor fmt.Println error ignoring.

*** Package organization
    **** Idiomatic layout (`cmd/`, `pkg/`, `internal/`, `api/`, `configs/`,
         `scripts/`, `test/`).
    **** How to structure for libraries vs applications—pros/cons.
    **** Naming conventions, keeping packages small and focused (one concern).
    **** Example: simple “heroes” service with handlers, models, store layers—full dir tree + main.go snippet.
    **** Common pitfalls: cyclic imports, god packages.
    **** Exercise: Restructure flat project into layers.

*** Makefile
    **** Why a Makefile is useful even with `go` commands.
    **** Sample `Makefile` targets: `build`, `run`, `test`, `lint`, `fmt`,
         `vet`, `deps`, `clean`.
    **** Integration with `mise` to ensure environment correctness.

*** Logging
    **** Why `fmt.Println` isn’t enough for production.
    **** `slog` – the new structured logger in Go 1.26.
         ***** Basic setup, severity levels, context, JSON/text output.
         ***** Adding hooks/sinks, global logger vs per-component loggers.
    **** `zed` – alternative structured logging library.
         ***** Example configuration, using contexts, sampling.
    **** Comparison table: stdlib `log` vs `slog` vs `zed`.
    **** Recommendations for different project types.

*** Configuration management
    **** Why external configuration matters (12-factor app).
    **** `viper` – reading JSON/YAML/TOML/env files.
         ***** Example config struct, `viper.AutomaticEnv()`, watching for
               changes.
    **** Alternatives & when to choose them.

*** Database access
    **** Landscape: database/sql, ORMs, query builders.
    **** `gorm` – full-featured ORM.
         ***** Defining models, migrations, associations, CRUD examples.
    **** `sqlx` – lightweight extension over database/sql.
         ***** Named queries, struct scans, transaction examples.
    **** When to prefer `gorm` vs `sqlx` vs plain `database/sql`.
    **** Brief mention of migrations (go-migrate, ent) and connection pooling.

*** REST APIs
    **** `gin` – high-performance HTTP framework.
         ***** Routing, middleware, binding/validation, JSON responses.
         ***** Example CRUD API, custom error handling.
    **** `echo` – another popular router/HTTP server.
         ***** Comparison with `gin` (performance, API, middleware).
         ***** Example with middleware, websockets.
    **** Testing handlers with `httptest`.
    **** OpenAPI/Swagger generation (swaggo/swag, echo-swagger).

*** Unit testing & tools
    **** `go test` basics, table-driven tests.
    **** `testify` and other helpers.
    **** Coverage, benchmarks, fuzzing.

*** Static analysis & linting
    **** `go vet`, `golangci-lint` (linters bundled: `errcheck`, `staticcheck`,
         `ineffassign`, etc.).
    **** `gofmt`/`goimports` for formatting.
    **** Running these via `Makefile` and CI.

*** CI/CD
    **** Explain CI/CD benefits.
    **** Short explainer about GitHub Actions; sample workflow building,
         testing, linting, releasing a module.

*** AI-Powered Development (LLMs)
    **** Introduction to LLMs in coding (GPT-4, Claude).
    **** Benefits: Error handling, boilerplate, unit tests.
    **** Tools: Cursor, Claude Code, Copilot.
    **** Best practices: Verification, context, learning with AI.
    **** Security note.

*** Advanced topics (glossed)
    **** Modules & versioning (`go mod tidy`, `replace`, semantic import
         versioning).
    **** Using `context.Context` properly.
    **** Dependency injection patterns (wire, fx).
    **** Building CLI tools with `cobra` and using `urfave/cli`.
    **** Building and deploying with Docker, Kubernetes basics.

* Stitch it all together.
    **** Explain how to create a new `mise` project and add dependencies.
    **** Create a comprehensive Makefile integrating all tools: build, test, lint, format, deps, clean.
    **** Step through the full lifecycle: write code, run `make test lint fmt`, build/deploy.
    **** Create a GitHub repo, push code.
    **** Configure GitHub Actions to run `make` targets on push/PR.
