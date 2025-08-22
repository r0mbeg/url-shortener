package main

import (
	"os"
	"url-shortener/internal/config"
	"url-shortener/internal/lib/logger/sl"
	"url-shortener/internal/storage/sqlite"

	"golang.org/x/exp/slog"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	_, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		return
	}

	// TODO: init router: chi, "chi render"

	// TODO: run server: run server
}

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		opts := &slog.HandlerOptions{Level: slog.LevelDebug}
		log = slog.New(
			opts.NewTextHandler(os.Stdout),
		)
	case envDev:
		opts := &slog.HandlerOptions{Level: slog.LevelDebug}
		log = slog.New(
			opts.NewJSONHandler(os.Stdout),
		)
	case envProd:
		opts := &slog.HandlerOptions{Level: slog.LevelInfo}
		log = slog.New(
			opts.NewJSONHandler(os.Stdout),
		)
	}
	return log
}
