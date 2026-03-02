# 12. Advanced Topics

[&larr; Back to [TOC](../README.md#table-of-contents)] | [&larr; [11-ci-cd.md](../11-ci-cd.md)] | [&rarr; [13-stitching-it-together.md](../13-stitching-it-together.md)]

## Modules & Versioning

`go mod tidy`, `replace` directives, semantic imports (`github.com/user/project/v2`).

## Context

Always pass `context.Context`:
```go
func GetHero(ctx context.Context, id int) (*Hero, error)
```

## Dependency Injection

`google/wire`:
```go
// wire.go
var HeroSet = wire.NewSet(NewStore, NewService, wire.Bind(new(Store), new(*MemoryStore)))
```

`uber/fx` for apps.

## CLI Tools

`spf13/cobra`: `cobra-cli init`

`urfave/cli` simpler.

## Containers

Dockerfile:
```dockerfile
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o heroes ./cmd/heroes

FROM alpine:latest
COPY --from=builder /app/heroes /heroes
CMD [\"./heroes\"]
```

K8s basics: Deployment/Service.

## Next

[13-stitching-it-together.md &rarr;](../13-stitching-it-together.md)

[&larr; Back to [TOC](../README.md#table-of-contents)]
