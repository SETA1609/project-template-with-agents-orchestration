# PLAN — MCP Server (Rust · HTTP · Docker · stdio)

## Overview

Build a **Model Context Protocol (MCP) server** in Rust that supports **two transports**, selectable at runtime:

| Mode   | Transport                                          | Use-case                                                                  |
|--------|----------------------------------------------------|---------------------------------------------------------------------------|
| `stdio` | Standard input/output (JSON-RPC over stdin/stdout) | Local dev, subprocess launched by an AI client (e.g. Claude Desktop, VS Code) |
| `http`  | HTTP + Server-Sent Events                          | Docker deployment, remote access, multi-client                            |

The server will be packaged as a Docker image for the HTTP mode and can also be run as a plain binary for stdio mode.

---

## Goals

- Implement the MCP spec (2024-11-05 or later) over **HTTP + Server-Sent Events (SSE)** transport
- Write the server in **Rust** for performance, safety, and a small container footprint
- Ship a **multi-stage Docker image** (builder → distroless/alpine runtime)
- Keep the project structure clean so new tools/resources can be added with minimal boilerplate
- Let the user pick transport via a `MCP_TRANSPORT` env var (`stdio` | `http`) or a `--transport` CLI flag

---

## Stack

| Layer          | Choice                           | Rationale                                                                |
|----------------|----------------------------------|--------------------------------------------------------------------------|
| Language       | Rust (stable)                    | Performance, safety, great async ecosystem                               |
| Async runtime  | Tokio                            | De-facto standard; well-supported by web frameworks                      |
| HTTP framework | Axum                             | Ergonomic, tower-compatible, SSE support (http mode)                     |
| MCP SDK        | `rmcp` crate (official Rust SDK) | Supports both stdio **and** HTTP/SSE transports                          |
| Serialisation  | serde + serde_json               | Ubiquitous, zero-copy-friendly                                           |
| Container      | Docker (multi-stage)             | Reproducible builds for http mode; minimal runtime image                 |
| Config         | `MCP_TRANSPORT` env var or `--transport` CLI flag | Zero-friction mode switching                        |

---

## Project Structure

```
my-mcp-agent-orchestrator/
├── Cargo.toml
├── Cargo.lock
├── Dockerfile
├── docker-compose.yml
├── .dockerignore
├── .env.example
├── src/
│   ├── main.rs
│   ├── server.rs
│   ├── config.rs
│   ├── transport/
│   │   ├── mod.rs
│   │   ├── stdio.rs
│   │   └── http_sse.rs
│   ├── tools/
│   │   ├── mod.rs
│   │   └── echo.rs
│   └── resources/
│       ├── mod.rs
│       └── health.rs
├── tests/
│   ├── stdio_integration.rs
│   └── http_integration.rs
└── docs/
```

---

## How to read this plan

Each step lists the agents that may claim it, in priority order: `[Primary|Fallback]`.
- If the primary is available, they should claim it.
- If the primary is busy or unavailable, the fallback may claim it instead.
- To claim a step, replace `[ ]` with your symbol. To complete it, replace your symbol with `[x]`.
- Steps marked `[D]` require human input and cannot be claimed by any other agent.

---

## Implementation Phases

### Phase 1 — Scaffold & Tooling

- [x] `[C|K]` Design module boundaries and `Transport` trait — define the interface all other agents build against before any code is written
- [x] `[O|G]` Initialise Cargo binary crate (`cargo new --bin`) using the module structure defined above
- [x] `[P|O]` Add dependencies to `Cargo.toml`: `axum`, `tokio`, `rmcp`, `serde`, `serde_json`, `dotenvy`, `clap`, `tracing`, `tracing-subscriber`
- [x] `[P|O]` Configure `rustfmt.toml` and `clippy.toml` (see `docs/CODE_STYLE.md`)
- [x] `[P|O]` Add `.env.example` with `MCP_TRANSPORT`, `HOST`, `PORT`, `LOG_LEVEL`

### Phase 2 — Transport Layer

> `[C|K]` must complete the Phase 1 design step before any Phase 2 work begins.

#### 2a — stdio transport
- [x] `[O|G]` Implement `src/transport/stdio.rs`: read JSON-RPC from `stdin`, write responses to `stdout`, async loop

#### 2b — HTTP + SSE transport
- [x] `[K|C]` Design the Axum router shape and shared app state — delegate to `rmcp::transport::sse::ServerSseTransport`; no custom Axum router or channel wiring needed
- [x] `[O|G]` Implement `src/transport/http_sse.rs`: `POST /message` and `GET /sse` endpoints using the design above
- [x] `[O|G]` Implement graceful shutdown (SIGTERM via `tokio::signal`)

#### 2c — Transport dispatch
- [x] `[O|P]` Implement dispatch in `src/main.rs`:
  ```rust
  match config.transport {
      Transport::Stdio => run_stdio(server).await,
      Transport::Http  => run_http(server, config).await,
  }
  ```

### Phase 3 — MCP Server Core

- [x] `[K|C]` Design capability registration pattern — how tools, resources, and prompts are registered with `rmcp`
- [x] `[O|G]` Implement `src/server.rs`: initialise `rmcp` server, capability negotiation (`initialize` / `initialized`)
- [x] `[O|G]` Register built-in capabilities: `list_tools` / `call_tool`, `list_resources` / `read_resource`, `list_prompts` / `get_prompt`
- [x] `[P|O]` Add sample **echo tool** (`src/tools/echo.rs`) as a reference implementation
- [x] `[P|O]` Add sample **health resource** (`src/resources/health.rs`) as a reference implementation

### Phase 4 — Configuration & Observability

- [ ] `[O|G]` Implement `src/config.rs`: load `MCP_TRANSPORT`, `HOST`, `PORT`, `LOG_LEVEL` from ENV; parse `--transport` CLI flag via `clap`
- [ ] `[P|O]` Wire `tracing` + `tracing-subscriber` for structured JSON logging at startup
- [ ] `[P|O]` Add `GET /health` liveness endpoint (outside MCP protocol)

### Phase 5 — Docker Packaging

- [ ] `[C|K]` Review Dockerfile strategy — confirm base images, layer order, and secrets handling before `[O]` writes it
- [ ] `[O|G]` Write multi-stage `Dockerfile`: builder stage (`rust:slim` → `cargo build --release`) + runtime stage (`distroless/cc`)
- [ ] `[P|O]` Add `.dockerignore` (exclude `target/`, `.git`, secrets)
- [ ] `[P|O]` Add `docker-compose.yml` for local development convenience
- [ ] `[D]` Validate image size target: **< 50 MB**

### Phase 6 — Testing & CI

- [ ] `[K|C]` Design integration test harness — how to spin up the server in-process for both transports
- [ ] `[O|G]` Write unit tests for each tool and resource handler
- [ ] `[O|K]` Write integration tests: `tests/stdio_integration.rs` and `tests/http_integration.rs`
- [ ] `[O|G]` Write GitHub Actions workflow: `cargo fmt --check`, `cargo clippy -- -D warnings`, `cargo test`, `docker build` smoke-test
- [ ] `[D]` Review and approve CI workflow before merging to `main`

---

## Open Questions — `[D]` to resolve before Phase 2

- [ ] `[D]` Authentication strategy (API key header? OAuth2 bearer token?)
- [ ] `[D]` Rate-limiting requirements
- [ ] `[D]` Persistence layer needed? (e.g., tool outputs cached to Redis/SQLite)
- [ ] `[D]` Target deployment environment (bare Docker, Compose, Kubernetes, fly.io…)

---

## API Surface (MVP)

### HTTP mode

| Method | Endpoint   | Description                               |
|--------|------------|-------------------------------------------|
| GET    | `/sse`     | Open SSE stream (MCP server → client)     |
| POST   | `/message` | Send JSON-RPC message to MCP server       |
| GET    | `/health`  | Liveness probe (outside MCP)              |

### stdio mode

No HTTP endpoints. Reads newline-delimited JSON-RPC from `stdin`, writes responses to `stdout`.

---

## Key Design Decisions

1. **Dual transport, single binary** — same compiled binary serves both modes
2. **Runtime selection** — `MCP_TRANSPORT` env var or `--transport` CLI flag; CLI wins if both are set
3. **SSE over WebSockets** — MCP HTTP spec uses SSE for server-push; simpler to proxy, firewall-friendly
4. **`rmcp` crate** — avoids reimplementing JSON-RPC framing, session state, and capability negotiation
5. **Distroless runtime image** — reduces attack surface; no shell or package manager in production
6. **ENV-only config for Docker** — secrets injected via Docker secrets or orchestrator env vars

---

## References

- [MCP Specification](https://spec.modelcontextprotocol.io)
- [`rmcp` crate (GitHub)](https://github.com/modelcontextprotocol/rust-sdk)
- [Axum docs](https://docs.rs/axum)
- [Tokio docs](https://tokio.rs)
