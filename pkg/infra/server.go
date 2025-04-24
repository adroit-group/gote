package infra

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"golang.org/x/sync/errgroup"
)

func RunHTTPServerWithGracefulShutdown(ctx context.Context, srv *http.Server) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		<-ctx.Done()

		slog.Info("shutting down server...")
		ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
		defer cancel()

		return srv.Shutdown(ctx)
	})

	eg.Go(func() error {
		slog.Info("starting server", "addr", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}

		return nil
	})

	return eg.Wait()
}
