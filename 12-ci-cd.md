# 12. CI/CD (Automation on GitHub)

In this chapter we run the same checks on every push.

[← Back to [TOC](README.md#table-of-contents)] | [← [11-static-analysis-linting.md](./11-static-analysis-linting.md)] | [→ [13-llms.md](./13-llms.md)]

## What is CI/CD?

- **Continuous Integration (CI)** — automatically test, vet, and lint every push and pull request so broken code never sits on the main branch unnoticed.
- **Continuous Delivery / Deployment (CD)** — automatically build a release artifact (and sometimes deploy it).

This chapter is CI. CD comes later, once you have somewhere to deploy.

---

## GitHub Actions

[**GitHub Actions**](https://github.com/features/actions) runs workflows from YAML files in `.github/workflows/`.

### Example: `.github/workflows/ci.yml`

```yaml
name: CI

on:
  push:
    branches: [main]
  pull_request:

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v6

      # Installs mise, then every tool in mise.toml (Go, golangci-lint, …)
      - uses: jdx/mise-action@v4

      - name: Lint
        run: make lint

      - name: Vet
        run: make vet

      - name: Test
        run: make test

      - name: Build
        run: make build
```

Pin major versions of actions (`@v4`, `@v6`); do not use `@master`. The [mise CI docs](https://mise.jdx.dev/continuous-integration.html) and the [mise-action README](https://github.com/jdx/mise-action) list the current tags — bump when you read their release notes. `actions/checkout` also publishes v5 and v7.

---

## Why mise in CI?

The whole point of `mise.toml` is that **laptop and CI install the same versions**. If you instead write `go-version: '1.26.5'` in `actions/setup-go` and `go = "1.26"` in mise, those two files will drift.

mise-action:

1. Installs mise
2. Reads `mise.toml`
3. Puts `go` and `golangci-lint` on `PATH`
4. Caches the installs so the next run is faster

Then `make test` is the same command you run locally.

This tutorial repo ships a working workflow: [`.github/workflows/ci.yml`](./.github/workflows/ci.yml) sets `working-directory: heroes-service` so the docs at the root are not treated as a Go module.

---

## A few upgrades when you are ready

- **`go test -race`** on a separate job (slower; still worth it on `main`).
- **`govulncheck ./...`** as its own step (Chapter 11).
- **Dependabot or Renovate** for `go.mod` and GitHub Actions versions.
- **CD**: build a container (Chapter 14) and push it to a registry on tagged releases only.

## Next step

Let's talk about using AI assistants without skipping review.

[13-llms.md →](./13-llms.md)

[← Back to [TOC](README.md#table-of-contents)]
