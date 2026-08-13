# 14. Advanced Topics

You have the foundations of a professional Go environment. This chapter is a map of what comes next — not a full course.

[← Back to [TOC](README.md#table-of-contents)] | [← [13-llms.md](./13-llms.md)] | [→ [15-stitching-it-together.md](./15-stitching-it-together.md)]

## 1. Modules, versions, and `go tool`

Release libraries with [semantic versioning](https://semver.org/) and Git tags (`v1.2.3`). A breaking change to an already-released `v1` module needs a `/v2` import path. See [Module version numbering](https://go.dev/doc/modules/version-numbers).

Since **Go 1.24**, executable tools belong in `go.mod` with a `tool` directive — no more `tools.go` blank imports:

```bash
go get -tool golang.org/x/vuln/cmd/govulncheck
go get -tool golang.org/x/tools/cmd/stringer
go tool govulncheck ./...
```

For a **multi-module repo**, look at [`go work`](https://go.dev/ref/mod#workspaces) (`go.work`). You do not need it for the single-module heroes service.

---

## 2. `context.Context`

Any function that waits — HTTP, database, gRPC, a sleep — should take `context.Context` as its **first** argument.

```go
func FetchHero(ctx context.Context, id int) (*Hero, error) {
    row := db.QueryRowContext(ctx, `SELECT id, name FROM heroes WHERE id = $1`, id)
    // ...
}
```

The context carries deadlines and cancellation. When the client hangs up, `r.Context()` is cancelled and a well-behaved query stops. Do not store a context on a struct; pass it down the call.

In `main`, pair the server with a signal:

```go
ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
defer stop()

go func() {
    <-ctx.Done()
    shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    _ = server.Shutdown(shutdownCtx)
}()
```

---

## 3. Dependency injection (without a framework)

"Dependency injection" means: **pass collaborators in**, do not construct them deep in the stack.

```go
type API struct {
    store  store.Store
    logger *slog.Logger
}

func NewAPI(s store.Store, logger *slog.Logger) *API {
    return &API{store: s, logger: logger}
}
```

`main` is the only place that opens the database and builds the graph. Tests pass a fake `store.Store`.

Frameworks such as [Uber fx](https://github.com/uber-go/fx) help in very large graphs. [Google Wire](https://github.com/google/wire) is in maintenance mode — do not start a new project on it. Most services never outgrow constructors.

---

## 4. CLIs

Go is a default choice for command-line tools. [**Cobra**](https://github.com/spf13/cobra) and [**urfave/cli**](https://github.com/urfave/cli) handle flags, subcommands, and help text. For a single command, `flag` in the standard library is enough.

---

## 5. Containers

[`heroes-service/Dockerfile`](./heroes-service/Dockerfile) is this image applied to the sample.

A small, reproducible image:

```dockerfile
# syntax=docker/dockerfile:1

FROM golang:1.26-alpine AS builder
WORKDIR /src
RUN apk add --no-cache ca-certificates
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/heroes ./cmd/heroes

FROM alpine:3.21
RUN adduser -D -H -u 65532 app
COPY --from=builder /out/heroes /heroes
USER app
EXPOSE 8080
ENTRYPOINT ["/heroes"]
```

Why the extra pieces:

- **`go mod download` before `COPY .`** — dependency layers stay cached when only app code changes.
- **`CGO_ENABLED=0`** — a static binary, no glibc surprises.
- **non-root `USER`** — the process cannot write wherever it wants if the container is broken into.

For a tighter runtime, look at [distroless](https://github.com/GoogleContainerTools/distroless) `static` images. Kubernetes is a deployment *platform*; learn Docker (or another OCI builder) first.

---

## 6. Goroutines and channels

Go's concurrency primitives are cheap goroutines and typed channels. The [Tour](https://go.dev/tour/concurrency/1) is still the right first hour.

```go
func main() {
    ball := make(chan string)
    go player("ping", ball)
    go player("pong", ball)

    ball <- "Serve!"
    time.Sleep(time.Second)
    <-ball
}

func player(name string, ch chan string) {
    for {
        msg := <-ch
        fmt.Println(name, msg)
        time.Sleep(100 * time.Millisecond)
        ch <- fmt.Sprintf("%s: ball!", name)
    }
}
```

Always run `go test -race` once you share memory across goroutines. Prefer communicating over a channel, or protecting the memory with `sync.Mutex`, over "I think this is safe."

Further reading: [Go Concurrency Patterns](https://go.dev/blog/pipelines), [Effective Go — concurrency](https://go.dev/doc/effective_go#concurrency).

## Next step

Put the pieces into one new-project checklist.

[15-stitching-it-together.md →](./15-stitching-it-together.md)

[← Back to [TOC](README.md#table-of-contents)]
