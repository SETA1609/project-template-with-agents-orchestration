// Package resources registers all MCP resources with the server.
package resources

import "github.com/mark3labs/mcp-go/server"

// RegisterAll adds every resource in this package to s.
// Phase 3 [O|G]: call each resource's Register function here.
func RegisterAll(s *server.MCPServer) {
	// health.Register(s)
	panic("not implemented")
}
