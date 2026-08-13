# Heroes service

A small REST API used by this tutorial. It is **stdlib only** (`net/http`, `log/slog`, `encoding/json`) so you can run it after Chapter 2 without extra modules.

Walkthrough: [Chapter 16](../16-end-to-end.md). Observability notes: [Chapter 17](../17-observability.md).

## Run

```bash
# from this directory
go test ./...
go run ./cmd/heroes
```

```bash
curl -s localhost:8080/healthz
curl -s -X POST localhost:8080/heroes -d '{"name":"Superman"}'
curl -s localhost:8080/heroes
```

Environment:

| Variable | Default | Meaning |
|----------|---------|---------|
| `PORT` | `8080` | Listen port |
| `LOG_LEVEL` | `info` | `debug`, `info`, `warn`, `error` |
| `PPROF` | off | Set `1` to expose `/debug/pprof/` |
| `SHUTDOWN_TIMEOUT` | `10s` | Graceful shutdown budget |

`make` runs tidy, fmt, vet, test, and build. `make lint` needs `golangci-lint` (pinned in `mise.toml`).
