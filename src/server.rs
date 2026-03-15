//! MCP server core — capability registration and handler dispatch.
//!
//! # Registration pattern ([C] design contract for [O|G] to implement)
//!
//! Each tool module must export:
//!   - `pub fn definition() -> rmcp::model::Tool`  — name, description, input schema
//!   - `pub async fn call(args: serde_json::Value) -> rmcp::model::CallToolResult`
//!
//! Each resource module must export:
//!   - `pub fn definition() -> rmcp::model::Resource`  — uri, name, description, mime_type
//!   - `pub async fn read(uri: &str) -> rmcp::model::ReadResourceResult`
//!
//! To register a new capability:
//!   1. Create the module under `src/tools/` or `src/resources/`.
//!   2. Add `pub mod <name>;` to the respective `mod.rs`.
//!   3. Add `tools::name::definition()` / `resources::name::definition()` to the vecs in `McpServer::new()`.
//!
//! No changes to `ServerHandler` are ever needed — the impl delegates to the vecs.

use crate::{resources, tools};
use rmcp::{
    ServerHandler,
    model::{
        CallToolResult, ListResourcesResult, ListToolsResult, ReadResourceResult, Resource,
        ServerCapabilities, ServerInfo, Tool, Prompt, ListPromptsResult, GetPromptResult,
    },
};

/// Holds all registered capabilities built once at startup.
pub struct McpServer {
    tools: Vec<Tool>,
    resources: Vec<Resource>,
    prompts: Vec<Prompt>,
}

impl McpServer {
    pub fn new() -> Self {
        Self {
            tools: vec![
                tools::echo::definition(),
            ],
            resources: vec![
                resources::health::definition(),
            ],
            prompts: vec![], // No prompts implemented yet
        }
    }
}

impl Default for McpServer {
    fn default() -> Self {
        Self::new()
    }
}

// --- ServerHandler impl ([O|G] fills in the method bodies) ---

impl ServerHandler for McpServer {
    fn get_info(&self) -> ServerInfo {
        ServerInfo {
            capabilities: ServerCapabilities::builder()
                .enable_tools()
                .enable_resources()
                .enable_prompts()
                .build(),
            ..Default::default()
        }
    }

    async fn list_tools(
        &self,
        _request: rmcp::model::PaginatedRequestParam,
        _context: rmcp::service::RequestContext<rmcp::RoleServer>,
    ) -> Result<ListToolsResult, rmcp::Error> {
        Ok(ListToolsResult {
            tools: self.tools.clone(),
            next_cursor: None,
            meta: None,
        })
    }

    async fn call_tool(
        &self,
        request: rmcp::model::CallToolRequestParam,
        _context: rmcp::service::RequestContext<rmcp::RoleServer>,
    ) -> Result<CallToolResult, rmcp::Error> {
        // Dispatch by tool name to the matching module's call() function.
        match request.name.as_str() {
            "echo" => tools::echo::call(request.arguments).await,
            name => Err(rmcp::Error::invalid_params(
                format!("unknown tool: {name}"),
                None,
            )),
        }
    }

    async fn list_resources(
        &self,
        _request: rmcp::model::PaginatedRequestParam,
        _context: rmcp::service::RequestContext<rmcp::RoleServer>,
    ) -> Result<ListResourcesResult, rmcp::Error> {
        Ok(ListResourcesResult {
            resources: self.resources.clone(),
            next_cursor: None,
            meta: None,
        })
    }

    async fn read_resource(
        &self,
        request: rmcp::model::ReadResourceRequestParam,
        _context: rmcp::service::RequestContext<rmcp::RoleServer>,
    ) -> Result<ReadResourceResult, rmcp::Error> {
        // Dispatch by URI to the matching module's read() function.
        match request.uri.as_str() {
            "mcp://health" => resources::health::read(request.uri.as_str()).await,
            uri => Err(rmcp::Error::invalid_params(
                format!("unknown resource: {uri}"),
                None,
            )),
        }
    }

    async fn list_prompts(
        &self,
        _request: rmcp::model::PaginatedRequestParam,
        _context: rmcp::service::RequestContext<rmcp::RoleServer>,
    ) -> Result<ListPromptsResult, rmcp::Error> {
        Ok(ListPromptsResult {
            prompts: self.prompts.clone(),
            next_cursor: None,
            meta: None,
        })
    }

    async fn get_prompt(
        &self,
        request: rmcp::model::GetPromptRequestParam,
        _context: rmcp::service::RequestContext<rmcp::RoleServer>,
    ) -> Result<GetPromptResult, rmcp::Error> {
        Err(rmcp::Error::invalid_params(
            format!("unknown prompt: {}", request.name),
            None,
        ))
    }
}
