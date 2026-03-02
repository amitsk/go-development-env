# 11. CI/CD

[&larr; Back to [TOC](../README.md#table-of-contents)] | [&larr; [10-static-analysis-linting.md](../10-static-analysis-linting.md)] | [&rarr; [12-advanced-topics.md](../12-advanced-topics.md)]

## What is CI/CD?

Continuous Integration/Deployment: auto-test/build/deploy on push/PR.

Benefits: catch bugs early, consistent envs.

## GitHub Actions

`.github/workflows/ci.yml`:
```yaml
name: CI
on: [push, pull_request]
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v4
    - uses: mise-action/setup@v2
      with:
        settings: |
          go: 1.23
    - run: make deps
    - run: make test lint fmt vet
    - run: make build
```

`mise-action` ensures Go version.

## Release Workflow

Add `goreleaser` for binaries.

## Next

[12-advanced-topics.md &rarr;](../12-advanced-topics.md)

[&larr; Back to [TOC](../README.md#table-of-contents)]
