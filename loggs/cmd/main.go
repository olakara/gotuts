package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("Info log")
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	jsonLogger.Info("hello", "count", 3)
	jsonLogger.Error("Error")
	//slog.Debug("Debug")
}
