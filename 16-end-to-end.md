# 16. The Heroes Service (End to End)

Chapters 3–15 introduced pieces in isolation. This chapter walks the **same ideas in one module** you can run: [`heroes-service/`](./heroes-service/).

[← Back to [TOC](README.md#table-of-contents)] | [← [15-stitching-it-together.md](./15-stitching-it-together.md)] | [→ [17-observability.md](./17-observability.md)]

## What you get

A stdlib HTTP API (no Gin, no Viper, no database) so it builds after Chapter 2:

| Method | Path | Meaning |
|--------|------|---------|
| `GET` | `/healthz` | Process is up |
| `GET` | `/heroes` | List heroes |
| `GET` | `/heroes/{id}` | Fetch one |
| `POST` | `/heroes` | Create `{"name":"…"}` |

```text
heroes-service/
├── cmd/heroes/main.go          # config, logger, listen, shutdown
├── internal/
│   ├── api/                    # HTTP + middleware + httptest
│   ├── config/                 # env vars only
│   ├── models/                 # Hero + ValidateName
│   └── store/                  # interface + in-memory impl
├── Makefile
├── mise.toml
├── Dockerfile
└── go.mod
```

That is the Chapter 4 layout, filled in.

## Run it

From the **module** directory, not the tutorial root:

```bash
cd heroes-service
go test ./...
go run ./cmd/heroes
```

In another terminal:

```bash
curl -s localhost:8080/healthz
curl -s -X POST localhost:8080/heroes \
  -H 'Content-Type: application/json' \
  -d '{"name":"Superman"}'
curl -s localhost:8080/heroes
```

`Ctrl+C` triggers graceful shutdown (`signal.NotifyContext` in `cmd/heroes/main.go`).

Or `make` — tidy, fmt, vet, test, build — the same targets Chapter 5 and Chapter 12 describe.

## How the pieces map to earlier chapters

Read these files in this order. Each is short.

### 1. Errors and a store (Chapters 3–4)

[`internal/store/store.go`](./heroes-service/internal/store/store.go) defines sentinels (`ErrNotFound`, `ErrInvalidName`). [`memory.go`](./heroes-service/internal/store/memory.go) wraps them with `%w` so handlers can `errors.Is`. Tests in [`memory_test.go`](./heroes-service/internal/store/memory_test.go) check that.

The `Store` interface is the seam: `main` injects `store.NewMemory()`. A later SQL implementation would satisfy the same interface (Chapter 8) without touching HTTP.

### 2. Config from the environment (Chapter 7)

[`internal/config/config.go`](./heroes-service/internal/config/config.go) is the “you might not need Viper” path: `PORT`, `LOG_LEVEL`, `PPROF`, `SHUTDOWN_TIMEOUT`. Invalid values fail at startup — fail fast.

### 3. HTTP (Chapter 9)

[`internal/api/server.go`](./heroes-service/internal/api/server.go) uses Go 1.22 `ServeMux` patterns (`GET /heroes/{id}`, `r.PathValue`). JSON encoding is `encoding/json`. Unknown JSON fields are rejected (`DisallowUnknownFields`).

`cmd/heroes` stays thin: load config, build the logger, construct `api.NewServer`, listen, shut down on signal (Chapter 14).

### 4. Tests (Chapter 10)

| File | Kind |
|------|------|
| `internal/models/hero_test.go` | Table-driven + a fuzz target |
| `internal/store/memory_test.go` | `errors.Is` on the store |
| `internal/config/config_test.go` | `t.Setenv` |
| `internal/api/server_test.go` | `httptest` — no TCP port |

```bash
go test ./...
go test -race -count=1 ./...
go test -fuzz=FuzzValidateName -fuzztime=5s ./internal/models
```

### 5. CI (Chapter 12)

The workflow at [`.github/workflows/ci.yml`](./.github/workflows/ci.yml) runs those `make` targets inside `heroes-service/` with `jdx/mise-action`, so the Go version comes from [`mise.toml`](./mise.toml).

## Try this

1. Add `PUT /heroes/{id}` (replace the name). Handle not-found vs bad JSON vs short name.
2. Swap `Memory` for a file-backed store without changing `api`.
3. Point `LOG_LEVEL=debug` and watch JSON logs while you curl.

## Next step

The service already logs requests and exposes `/healthz`. Chapter 17 explains *why*, and how to add pprof and trace IDs without a full APM install.

[17-observability.md →](./17-observability.md)

[← Back to [TOC](README.md#table-of-contents)]
