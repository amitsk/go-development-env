# 13. Stitching It All Together

[&larr; Back to [TOC](../README.md#table-of-contents)] | [&larr; [12-advanced-topics.md](../12-advanced-topics.md)]

## Create New Project

```bash
mkdir heroes-api && cd heroes-api
mise init
go mod init github.com/yourusername/heroes-api
mise add github.com/gin-gonic/gin github.com/spf13/viper gorm.io/gorm gorm.io/driver/sqlite github.com/rs/zerolog golangci-lint
```

## Comprehensive Makefile

Copy from Ch.4, add:
```
docker-build:
	docker build -t heroes-api .
```

## Full Lifecycle

1. **Structure**: `mkdir -p cmd/heroes internal/{api,models,store,config,logger}`

2. **Config**: `internal/config/config.go` (Viper from Ch.6)

3. **Models/Store**: GORM/sqlx from Ch.7

4. **Logger**: slog/zerolog wrapper in `internal/logger/logger.go`

5. **API**: Gin router in `internal/api/handlers.go`, wire deps.

6. **Main**: `cmd/heroes/main.go`
```go
package main

import (
    \"log\"
    \"yourproject/internal/config\"
    \"yourproject/internal/logger\"
    \"yourproject/cmd/heroes/internal/server\"
)

func main() {
    cfg, _ := config.Load()
    logger.Init(cfg.LogLevel)
    srv := server.New(cfg)
    log.Fatal(srv.Start())
}
```

7. **Test/Lint/Build**:
```bash
make deps fmt lint test build run
```

8. **GitHub**: `git init`, push, add `.github/workflows/ci.yml`

9. **Deploy**: Docker build/push, or `go build` binary.

**Run**: `curl http://localhost:8080/heroes` → JSON list!

## Cookiecutter Alternative

Use [cookiecutter-go](https://github.com/... ) for templates.

## Done!

Production-ready Go app with best practices.

[&larr; Back to [TOC](../README.md#table-of-contents)]
