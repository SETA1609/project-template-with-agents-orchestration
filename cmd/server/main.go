// Entry-point: resolves config, builds the MCP server, dispatches to transport.
// Dispatch body implemented by [O|P] in Phase 2c.
package main

import (
	"context"
	"log"

	"github.com/SETA1609/my-mcp-agent-orchestrator/internal/config"
	internalserver "github.com/SETA1609/my-mcp-agent-orchestrator/internal/server"
	"github.com/SETA1609/my-mcp-agent-orchestrator/internal/transport"
	"github.com/SETA1609/my-mcp-agent-orchestrator/internal/transport/sse"
	"github.com/SETA1609/my-mcp-agent-orchestrator/internal/transport/stdio"
)

func main() {
	cfg := config.Load()
	s := internalserver.New()

	var t transport.Transport
	switch cfg.Transport {
	case "stdio":
		t = stdio.StdioTransport{}
	case "http":
		t = sse.SSETransport{
			Addr:    cfg.Addr(),
			BaseURL: cfg.BaseURL(),
		}
	default:
		log.Fatalf("unknown transport: %q (set MCP_TRANSPORT to 'stdio' or 'http')", cfg.Transport)
	}

	if err := t.Serve(context.Background(), s); err != nil {
		log.Fatal(err)
	}
}
