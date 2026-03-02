# 03. Package Organization

[&larr; Back to [TOC](../README.md#table-of-contents)] | [&larr; [02-go-workspace-modules.md](../02-go-workspace-modules.md)] | [&rarr; [04-makefile.md](../04-makefile.md)]

## Idiomatic Go Layout

Keep packages small, focused. Standard dirs:

```
project/
├── cmd/          # Entry points (main packages)
│   └── heroes/
│       └── main.go
├── internal/     # Private code
│   ├── api/
│   ├── models/
│   └── store/
├── pkg/          # Reusable public code
├── configs/      # Config files
├── scripts/      # Build/deploy scripts
├── test/         # Integration tests
├── go.mod
├── Makefile
└── README.md
```

## Libraries vs Apps

- **App**: `cmd/myapp/main.go` imports `internal/`.
- **Library**: Public APIs in `pkg/`.

Naming: lowercase, short, descriptive (e.g., `heroesv1` for API v1).

## Heroes Service Example

`internal/models/hero.go`:
```go
package models

type Hero struct {
    ID   int    `json:\"id\"
    Name string `json:\"name\"
}
```

`internal/store/store.go`:
```go
package store

import \"yourproject/internal/models\"

type Store interface {
    GetHeroes() ([]models.Hero, error)
}
```

`cmd/heroes/main.go`:
```go
package main

import (
    \"fmt\"
    \"yourproject/internal/store\"
)

func main() {
    s := store.NewMemoryStore()
    heroes, _ := s.GetHeroes()
    fmt.Println(heroes)
}
```

## Best Practices

- One responsibility per package.
- Avoid deep nesting.
- Use `internal/` for unexported.

## Next

[04-makefile.md &rarr;](../04-makefile.md)

[&larr; Back to [TOC](../README.md#table-of-contents)]
