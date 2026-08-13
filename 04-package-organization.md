# 04. Package Organization

Organizing your code is one of the most important skills in professional software development. This chapter covers what the Go team actually recommends, then a layout that works well for a small HTTP service.

[← Back to [TOC](README.md#table-of-contents)] | [← [03-error-handling.md](./03-error-handling.md)] | [→ [05-makefile.md](./05-makefile.md)]

## Start smaller than you think

When the project is small, **one package in the repo root is fine**. The official guide, [Organizing a Go module](https://go.dev/doc/modules/layout), is explicit about this:

- A tiny command can be `main.go` + `go.mod` at the root.
- Split into extra packages only when a file is doing two jobs.
- Prefer `internal/` for supporting packages you do not want other modules to import.

The popular [golang-standards/project-layout](https://github.com/golang-standards/project-layout) repo is **not** an official standard. Empty `pkg/`, `api/`, `scripts/`, and `configs/` folders do not make a project professional. Add a directory when you have a second file that belongs there.

---

## A practical service layout

Once you have an HTTP API, a data layer, and maybe a second binary, this shape is common and matches the official "server project" advice (`cmd/` + `internal/`):

```text
heroes-service/
├── cmd/
│   └── heroes/
│       └── main.go          # package main — wiring only
├── internal/
│   ├── api/                 # HTTP handlers
│   ├── models/              # data types
│   └── store/               # persistence
├── configs/                 # optional: YAML / example env files
├── go.mod
├── mise.toml
├── Makefile
└── README.md
```

Skip `pkg/` until you are *sure* another module should import that code. Most services never need it.

This repository's [`heroes-service/`](./heroes-service/) is that layout with real files. Chapter 16 walks it end to end.

### 1. `cmd/` — entry points

Each subdirectory is a separate binary. Keep `main` thin: parse flags, load config, construct dependencies, call `ListenAndServe`. Business logic lives under `internal/`.

### 2. `internal/` — compiler-enforced privacy

Any package under `internal/` **cannot** be imported by other modules. That is a language rule, not a convention. It lets you refactor freely.

### 3. Libraries vs applications

| | Library (importable package) | Application / server |
|--|------------------------------|----------------------|
| Root package | Public API (`package heroes`) | Usually no public Go API |
| Binaries | Optional `cmd/` | `cmd/<name>/main.go` |
| Internals | `internal/` for unexported helpers | Almost everything in `internal/` |

---

## Example: the Heroes service

### The model (`internal/models/hero.go`)

```go
package models

// Hero represents a single hero in our system.
type Hero struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}
```

### The store (`internal/store/store.go`)

```go
package store

import "github.com/yourusername/heroes-service/internal/models"

// Store is how the rest of the app reads and writes heroes.
type Store interface {
    GetHeroes() ([]models.Hero, error)
}
```

Using an interface here makes tests easy later (Chapter 10): you can swap a fake store for a database.

### The app entry point (`cmd/heroes/main.go`)

```go
package main

import (
    "fmt"

    "github.com/yourusername/heroes-service/internal/models"
)

func main() {
    h := models.Hero{ID: 1, Name: "Super Go"}
    fmt.Printf("Hello, %s!\n", h.Name)
}
```

Run it from the module root:

```bash
go run ./cmd/heroes
```

---

## Naming and pitfalls

- **Package names** are short, lowercase, singular (`store`, not `stores` or `store_utils`).
- **No "god" packages** named `util`, `common`, or `helpers`. If you cannot name the package after what it *does*, the code probably belongs next to its caller.
- **Cyclic imports** (`api` imports `store` imports `api`) mean the packages are not really separate. Extract a third package (often `models`) or merge them.
- **One package per directory.** All `.go` files in a folder share the same `package` name.

## Try this

Start from the single-file `main.go` in Chapter 2. Move it to `cmd/heroes/main.go` and put `Hero` in `internal/models`. Confirm `go run ./cmd/heroes` still works.

## Next step

Now that we know how to organize the code, let's automate the development tasks.

[05-makefile.md →](./05-makefile.md)

[← Back to [TOC](README.md#table-of-contents)]
