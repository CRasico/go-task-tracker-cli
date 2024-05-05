package main

import (
	"flag"
	"log/slog"
	"os"
	"os/exec"
	"time"

	"github.com/crasico/go-task-tracker-cli/pkg/ui"
)

func main() {
	var isDebug bool
	flag.BoolVar(&isDebug, "debug", false, "Is Debug Mode")

	flag.Parse()

	logger := setupLogger(isDebug)
	logger.Debug("Starting Up Go Task Tracker CLI")

	window := ui.NewWindow(ui.WindowArgs{Width: 20, Height: 10, HasBorder: true})

	window.Render()

	for i := range window.Pixels[0] {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		window.Draw(3, i, '+')
		window.Render()

		time.Sleep(300 * time.Millisecond)
	}
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
