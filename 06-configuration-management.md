# 06. Configuration Management

[&larr; Back to [TOC](../README.md#table-of-contents)] | [&larr; [05-logging.md](../05-logging.md)] | [&rarr; [07-database-access.md](../07-database-access.md)]

## Why External Config?

12-factor: config in env. Supports dev/prod.

## Viper

`mise add github.com/spf13/viper`

```go
package main

import \"github.com/spf13/viper\"

type Config struct {
    DBURL string `mapstructure:\"db_url\"
    LogLevel string
}

func loadConfig() (*Config, error) {
    viper.SetConfigName(\"config\")
    viper.SetConfigType(\"yaml\")
    viper.AddConfigPath(\".\")
    viper.AutomaticEnv() // Overrides

    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }

    cfg := &Config{}
    viper.Unmarshal(cfg)
    return cfg, nil
}
```

`config.yaml`:
```yaml
db_url: postgres://localhost/heroes
log_level: info
```

Watch changes: `viper.WatchConfig()`

## Alternatives

- `koanf` lighter
- Flags/env only for CLIs

## Usage in Makefile

`make run CONFIG=prod.yaml`

## Next

[07-database-access.md &rarr;](../07-database-access.md)

[&larr; Back to [TOC](../README.md#table-of-contents)]
