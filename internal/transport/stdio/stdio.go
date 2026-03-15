// Package stdio implements the stdio transport backend.
// Implemented by [O|G] in Phase 2a.
package stdio

import (
	"context"

	"github.com/mark3labs/mcp-go/server"
)

// StdioTransport serves MCP over stdin/stdout.
type StdioTransport struct{}

// Serve runs the MCP server over stdin/stdout until ctx is cancelled.
func (t StdioTransport) Serve(ctx context.Context, s *server.MCPServer) error {
	// Phase 2a [O|G]: call server.ServeStdio(s)
	panic("not implemented")
}
