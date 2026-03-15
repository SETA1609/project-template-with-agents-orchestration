// Package server constructs and configures the MCP server instance.
//
// # Registration pattern ([C] design contract for [O|G])
//
// Each tool module must export:
//   - Register(s *mcp.Server)  — adds the tool to the server at startup
//
// Each resource module must export:
//   - Register(s *mcp.Server)  — adds the resource to the server at startup
//
// To add a new capability:
//  1. Create internal/tools/<name>.go or internal/resources/<name>.go.
//  2. Implement Register(s *mcp.Server).
//  3. Call it from internal/tools/tools.go or internal/resources/resources.go.
//  4. No changes to this file are needed.
package server

import (
	"github.com/mark3labs/mcp-go/server"
)

// New creates a configured MCPServer with all capabilities registered.
// Phase 3 [O|G]: call tools.RegisterAll and resources.RegisterAll here.
func New() *server.MCPServer {
	panic("not implemented")
}
