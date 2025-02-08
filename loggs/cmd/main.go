package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("Info log")
	slog.Info("hello", "count", 3)
	//slog.Error("Error")
	//slog.Debug("Debug")
}
