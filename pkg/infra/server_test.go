package infra

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/adroit-group/go-template/pkg/logger"
	"github.com/stretchr/testify/require"
)

func TestRunHTTPServerWithGracefulShutdown(t *testing.T) {
	logger.SetupSlog("test", io.Discard)

	// Create a test HTTP server
	var serverCalled bool
	var mu sync.Mutex
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		serverCalled = true
		mu.Unlock()
		w.WriteHeader(http.StatusOK)
	})

	// Find an available port
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("Failed to find available port: %v", err)
	}
	port := listener.Addr().(*net.TCPAddr).Port
	err = listener.Close()
	require.NoError(t, err)

	addr := fmt.Sprintf("localhost:%d", port)
	srv := &http.Server{
		Addr:    addr,
		Handler: testHandler,
	}

	// Create a cancellable context
	ctx, cancel := context.WithCancel(context.Background())

	// Run the server in a goroutine
	serverErrCh := make(chan error, 1)
	go func() {
		serverErrCh <- RunHTTPServerWithGracefulShutdown(ctx, srv)
	}()

	// Give the server time to start
	time.Sleep(100 * time.Millisecond)

	// Test that server is running by making a request
	client := &http.Client{Timeout: 1 * time.Second}
	resp, err := client.Get("http://" + addr)
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	err = resp.Body.Close()
	require.NoError(t, err)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}

	mu.Lock()
	if !serverCalled {
		t.Error("Server handler was not called")
	}
	mu.Unlock()

	// Test graceful shutdown
	cancel() // Trigger shutdown

	// Wait for server to stop with timeout
	select {
	case err := <-serverErrCh:
		if err != nil {
			t.Errorf("Server returned error on shutdown: %v", err)
		}
	case <-time.After(10 * time.Second): // Should be longer than the 5s shutdown timeout
		t.Error("Server did not shut down within expected time")
	}
}
