//! Echo tool — reference implementation for [P|O].
//!
//! Contract (must match the registration pattern in server.rs):
//!   - `definition()` returns the Tool metadata registered at startup.
//!   - `call(args)` is dispatched by name from ServerHandler::call_tool.

use rmcp::model::{CallToolResult, Tool, ToolContent};
use serde::{Deserialize, Serialize};
use schemars::JsonSchema;
use serde_json::Value;

#[derive(Debug, Serialize, Deserialize, JsonSchema)]
pub struct EchoArgs {
    pub input: String,
}

/// Tool metadata registered in McpServer::new().
pub fn definition() -> Tool {
    Tool {
        name: "echo".into(),
        description: Some("Echoes the input string back".into()),
        input_schema: schemars::schema_for!(EchoArgs).into(),
    }
}

/// Echoes the `input` field from args back as text content.
pub async fn call(args: Value) -> Result<CallToolResult, rmcp::Error> {
    let echo_args: EchoArgs = serde_json::from_value(args)?;
    Ok(CallToolResult {
        content: vec![ToolContent::text(echo_args.input)],
        is_error: None,
    })
}
