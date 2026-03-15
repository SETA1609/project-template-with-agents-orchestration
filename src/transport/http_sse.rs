// Phase 2b — designed by [K|C], implemented by [O|G]
// Axum router with POST /message and GET /sse endpoints.

use super::Transport;
use crate::server::McpServer;
use anyhow::Result;
use async_trait::async_trait;
use rmcp::transport::sse::ServerSseTransport;
use std::sync::Arc;
use tokio::signal;

pub struct HttpSseTransport {
    pub host: String,
    pub port: u16,
}

#[async_trait]
impl Transport for HttpSseTransport {
    async fn serve(self, server: Arc<McpServer>) -> Result<()> {
        let transport = ServerSseTransport::new(self.host, self.port);
        let serve_future = server.serve(transport);
        tokio::select! {
            result = serve_future => result,
            _ = signal::ctrl_c() => Ok(()),
        }
    }
}
