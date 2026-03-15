// Echo tool — reference implementation for [P|GF].
//
// Contract:
//   - Register(s) adds the tool definition to the MCP server.
//   - The handler func is passed directly to mcp-go; it receives a context
//     and the parsed request, and returns content + error.
package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Register adds the echo tool to s.
// Phase 3 [P|GF]: implement this function.
func RegisterEcho(s *server.MCPServer) {
	tool := mcp.NewTool("echo",
		mcp.WithDescription("Echoes the input string back"),
		mcp.WithString("input", mcp.Required(), mcp.Description("The string to echo")),
	)
	s.AddTool(tool, echoHandler)
}

func echoHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Phase 3 [P|GF]: extract req.Params.Arguments["input"], return mcp.NewToolResultText(...)
	panic("not implemented")
}
