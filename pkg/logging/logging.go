package logging

import (
	"log/slog"
	"os"
)

// Default logger
func InitSlog(env string) *slog.Logger {
	switch env {
	case "prod":
		return slog.New(slog.NewTextHandler(os.Stdout, nil))
	case "dev":
		return slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}
	return nil
}
