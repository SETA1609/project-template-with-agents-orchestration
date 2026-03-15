//! Health resource — reference implementation for [P|O].
//!
//! Contract (must match the registration pattern in server.rs):
//!   - `definition()` returns the Resource metadata registered at startup.
//!   - `read()` is dispatched by URI from ServerHandler::read_resource.

use rmcp::model::{ReadResourceResult, Resource, ResourceContents};

/// Resource metadata registered in McpServer::new().
pub fn definition() -> Resource {
    Resource {
        uri: "mcp://health".into(),
        name: Some("Health Status".into()),
        description: Some("Server health status".into()),
        mime_type: Some("application/json".into()),
    }
}

/// Returns current server health as a JSON resource.
pub async fn read(_uri: &str) -> Result<ReadResourceResult, rmcp::Error> {
    Ok(ReadResourceResult {
        contents: vec![ResourceContents::text(r#"{"status": "ok"}"#, "mcp://health")],
    })
}
