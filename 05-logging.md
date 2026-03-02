# 05. Logging

[&larr; Back to [TOC](../README.md#table-of-contents)] | [&larr; [04-makefile.md](../04-makefile.md)] | [&rarr; [06-configuration-management.md](../06-configuration-management.md)]

## Why Not fmt.Println?

Production needs: levels, structured data, rotation, multi-sink.

## slog (Go 1.21+)

Structured, stdlib.

```go
package main

import (
    \"log/slog\"
    \"os\"
)

func main() {
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
    logger.Info(\"hero created\", slog.Int(\"id\", 1), slog.String(\"name\", \"Superman\"))
}
```

Levels: Debug, Info, Warn, Error. Context, attrs.

Global: `slog.SetDefault(logger)`

## Zerolog (\"zed\"-like high-perf)

`go get github.com/rs/zerolog`

```go
import \"github.com/rs/zerolog\"

logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
logger.Info().Int(\"id\", 1).Str(\"name\", \"Batman\").Msg(\"hero updated\")
```

Sampling, hooks.

## Comparison

| Feature | slog | zerolog |
|---------|------|---------|
| Stdlib  | Yes  | No     |
| Perf    | Good | Best   |
| JSON    | Built-in | Built-in |
| Hooks   | Yes  | Yes    |

**Rec**: slog for new, zerolog for perf-critical.

## Next

[06-configuration-management.md &rarr;](../06-configuration-management.md)

[&larr; Back to [TOC](../README.md#table-of-contents)]
