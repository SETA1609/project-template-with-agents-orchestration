// Package config loads runtime configuration from environment variables and
// CLI flags. Implemented by [O|G] in Phase 4.
package config

// Config holds all runtime settings resolved at startup.
// ENV vars are read first; CLI flags override them if provided.
type Config struct {
	Transport string // "stdio" | "http"  (MCP_TRANSPORT / --transport)
	Host      string // bind host         (HOST / --host)
	Port      string // bind port         (PORT / --port)
	LogLevel  string // slog level        (LOG_LEVEL / --log-level)
}

// BaseURL returns the HTTP base URL derived from Host and Port.
func (c Config) BaseURL() string {
	return "http://" + c.Host + ":" + c.Port
}

// Addr returns the host:port string for net.Listen.
func (c Config) Addr() string {
	return c.Host + ":" + c.Port
}

// Load reads Config from ENV then applies any CLI flag overrides.
// Phase 4 [O|G]: implement this function.
func Load() Config {
	panic("not implemented")
}
