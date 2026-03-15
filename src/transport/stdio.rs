// Phase 2a — implemented by [O|G]
// Reads newline-delimited JSON-RPC from stdin, writes responses to stdout.

use super::Transport;
use crate::server::McpServer;
use anyhow::Result;
use async_trait::async_trait;
use std::sync::Arc;

pub struct StdioTransport;

#[async_trait]
impl Transport for StdioTransport {
    async fn serve(self, server: Arc<McpServer>) -> Result<()> {
        let transport = (tokio::io::stdin(), tokio::io::stdout());
        server.serve(transport).await
    }
}
