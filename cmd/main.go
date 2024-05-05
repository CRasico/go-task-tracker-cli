package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/crasico/go-task-tracker-cli/pkg/ui"
)

func main() {
	var isDebug bool
	flag.BoolVar(&isDebug, "debug", false, "Is Debug Mode")

	flag.Parse()

	logger := setupLogger(isDebug)
	logger.Debug("Starting Up Go Task Tracker CLI")

	window := ui.NewWindow(ui.WindowArgs{Width: 10, Height: 10, HasBorder: true})

	window.Draw(3, 3, '+')
	window.Render()
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
