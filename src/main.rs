// Phase 2c — dispatch implemented by [O|P]

mod config;
mod resources;
mod server;
mod tools;
mod transport;

use config::Config;
use server::McpServer;
use transport::{Transport, TransportKind, http_sse::HttpSseTransport, stdio::StdioTransport};
use std::sync::Arc;

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    let config = Config::load();
    let server = Arc::new(McpServer::new());

    match config.transport {
        TransportKind::Stdio => StdioTransport.serve(server.clone()).await,
        TransportKind::Http => HttpSseTransport {
            host: config.host,
            port: config.port,
        }
        .serve(server.clone())
        .await,
    }
}
