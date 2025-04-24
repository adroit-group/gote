package logger

import (
	"io"
	"log/slog"
)

func SetupSlog(service string, output io.Writer) *slog.Logger {
	l := slog.New(slog.NewJSONHandler(output, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})).With("service", service)

	slog.SetDefault(l)

	return l
}
