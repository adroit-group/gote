package infra

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func RunGRPCServerWithGracefulShutdown(ctx context.Context, srv *grpc.Server, ls net.Listener) error {
	return RunGRPCServerWithHealthCheck(ctx, srv, ls, nil)
}

func RunGRPCServerWithHealthCheck(ctx context.Context, srv *grpc.Server, ls net.Listener, healthServer *health.Server) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		<-ctx.Done()

		srv.GracefulStop()

		if err := ls.Close(); err != nil && !isNetworkCloseError(err) {
			return err
		}

		return nil
	})

	eg.Go(func() error {
		if healthServer == nil {
			healthServer = health.NewServer()
		}
		grpc_health_v1.RegisterHealthServer(srv, healthServer)
		healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

		reflection.Register(srv)
		slog.Info("starting server", "addr", ls.Addr().String())
		if err := srv.Serve(ls); err != nil && err != grpc.ErrServerStopped {
			return err
		}

		return nil
	})

	return eg.Wait()
}

// isNetworkCloseError checks if an error is expected when closing network connections
func isNetworkCloseError(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, net.ErrClosed) {
		return true
	}

	// Check for *net.OpError with EBADF (bad file descriptor) which occurs when closing already closed connections
	var opErr *net.OpError
	if errors.As(err, &opErr) {
		if errors.Is(opErr.Err, syscall.EBADF) || errors.Is(opErr.Err, syscall.ECONNRESET) {
			return true
		}
	}

	return false
}
