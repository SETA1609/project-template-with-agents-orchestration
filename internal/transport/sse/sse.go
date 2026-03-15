// Package sse implements the HTTP + SSE transport backend.
// Designed by [K|C] in Phase 2b, implemented by [O|G].
package sse

import (
	"context"

	"github.com/mark3labs/mcp-go/server"
)

// SSETransport serves MCP over HTTP + Server-Sent Events.
type SSETransport struct {
	Addr    string // e.g. "127.0.0.1:3000"
	BaseURL string // e.g. "http://localhost:3000"
}

// Serve starts the SSE HTTP server and blocks until ctx is cancelled.
func (t SSETransport) Serve(ctx context.Context, s *server.MCPServer) error {
	// Phase 2b [O|G]: create server.NewSSEServer(s, t.BaseURL), listen on t.Addr,
	// and shut down cleanly when ctx is cancelled.
	panic("not implemented")
}
