// Health resource — reference implementation for [P|GF].
//
// Contract:
//   - RegisterHealth(s) adds the resource definition to the MCP server.
//   - The handler receives the URI and returns resource contents + error.
package resources

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterHealth adds the health resource to s.
// Phase 3 [P|GF]: implement this function.
func RegisterHealth(s *server.MCPServer) {
	resource := mcp.NewResource(
		"mcp://health",
		"Health Status",
		mcp.WithResourceDescription("Server health status"),
		mcp.WithMIMEType("application/json"),
	)
	s.AddResource(resource, healthHandler)
}

func healthHandler(_ context.Context, _ mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	// Phase 3 [P|GF]: return mcp.TextResourceContents with JSON status payload
	panic("not implemented")
}
