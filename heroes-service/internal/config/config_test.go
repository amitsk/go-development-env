package config

import (
	"testing"
	"time"
)

func TestFromEnvDefaults(t *testing.T) {
	t.Setenv("PORT", "")
	t.Setenv("LOG_LEVEL", "")
	t.Setenv("PPROF", "")
	t.Setenv("SHUTDOWN_TIMEOUT", "")

	cfg, err := FromEnv()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Addr != ":8080" || cfg.LogLevel != "info" || cfg.Pprof {
		t.Fatalf("defaults: %+v", cfg)
	}
	if cfg.ShutdownTimeout != 10*time.Second {
		t.Fatalf("timeout: %s", cfg.ShutdownTimeout)
	}
}

func TestFromEnvOverrides(t *testing.T) {
	t.Setenv("PORT", "9090")
	t.Setenv("LOG_LEVEL", "DEBUG")
	t.Setenv("PPROF", "1")
	t.Setenv("SHUTDOWN_TIMEOUT", "3s")

	cfg, err := FromEnv()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Addr != ":9090" || cfg.LogLevel != "debug" || !cfg.Pprof {
		t.Fatalf("overrides: %+v", cfg)
	}
	if cfg.ShutdownTimeout != 3*time.Second {
		t.Fatalf("timeout: %s", cfg.ShutdownTimeout)
	}
}

func TestFromEnvRejectsBadPort(t *testing.T) {
	t.Setenv("PORT", "abc")
	if _, err := FromEnv(); err == nil {
		t.Fatal("expected error")
	}
}
