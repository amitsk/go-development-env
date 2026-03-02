# 07. Configuration Management

In this chapter, we'll learn how to handle configuration for our application using **Viper**.

[&larr; Back to [TOC](../README.md#table-of-contents)] | [&larr; [06-logging.md](./06-logging.md)] | [&rarr; [08-database-access.md](./08-database-access.md)]

## Why Use External Configuration?

When you're building a professional application, you'll need it to run in different environments: your laptop (Development), a testing server (Staging), and the real world (Production).

The **12-Factor App** methodology states that you should never hardcode configuration values (like database passwords or API keys) in your code. Instead, you should keep them separate so you can:

1. **Keep Secrets Safe**: Don't commit passwords to your GitHub repository!
2. **Change Behavior Without Rebuilding**: Change the log level or database URL without having to re-compile your code.
3. **Be Environment-Agnostic**: Use a different database on your laptop than in production.

---

## Introducing Viper

[**Viper**](https://github.com/spf13/viper) is the most popular configuration library for Go. It can read configuration from:
- Files (JSON, TOML, YAML, HCL, envfile)
- Environment variables
- Command-line flags
- Remote configuration systems (like Consul or Etcd)

---

## Example: Using Viper

Let's see how we can use Viper to load configuration into a Go `struct`.

### 1. Define Your Configuration Struct

We use \"tags\" (the text in backticks) to tell Viper how to map the configuration file to our struct.

```go
type Config struct {
    Port     int    `mapstructure:\"port\"`
    Database string `mapstructure:\"database_url\"`
    LogLevel string `mapstructure:\"log_level\"`
}
```

### 2. Load the Configuration

```go
package main

import (
    \"fmt\"
    \"github.com/spf13/viper\"
)

func LoadConfig() (*Config, error) {
    // Tell Viper where to look for the file
    viper.SetConfigName(\"config\") // Name of the file (without extension)
    viper.SetConfigType(\"yaml\")   // Extension of the file
    viper.AddConfigPath(\".\")      // Look in the current directory

    // Automatically read from environment variables too!
    // For example, DATABASE_URL will override the value in the file.
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        return nil, fmt.Errorf(\"failed to read config: %w\", err)
    }

    var cfg Config
    if err := viper.Unmarshal(&cfg); err != nil {
        return nil, fmt.Errorf(\"failed to unmarshal config: %w\", err)
    }

    return &cfg, nil
}
```

### 3. The Configuration File (`config.yaml`)

```yaml
port: 8080
database_url: \"postgres://localhost/heroes\"
log_level: \"info\"
```

---

## Best Practices

- **Use Environment Variables for Secrets**: Keep your database password in an environment variable (`DB_PASSWORD`), and only put non-sensitive defaults in your `config.yaml`.
- **Set Defaults**: Use `viper.SetDefault(\"port\", 8080)` to ensure your app always has a working configuration even if the file is missing.
- **Fail Fast**: If a critical configuration value is missing (like a database URL), your application should stop and print a clear error message.

## Next Step

Now that we can configure our app, let's learn how to store and retrieve data from a database.

[08-database-access.md &rarr;](./08-database-access.md)

[&larr; Back to [TOC](../README.md#table-of-contents)]
