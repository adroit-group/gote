package infra

import (
	"context"
	"io"
	"net"
	"testing"
	"time"

	"github.com/adroit-group/gote/pkg/logger"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestRunGRPCServerWithGracefulShutdown(t *testing.T) {
	logger.SetupSlog("test", io.Discard)

	srv := grpc.NewServer()
	listener, err := net.Listen("tcp", "localhost:0")
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())

	// Run server in goroutine
	done := make(chan error, 1)
	go func() {
		done <- RunGRPCServerWithGracefulShutdown(ctx, srv, listener)
	}()

	// Give server time to start
	time.Sleep(50 * time.Millisecond)

	// Trigger shutdown
	cancel()

	// Should complete without hanging
	select {
	case err := <-done:
		// Should return nil now that implementation handles close errors properly
		if err != nil {
			t.Errorf("Expected nil, got: %v", err)
		}
	case <-time.After(5 * time.Second):
		t.Error("Server did not shut down within expected time")
	}
}

func TestRunGRPCServerWithGracefulShutdown_ListenerError(t *testing.T) {
	logger.SetupSlog("test", io.Discard)

	srv := grpc.NewServer()

	// Create a listener and close it immediately to simulate error
	listener, err := net.Listen("tcp", "localhost:0")
	require.NoError(t, err)
	err = listener.Close()
	require.NoError(t, err)

	ctx := context.Background()

	// Run the server - should return error due to closed listener
	err = RunGRPCServerWithGracefulShutdown(ctx, srv, listener)
	if err == nil {
		t.Error("Expected error when using closed listener, but got nil")
	}
}

