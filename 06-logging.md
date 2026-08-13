# 06. Structured Logging

In this chapter we record what the application is doing with **structured logging**.

[← Back to [TOC](README.md#table-of-contents)] | [← [05-makefile.md](./05-makefile.md)] | [→ [07-configuration-management.md](./07-configuration-management.md)]

## Why not just `fmt.Println`?

`fmt.Println` is fine while you learn. In production it falls short:

1. **No levels** — you cannot tell a routine "info" from a crash.
2. **Hard to search** — thousands of free-form lines do not query well.
3. **Machine-unfriendly** — dashboards and alerts want key/value fields, not English sentences.

---

## What is structured logging?

Write logs as consistent, machine-readable records (usually JSON). Instead of `User 123 logged in`:

```json
{"time":"2026-08-12T12:00:00Z","level":"INFO","msg":"user login","user_id":123}
```

---

## `log/slog` (standard library since Go 1.21)

Since **Go 1.21** (August 2023), the standard library includes `log/slog`. You do not need an extra module for structured logs on any supported Go release.

### Basic setup

```go
package main

import (
    "log/slog"
    "os"
)

func main() {
    handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
        Level: slog.LevelInfo,
    })
    logger := slog.New(handler)
    slog.SetDefault(logger)

    logger.Info("hero created",
        slog.Int("id", 1),
        slog.String("name", "Superman"),
    )
}
```

Use **`slog.NewTextHandler`** while developing (easier to read) and JSON in staging/production.

### Always attach the error as a field

```go
if err != nil {
    logger.Error("create hero failed",
        slog.String("name", name),
        slog.Any("err", err),
    )
    return err
}
```

A message of `"error: " + err.Error()` loses structure. Keep `err` as its own attribute.

### Child loggers and groups

```go
reqLog := logger.With(slog.String("request_id", reqID))
reqLog.Info("handler start")

// Groups nest related keys: db.query, db.rows
dbLog := logger.WithGroup("db")
dbLog.Info("query", slog.String("sql", "SELECT …"))
```

Pass the `*slog.Logger` (or a `context.Context` with `slog.NewContext` / `Logger.With`) into handlers and stores instead of calling a global logger from deep packages. Globals are fine at the edge (`main`).

### Multiple destinations (Go 1.26)

Go 1.26 added [`slog.NewMultiHandler`](https://pkg.go.dev/log/slog#NewMultiHandler): fan the same records out to several handlers (for example JSON to stdout *and* a file, or a test buffer).

```go
logger := slog.New(slog.NewMultiHandler(
    slog.NewJSONHandler(os.Stdout, nil),
    slog.NewTextHandler(debugFile, &slog.HandlerOptions{Level: slog.LevelDebug}),
))
```

---

## Log levels

- **DEBUG** — noisy detail for developers. Off in production by default.
- **INFO** — normal operation ("server started", "hero created").
- **WARN** — unexpected but the request or process can continue.
- **ERROR** — an operation failed.

Set the minimum level from configuration (Chapter 7), not a hard-coded constant.

---

## Alternatives

| Feature | `log` (old stdlib) | `slog` | [zap](https://github.com/uber-go/zap) / [zerolog](https://github.com/rs/zerolog) |
|---------|--------------------|--------|----------------------------------|
| In the standard library? | Yes | Yes (Go 1.21+) | No |
| Structured fields | No | Yes | Yes |
| Performance | Fine | Very good | Best-in-class |
| When to choose | Tiny CLIs | **Default for new code** | Existing codebases or extreme throughput |

**Recommendation:** use `slog` for new projects. Reach for zap or zerolog only if you already depend on them or you have measured a logging bottleneck.

Chapter 17 attaches a `request_id` to these records in [`heroes-service`](./heroes-service/internal/api/middleware.go).

## Next step

Now that we can log what the app is doing, let's configure it without changing the code.

[07-configuration-management.md →](./07-configuration-management.md)

[← Back to [TOC](README.md#table-of-contents)]
