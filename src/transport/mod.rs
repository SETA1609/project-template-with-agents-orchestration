use anyhow::Result;
use async_trait::async_trait;
use std::sync::Arc;

use crate::server::McpServer;

/// Selects which transport backend to start at runtime.
/// Resolved from `MCP_TRANSPORT` env var or `--transport` CLI flag.
#[derive(Debug, Clone, Copy, PartialEq, Eq, Default)]
pub enum TransportKind {
    Stdio,
    #[default]
    Http,
}

/// Every transport backend implements this trait.
/// It consumes itself and runs until the server shuts down or an error occurs.
#[async_trait]
pub trait Transport: Send + 'static {
    async fn serve(self, server: Arc<McpServer>) -> Result<()>;
}

pub mod http_sse;
pub mod stdio;
