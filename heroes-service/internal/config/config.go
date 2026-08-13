package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Config is process settings loaded from the environment.
type Config struct {
	Addr            string
	LogLevel        string
	Pprof           bool
	ShutdownTimeout time.Duration
}

// FromEnv reads PORT, LOG_LEVEL, PPROF, and SHUTDOWN_TIMEOUT.
func FromEnv() (Config, error) {
	cfg := Config{
		Addr:            ":8080",
		LogLevel:        "info",
		ShutdownTimeout: 10 * time.Second,
	}

	if port := os.Getenv("PORT"); port != "" {
		if _, err := strconv.Atoi(port); err != nil {
			return Config{}, fmt.Errorf("PORT %q is not a number", port)
		}
		cfg.Addr = ":" + port
	}

	if level := strings.ToLower(strings.TrimSpace(os.Getenv("LOG_LEVEL"))); level != "" {
		switch level {
		case "debug", "info", "warn", "error":
			cfg.LogLevel = level
		default:
			return Config{}, fmt.Errorf("LOG_LEVEL %q is not debug|info|warn|error", level)
		}
	}

	if v := os.Getenv("PPROF"); v == "1" || strings.EqualFold(v, "true") {
		cfg.Pprof = true
	}

	if raw := os.Getenv("SHUTDOWN_TIMEOUT"); raw != "" {
		d, err := time.ParseDuration(raw)
		if err != nil {
			return Config{}, fmt.Errorf("SHUTDOWN_TIMEOUT: %w", err)
		}
		cfg.ShutdownTimeout = d
	}

	return cfg, nil
}
