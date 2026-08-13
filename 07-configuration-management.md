# 07. Configuration Management

In this chapter we keep settings out of the source code.

[← Back to [TOC](README.md#table-of-contents)] | [← [06-logging.md](./06-logging.md)] | [→ [08-database-access.md](./08-database-access.md)]

## Why external configuration?

A professional app runs in more than one place: your laptop (development), a shared test environment (staging), and production.

The [Twelve-Factor App](https://12factor.net/config) rule is: **never hardcode** config (ports, database URLs, API keys). Keep it outside the binary so you can:

1. **Keep secrets out of Git** — no passwords in the repo.
2. **Change behavior without rebuilding** — flip the log level or database URL.
3. **Stay environment-agnostic** — a different database on your laptop than in production.

---

## Do you need a library?

Many Go services only need environment variables. The standard library is enough:

```go
port := os.Getenv("PORT")
if port == "" {
    port = "8080"
}
```

Small typed helpers such as [caarlos0/env](https://github.com/caarlos0/env) or [kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig) stay close to that idea.

Use a fuller library when you genuinely need files + env + flags merged together. That is where Viper earns its keep.

[`heroes-service/internal/config`](./heroes-service/internal/config/config.go) is the env-only path: `PORT`, `LOG_LEVEL`, `PPROF`.

---

## Viper

[**Viper**](https://github.com/spf13/viper) can read:

- Files (YAML, JSON, TOML, envfile, …)
- Environment variables
- Command-line flags
- Remote sources (Consul, etcd) — skip these until you need them

### 1. Define a config struct

```go
type Config struct {
    Port     int    `mapstructure:"port"`
    Database string `mapstructure:"database_url"`
    LogLevel string `mapstructure:"log_level"`
}
```

### 2. Load it

```go
package main

import (
    "fmt"
    "strings"

    "github.com/spf13/viper"
)

func LoadConfig() (*Config, error) {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    viper.AddConfigPath("./configs")

    viper.SetDefault("port", 8080)
    viper.SetDefault("log_level", "info")

    // PORT → port, DATABASE_URL → database_url
    viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
    viper.AutomaticEnv()

    // AutomaticEnv does not invent names. Bind the ones you care about:
    _ = viper.BindEnv("database_url", "DATABASE_URL")
    _ = viper.BindEnv("port", "PORT")
    _ = viper.BindEnv("log_level", "LOG_LEVEL")

    if err := viper.ReadInConfig(); err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
            return nil, fmt.Errorf("read config: %w", err)
        }
        // File is optional if env vars provide the required values.
    }

    var cfg Config
    if err := viper.Unmarshal(&cfg); err != nil {
        return nil, fmt.Errorf("unmarshal config: %w", err)
    }
    if cfg.Database == "" {
        return nil, fmt.Errorf("database_url is required")
    }
    return &cfg, nil
}
```

`AutomaticEnv` alone will **not** map `DATABASE_URL` onto `database_url`. You need `BindEnv` (or a naming scheme that already matches). This is the most common Viper surprise.

### 3. Example file (`config.yaml` or `configs/config.yaml`)

```yaml
port: 8080
database_url: "postgres://localhost/heroes"
log_level: "info"
```

Commit an example (`config.example.yaml`) with fake values. Put the real `config.yaml` in `.gitignore` if it can contain secrets. Prefer env vars for anything sensitive.

---

## Best practices

- **Secrets live in the environment** (or a secret manager), never in Git.
- **Set defaults** for safe local values (`port`, `log_level`).
- **Fail fast** if a required value (database URL, signing key) is missing.
- **Don't watch files in production** unless you have a concrete hot-reload need. Restarting the process is simpler and safer.

## Next step

Now that we can configure the app, let's store and retrieve data.

[08-database-access.md →](./08-database-access.md)

[← Back to [TOC](README.md#table-of-contents)]
