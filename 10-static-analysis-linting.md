# 10. Static Analysis & Linting

[&larr; Back to [TOC](../README.md#table-of-contents)] | [&larr; [09-unit-testing.md](../09-unit-testing.md)] | [&rarr; [11-ci-cd.md](../11-ci-cd.md)]

## Built-ins

`go fmt ./...` / `goimports -w ./...` (formatting)

`go vet ./...` (simple checks)

## golangci-lint

`mise install golangci-lint@latest`

`.golangci.yml`:
```yaml
linters-settings:
  errcheck:
    check-type-assertions: true
  staticcheck:
    checks: [\"all\",\"-SA1019\"]
issues:
  exclude-use-default: false
```

Run: `golangci-lint run`

Bundles: staticcheck, errcheck, govet, ineffassign.

## Makefile

```
lint:
	golangci-lint run --fix
fmt:
	gofmt -w -s . && goimports -w .
```

## CI Integration

Fail on lint errors.

**Pro tip**: `golangci-lint run --fast` for quick checks.

## Next

[11-ci-cd.md &rarr;](../11-ci-cd.md)

[&larr; Back to [TOC](../README.md#table-of-contents)]
