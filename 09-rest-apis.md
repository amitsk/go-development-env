# 09. Building REST APIs

In this chapter we expose the application over HTTP.

[← Back to [TOC](README.md#table-of-contents)] | [← [08-database-access.md](./08-database-access.md)] | [→ [10-unit-testing.md](./10-unit-testing.md)]

## What is a REST API?

A **REST API** is two programs talking over HTTP — the same protocol your browser uses.

- **Resources**: the things you manage (`/heroes`).
- **Methods**:
  - `GET` — read
  - `POST` — create
  - `PUT` / `PATCH` — update
  - `DELETE` — remove

---

## Start with `net/http` (Go 1.22+)

Since **Go 1.22**, the standard library `ServeMux` understands methods and path parameters. For many services you do not need a framework:

```go
mux := http.NewServeMux()

mux.HandleFunc("GET /heroes", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(map[string]any{
        "heroes": []string{"Superman", "Batman", "Wonder Woman"},
    })
})

mux.HandleFunc("GET /heroes/{id}", func(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    fmt.Fprintln(w, id)
})

server := &http.Server{Addr: ":8080", Handler: mux}
log.Fatal(server.ListenAndServe())
```

Learn this first. Frameworks are thin wrappers around the same ideas. [`heroes-service/internal/api`](./heroes-service/internal/api/server.go) is this style with create/list/get and tests.

---

## Gin, when you want batteries included

[**Gin**](https://github.com/gin-gonic/gin) (v1.12 as of 2026) is still the most common third-party HTTP framework: routing, JSON binding, validation tags, and middleware in a small API.

```go
package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type CreateHeroRequest struct {
    Name string `json:"name" binding:"required"`
}

func main() {
    r := gin.Default()

    r.POST("/heroes", func(c *gin.Context) {
        var req CreateHeroRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusCreated, gin.H{
            "id":   1,
            "name": req.Name,
        })
    })

    r.GET("/heroes", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "heroes": []string{"Superman", "Batman", "Wonder Woman"},
        })
    })

    if err := r.Run(":8080"); err != nil {
        panic(err)
    }
}
```

Install it in the module:

```bash
go get github.com/gin-gonic/gin@v1.12.0
```

---

## JSON with the standard library

Gin's `ShouldBindJSON` uses `encoding/json`. You should know the stdlib API:

```go
hero := models.Hero{ID: 1, Name: "Superman"}
bytes, err := json.Marshal(hero) // {"id":1,"name":"Superman"}

var h models.Hero
err = json.Unmarshal(bytes, &h) // pass a pointer
```

Struct tags:

```go
type Hero struct {
    ID     int    `json:"id"`
    Name   string `json:"name,omitempty"`
    Secret bool   `json:"-"`
}
```

Pitfalls: unmarshal into a pointer; `json.NewDecoder(r.Body).Decode` is better for HTTP bodies than `io.ReadAll` + `Unmarshal` (it streams and enforces one value).

Go 1.26 also lets you write `Age: new(yearsSince(born))` when a JSON field is a pointer — handy for optional values.

---

## Key HTTP ideas

### Context

In Gin, `c *gin.Context` holds the request and lets you write the response. The standard library uses `http.ResponseWriter`, `*http.Request`, and `r.Context()` (cancellation when the client hangs up). Prefer `r.Context()` for downstream calls.

### Status codes

- `200 OK` — success
- `201 Created` — a new resource exists
- `400 Bad Request` — the client sent junk
- `404 Not Found`
- `500 Internal Server Error` — *your* bug; log it, do not leak internals to the client

### Middleware

Functions that wrap every request: request IDs, structured logs, auth, panic recovery. Gin's `gin.Default()` already installs logger + recovery. In stdlib, wrap `http.Handler`.

### Graceful shutdown

Production servers should handle `SIGINT` / `SIGTERM` and call `server.Shutdown(ctx)` so in-flight requests finish. See Chapter 14.

---

## Alternatives

[**Echo**](https://echo.labstack.com/) and [**chi**](https://github.com/go-chi/chi) are the usual other choices. chi is a thin router that feels like stdlib. Echo is closer to Gin.

**Recommendation:** learn `net/http` first, then use **Gin** or **chi** if you want binding and a middleware ecosystem. Generate OpenAPI later with [swag](https://github.com/swaggo/swag) or an API-first tool — not on day one.

Handler tests belong with Chapter 10 (`net/http/httptest`).

## Next step

Let's make sure the API actually works.

[10-unit-testing.md →](./10-unit-testing.md)

[← Back to [TOC](README.md#table-of-contents)]
