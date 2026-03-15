# MCP Agent Orchestrator Template (Go)

This repository is a template for building an MCP server while coordinating multiple AI agents with explicit ownership, review, and merge rules.

It gives you two things in one project:
- an MCP server implementation in Go
- a structured agent workflow (`AGENTS.md` + `PLAN.md`) so work can be delegated safely

## What this template is for

Use this template when you want to:
- build an MCP server quickly with clear extension points for tools/resources
- split work across multiple agents without stepping on each other
- keep architecture decisions explicit and implementation tasks parallel

## How the agent workflow works

Start here:
- `AGENTS.md`: global coordination rules, role boundaries, merge strategy
- `PLAN.md`: active implementation checklist with `[Primary|Fallback]` ownership tags
- `docs/agents/*.md`: behavior and constraints per agent

Core idea:
- architecture gates first (owned by `[C|K]`)
- implementation after interfaces are defined
- one step claimed at a time (`[ ] -> [symbol] -> [x]`)

## MCP server overview

The server runs in a single binary and supports two transports selected at runtime:
- `stdio`: JSON-RPC over stdin/stdout for local client integrations
- `http`: HTTP + SSE for remote/docker usage

Current runtime behavior:
- transport selected via `MCP_TRANSPORT` or `--transport`
- bind address via `HOST` + `PORT` or `--addr host:port`
- JSON logs via `log/slog`, level from `LOG_LEVEL`
- HTTP mode exposes:
  - `GET /sse`
  - `POST /message`
  - `GET /health`

## Project structure

- `cmd/server/main.go`: app entrypoint and transport dispatch
- `internal/config/`: env/flag config loading
- `internal/server/`: MCP server construction + capability registration
- `internal/tools/`: MCP tool modules and registry
- `internal/resources/`: MCP resource modules and registry
- `internal/transport/`: stdio and HTTP/SSE backends
- `docs/`: process and coding standards

## Extend the server

Add a tool:
1. create a file in `internal/tools/`
2. define and register it using `s.AddTool(...)`
3. call its register function from `internal/tools/tools.go`

Add a resource:
1. create a file in `internal/resources/`
2. define and register it using `s.AddResource(...)`
3. call its register function from `internal/resources/resources.go`

No changes are required in `internal/server/server.go` when following this pattern.

## Requirements

- Go 1.23+ recommended
- Docker (optional, for HTTP container mode)
- Docker Compose (optional, for local HTTP stack)

## Run locally (binary)

1) Copy environment defaults:

```bash
cp .env.example .env
```

2) Run in stdio mode:

```bash
MCP_TRANSPORT=stdio go run ./cmd/server --transport stdio
```

3) Run in HTTP mode:

```bash
MCP_TRANSPORT=http HOST=127.0.0.1 PORT=3000 go run ./cmd/server
```

4) Health check (HTTP mode):

```bash
curl http://127.0.0.1:3000/health
```

## Run with Docker

Build image:

```bash
docker build -t my-mcp-agent-orchestrator:local .
```

Run container:

```bash
docker run --rm -p 3000:3000 -e MCP_TRANSPORT=http my-mcp-agent-orchestrator:local
```

Run with Compose:

```bash
docker compose up --build
```

## Useful dev commands

```bash
gofmt -w ./...
go test ./...
docker compose config
```

## Notes

- `PLAN.md` and `REVIEW.md` are workflow artifacts; follow `AGENTS.md` for the exact process.
- For shared/core changes, use a PR and keep commit history aligned with `docs/COMMIT_STYLE.md`.
