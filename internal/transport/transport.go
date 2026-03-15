// Package transport defines the interface every transport backend must satisfy.
//
// # Design contract ([C] — for [O|G] to implement)
//
// Each transport backend must implement [Transport]. main.go resolves the
// configured transport and calls Serve — it never imports a concrete backend
// directly, only this interface.
//
// Adding a new transport:
//  1. Create a sub-package (e.g. internal/transport/streamable).
//  2. Implement Transport.
//  3. Add a case to the switch in cmd/server/main.go.
//  4. No other files need to change.
package transport

import (
	"context"

	"github.com/mark3labs/mcp-go/server"
)

// Transport is the contract every backend must fulfil.
// It receives a fully-configured MCP server and runs until ctx is cancelled
// or a fatal error occurs.
type Transport interface {
	Serve(ctx context.Context, s *server.MCPServer) error
}
