package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/adroit-group/go-template/internal/httpserver"
	"github.com/adroit-group/go-template/pkg/infra"
	"github.com/adroit-group/go-template/pkg/logger"
	"github.com/go-playground/validator/v10"
)

func main() {
	logger.SetupSlog("template", os.Stdout)

	valdate := validator.New()
	ctx := context.Background()
	h := httpserver.NewServerHandler(valdate)

	h.RegisterRoutes("/api")

	srv := &http.Server{
		Addr:              ":80",
		Handler:           h,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
	}

	err := infra.RunHTTPServerWithGracefulShutdown(ctx, srv)
	if err != nil {
		slog.Error("failed handle server", "error", err)
		os.Exit(1)
	}
}
