package main

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"

	"github.com/adroit-group/gote/internal"
	"github.com/adroit-group/gote/internal/httpserver"
	"github.com/adroit-group/gote/pkg/config"
	"github.com/adroit-group/gote/pkg/infra"
	"github.com/adroit-group/gote/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

func main() {
	logger.SetupSlog("template", os.Stdout)
	config.AutoConfigure(internal.Configuration, viper.GetViper())

	validate := validator.New()
	ctx := context.Background()
	h := httpserver.NewServerHandler(validate)

	h.RegisterRoutes(viper.GetString(internal.ConfigHTTPBasePath))

	srv := &http.Server{
		Addr:              net.JoinHostPort("", viper.GetString(internal.ConfigHTTPPort)),
		Handler:           h,
		ReadTimeout:       viper.GetDuration(internal.ConfigHTTPReadTimeout),
		ReadHeaderTimeout: viper.GetDuration(internal.ConfigHTTPReadHeaderTimeout),
		WriteTimeout:      viper.GetDuration(internal.ConfigHTTPWriteTimeout),
		IdleTimeout:       viper.GetDuration(internal.ConfigHTTPIdleTimeout),
	}

	err := infra.RunHTTPServerWithGracefulShutdown(ctx, srv)
	if err != nil {
		slog.Error("failed handle server", "error", err)
		os.Exit(1)
	}
}
