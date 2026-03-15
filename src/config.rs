// Phase 4 — implemented by [O|G]
// Loads MCP_TRANSPORT, HOST, PORT, LOG_LEVEL from ENV; --transport from CLI.

use crate::transport::TransportKind;

/// Runtime configuration resolved at startup.
/// ENV vars are loaded first; CLI flags win if both are set.
#[derive(Debug)]
pub struct Config {
    pub transport: TransportKind,
    pub host: String,
    pub port: u16,
    pub log_level: String,
}

impl Config {
    pub fn load() -> Self {
        todo!("Phase 4: implement config loading from ENV + CLI")
    }
}
