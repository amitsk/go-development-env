package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/amitsk/go-development-env/heroes-service/internal/api"
	"github.com/amitsk/go-development-env/heroes-service/internal/config"
	"github.com/amitsk/go-development-env/heroes-service/internal/store"
)

func main() {
	os.Exit(run())
}

func run() int {
	cfg, err := config.FromEnv()
	if err != nil {
		slog.Error("config", slog.Any("err", err))
		return 1
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: parseLevel(cfg.LogLevel),
	}))
	slog.SetDefault(logger)

	srv := &http.Server{
		Addr:    cfg.Addr,
		Handler: api.NewServer(store.NewMemory(), logger, cfg.Pprof).Handler(),
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	errCh := make(chan error, 1)
	go func() {
		logger.Info("listening", slog.String("addr", cfg.Addr), slog.Bool("pprof", cfg.Pprof))
		errCh <- srv.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			logger.Error("shutdown", slog.Any("err", err))
			return 1
		}
		logger.Info("stopped")
		return 0
	case err := <-errCh:
		if err != nil && err != http.ErrServerClosed {
			logger.Error("listen", slog.Any("err", err))
			return 1
		}
		return 0
	}
}

func parseLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
