# 17. Observability (Just Enough)

You cannot debug a process you cannot see. This chapter covers the three signals a small Go service should have on day one — **logs, health, profiles** — using only the standard library. OpenTelemetry is named at the end; you do not need it yet.

The running example is still [`heroes-service/`](./heroes-service/).

[← Back to [TOC](README.md#table-of-contents)] | [← [16-end-to-end.md](./16-end-to-end.md)]

## 1. Structured logs with a request id

Every log line about a request should share one id so you can grep a single curl through the stack.

The sample does three things in [`internal/api/middleware.go`](./heroes-service/internal/api/middleware.go):

1. Honor `X-Request-ID` if the client (or a proxy) sent one; otherwise generate one.
2. Echo it on the response (`X-Request-ID`).
3. Store it on `r.Context()` and include `request_id` on the access log.

```go
logger.Info("request",
    slog.String("request_id", RequestIDFrom(r.Context())),
    slog.String("method", r.Method),
    slog.String("path", r.URL.Path),
    slog.Int("status", sw.code),
    slog.Duration("duration", time.Since(start)),
)
```

Errors in handlers use the same field ([`fail` in `server.go`](./heroes-service/internal/api/server.go)). That is Chapter 6 applied to HTTP.

Try it:

```bash
curl -s -D - -H 'X-Request-ID: lab-42' localhost:8080/healthz -o /dev/null
```

The response headers include `X-Request-ID: lab-42`, and the process prints a JSON line with `"request_id":"lab-42"`.

Do **not** log bodies, cookies, or authorization headers.

## 2. Health endpoints

Orchestrators and load balancers need a cheap “is this process alive?” check that does not hit the database.

`GET /healthz` in the sample returns `{"status":"ok"}`. That is a **liveness** probe: the HTTP server is accepting connections.

When you add a real store, add **readiness** separately:

| Path | Question | Fail when |
|------|----------|-----------|
| `/healthz` | Is the process running? | Almost never (deadlock, total crash) |
| `/readyz` | Can I take traffic? | DB ping fails, required config missing |

Readiness should not be a full integration test. One `PingContext` with a 300ms timeout is enough.

## 3. pprof

The standard library can profile a running process. The sample exposes it only when you opt in — pprof is useful and also a data leak if you leave it on a public port.

```bash
PPROF=1 go run ./cmd/heroes
```

Then, in another terminal (Go toolchain required):

```bash
go tool pprof http://localhost:8080/debug/pprof/heap
go tool pprof http://localhost:8080/debug/pprof/profile?seconds=10
```

Handlers live under `/debug/pprof/` and are registered only if `PPROF=1` (see [`NewServer`](./heroes-service/internal/api/server.go)). In production, bind pprof to localhost or protect it; do not put it on the public mux.

## 4. When to add OpenTelemetry

Logs + health + occasional pprof will carry a student project and many internal services.

Add [OpenTelemetry](https://opentelemetry.io/docs/languages/go/) when you have **more than one process** and you need to follow one request across them (API → worker → database). The usual shape:

- A **trace id** (wider than our request id) stored on `context.Context`
- A **span** per outbound call
- Metrics (`http_server_request_duration_seconds`) scraped by Prometheus or pushed to your vendor

Put the trace id on the same `slog` records (`logger.With(slog.String("trace_id", …))`) so logs and traces join. Do not start with a vendor agent “just in case.”

The Go 1.26 experimental `goroutineleak` pprof profile (`GOEXPERIMENT=goroutineleakprofile`) is worth knowing about once you have concurrency bugs. It is not a day-one switch.

## What “good enough” looks like

- [ ] JSON logs on stdout (Chapter 6)
- [ ] `request_id` on every request log and error
- [ ] `/healthz` (and `/readyz` once you have a dependency)
- [ ] pprof behind a flag or a private port
- [ ] `go test -race` in CI (Chapter 12) so you notice shared-memory bugs before production does

That is the whole observability budget for this tutorial. Kubernetes probes, dashboards, and SLOs come after the service is actually deployed.

[← Back to [TOC](README.md#table-of-contents)]
