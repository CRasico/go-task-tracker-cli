package main

import (
	"flag"
	"log/slog"
	"os"
)

func main() {
	var isDebug bool
	flag.BoolVar(&isDebug, "debug", false, "Is Debug Mode")

	flag.Parse()

	logger := setupLogger(isDebug)
	logger.Debug("Starting Up Go Task Tracker CLI")
}

func setupLogger(isDebug bool) *slog.Logger {
	if isDebug {
		return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	}

	return slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelError,
	}))
}
