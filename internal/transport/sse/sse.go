// Package sse implements the HTTP + SSE transport backend using mcp-go v0.45.
//
// Confirmed API shape (verified against source):
//
//	server.NewSSEServer(s, server.WithBaseURL(baseURL))  — options pattern, NOT positional
//	(*SSEServer).Start(addr)                             — blocks on ListenAndServe
//	(*SSEServer).Shutdown(ctx)                           — graceful shutdown, closes all sessions
//
// Default endpoints registered by mcp-go:
//
//	GET  /sse      — SSE stream (server → client)
//	POST /message  — JSON-RPC messages (client → server, responds via SSE)
package sse

import (
	"context"
	"net/http"
	"time"

	"github.com/mark3labs/mcp-go/server"
)

// SSETransport serves MCP over HTTP + Server-Sent Events.
type SSETransport struct {
	Addr    string // e.g. "127.0.0.1:3000"
	BaseURL string // e.g. "http://localhost:3000"
}

// Serve starts the SSE HTTP server and blocks until ctx is cancelled or a
// fatal error occurs. On ctx cancellation a 5-second graceful shutdown is
// attempted before returning.
func (t SSETransport) Serve(ctx context.Context, s *server.MCPServer) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	sseServer := server.NewSSEServer(
		s,
		server.WithBaseURL(t.BaseURL),
		server.WithHTTPServer(&http.Server{Addr: t.Addr, Handler: mux}),
	)
	mux.Handle("/", sseServer)

	errCh := make(chan error, 1)
	go func() {
		errCh <- sseServer.Start(t.Addr)
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return sseServer.Shutdown(shutdownCtx)
	}
}
