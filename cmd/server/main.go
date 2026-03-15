// Entry-point: resolves config, builds the MCP server, dispatches to transport.
// Dispatch body implemented by [O|P] in Phase 2c.
package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/SETA1609/my-mcp-agent-orchestrator/internal/config"
	internalserver "github.com/SETA1609/my-mcp-agent-orchestrator/internal/server"
	"github.com/SETA1609/my-mcp-agent-orchestrator/internal/transport"
	"github.com/SETA1609/my-mcp-agent-orchestrator/internal/transport/sse"
	"github.com/SETA1609/my-mcp-agent-orchestrator/internal/transport/stdio"
)

func main() {
	cfg := config.Load()
	configureLogging(cfg.LogLevel)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

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
		slog.Error("invalid transport", "transport", cfg.Transport)
		os.Exit(1)
	}

	if err := t.Serve(ctx, s); err != nil {
		slog.Error("server exited with error", "error", err)
		os.Exit(1)
	}
}

func configureLogging(level string) {
	logLevel := parseLogLevel(level)
	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel})
	slog.SetDefault(slog.New(h))
}

func parseLogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
