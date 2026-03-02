# 02. Go Workspace & Modules

[&larr; Back to [TOC](../README.md#table-of-contents)] | [&larr; [01-introduction.md](../01-introduction.md)] | [&rarr; [03-package-organization.md](../03-package-organization.md)]

## What is a Go Module?

Go modules manage dependencies since Go 1.11. Run `go mod init github.com/yourusername/project`.

Check GOPATH: `go env GOPATH` (legacy, modules replace it).

## Mise for Go Environments

**mise** handles Go versions/tools like pyenv/uv.

```bash
# Install Go 1.23
mise use go@1.23
mise install go@1.23  # First time

# Project-local .mise.toml or mise.toml
echo 'tasks:
  default:
    run: go run .
  test:
    run: go test ./...
' > mise.toml
```

`mise run go build` ensures correct Go version.

## Working with Go Commands

```bash
go mod init example.com/heroes
go mod tidy  # Add deps
go run .     # Run main
go test ./... # Tests
go install   # Binary to $GOBIN
```

**Pro tip**: `mise activate` for shell integration.

## Example

Create `main.go`:
```go
package main

import \"fmt\"

func main() {
    fmt.Println(\"Hello, Go modules!\")
}
```

Run: `mise run go run .`

## Next

[03-package-organization.md &rarr;](../03-package-organization.md)

[&larr; Back to [TOC](../README.md#table-of-contents)]
