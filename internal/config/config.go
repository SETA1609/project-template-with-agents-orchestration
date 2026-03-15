// Package config loads runtime configuration from environment variables and
// CLI flags.
package config

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

// Config holds all runtime settings resolved at startup.
// ENV vars are read first; CLI flags override them if provided.
type Config struct {
	Transport string // "stdio" | "http"  (MCP_TRANSPORT / --transport)
	Host      string // bind host         (HOST / --addr)
	Port      string // bind port         (PORT / --addr)
	LogLevel  string // slog level        (LOG_LEVEL)
}

// BaseURL returns the HTTP base URL derived from Host and Port.
func (c Config) BaseURL() string {
	return "http://" + c.Host + ":" + c.Port
}

// Addr returns the host:port string for net.Listen.
func (c Config) Addr() string {
	return c.Host + ":" + c.Port
}

// Load reads Config from ENV then applies CLI overrides.
// Supported flags:
//   - --transport stdio|http
//   - --addr host:port
func Load() Config {
	cfg := Config{
		Transport: envOrDefault("MCP_TRANSPORT", "http"),
		Host:      envOrDefault("HOST", "127.0.0.1"),
		Port:      envOrDefault("PORT", "3000"),
		LogLevel:  envOrDefault("LOG_LEVEL", "info"),
	}

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	transportFlag := fs.String("transport", "", "transport mode: stdio or http")
	addrFlag := fs.String("addr", "", "bind address (host:port)")
	_ = fs.Parse(os.Args[1:])

	if *transportFlag != "" {
		cfg.Transport = *transportFlag
	}

	cfg.Transport = strings.ToLower(strings.TrimSpace(cfg.Transport))
	if cfg.Transport != "stdio" && cfg.Transport != "http" {
		panic(fmt.Sprintf("invalid transport %q: must be 'stdio' or 'http'", cfg.Transport))
	}

	if *addrFlag != "" {
		host, port, err := splitHostPort(*addrFlag)
		if err != nil {
			panic(fmt.Sprintf("invalid --addr value %q: %v", *addrFlag, err))
		}
		cfg.Host = host
		cfg.Port = port
	}

	return cfg
}

func envOrDefault(key, fallback string) string {
	v := strings.TrimSpace(os.Getenv(key))
	if v == "" {
		return fallback
	}
	return v
}

func splitHostPort(addr string) (string, string, error) {
	host, port, err := net.SplitHostPort(strings.TrimSpace(addr))
	if err != nil {
		return "", "", err
	}
	if host == "" {
		host = "0.0.0.0"
	}
	if port == "" {
		return "", "", fmt.Errorf("missing port")
	}
	return host, port, nil
}
