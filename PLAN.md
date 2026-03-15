# PLAN — MCP Server (Go · HTTP · Docker · stdio)

## Overview

Build a **Model Context Protocol (MCP) server** in Go that supports **two transports**, selectable at runtime:

| Mode    | Transport                                          | Use-case                                                                      |
| ------- | -------------------------------------------------- | ----------------------------------------------------------------------------- |
| `stdio` | Standard input/output (JSON-RPC over stdin/stdout) | Local dev, subprocess launched by an AI client (e.g. Claude Desktop, VS Code) |
| `http`  | HTTP + Server-Sent Events                          | Docker deployment, remote access, multi-client                                |

The server will be packaged as a Docker image for HTTP mode and run as a plain binary for stdio mode.

---

## Goals

- Implement the MCP spec (2024-11-05 or later) over **HTTP + SSE** and **stdio** transports
- Write the server in **Go** for simplicity, fast compile times, and a small container footprint
- Ship a **multi-stage Docker image** (builder → distroless/static runtime)
- Keep the project structure clean so new tools/resources can be added with minimal boilerplate
- Let the user pick transport via a `MCP_TRANSPORT` env var (`stdio` | `http`) or a `--transport` CLI flag

---

## Stack

| Layer       | Choice                                   | Rationale                                                         |
| ----------- | ---------------------------------------- | ----------------------------------------------------------------- |
| Language    | Go (stable, latest)                      | Fast compile, simple concurrency, tiny binaries                   |
| MCP SDK     | `github.com/mark3labs/mcp-go`            | First-class Go MCP SDK; supports both stdio and SSE transports    |
| HTTP router | `net/http` stdlib                        | No extra dependency needed; mcp-go wraps it                       |
| Config      | `MCP_TRANSPORT` env + `--transport` flag | Parsed via `flag` stdlib or `github.com/spf13/cobra`              |
| Logging     | `log/slog` (stdlib, Go 1.21+)            | Structured JSON logging with zero extra deps                      |
| Container   | Docker (multi-stage)                     | `golang:1.23-alpine` builder → `gcr.io/distroless/static` runtime |
| Linting     | `golangci-lint`                          | Runs `gofmt`, `govet`, `staticcheck`, and more                    |

---

## Project Structure

```
my-mcp-agent-orchestrator/
├── go.mod
├── go.sum
├── Dockerfile               # used for http mode only
├── docker-compose.yml
├── .dockerignore
├── .env.example             # MCP_TRANSPORT=http  (or stdio)
├── cmd/
│   └── server/
│       └── main.go          # entry-point: reads config, dispatches to transport
├── internal/
│   ├── config/
│   │   └── config.go        # ENV + flag parsing (transport selection)
│   ├── server/
│   │   └── server.go        # MCP server setup, tool/resource registration
│   ├── tools/
│   │   ├── tools.go         # RegisterAll() — registers all tools with the server
│   │   └── echo.go          # Example tool: echo
│   └── resources/
│       ├── resources.go     # RegisterAll() — registers all resources with the server
│       └── health.go        # Example resource: server health
├── tests/
│   ├── stdio_test.go        # Spawn binary, talk over pipes
│   └── http_test.go         # Start HTTP server, send HTTP requests
└── docs/
    └── ...
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

- [x] `[C|K]` Design module boundaries and transport interface — define the contract all other agents build against
- [x] `[O|G]` Initialise Go module: `go mod init github.com/SETA1609/my-mcp-agent-orchestrator`
- [x] `[O|G]` Create `cmd/server/main.go` entry-point and `internal/` package skeleton
- [x] `[P|GF]` Add `go.mod` dependency: `github.com/mark3labs/mcp-go`
- [x] `[P|GF]` Add `.env.example` with `MCP_TRANSPORT` (`stdio`|`http`), `HOST`, `PORT`, `LOG_LEVEL`
- [x] `[P|O]` Configure `golangci-lint` (`.golangci.yml`) and `gofmt` as pre-commit check

### Phase 2 — Transport Layer

> `[C|K]` must complete the Phase 1 design step before any Phase 2 work begins.

#### 2a — stdio transport

- [x] `[O|G]` Use `mcp-go`'s built-in stdio transport: `server.ServeStdio(s)`

#### 2b — HTTP + SSE transport

- [x] `[K|C]` Confirm `mcp-go` SSE handler setup — `server.NewSSEServer(s, server.WithBaseURL(baseURL))` + `Start(addr)` / `Shutdown(ctx)`
- [x] `[O|G]` Implement `internal/` HTTP transport: spin up `mcp-go` SSE server on configured host/port
- [x] `[O|G]` Implement graceful shutdown (SIGTERM via `signal.NotifyContext`)

#### 2c — Transport dispatch

- [x] `[O|P]` Implement dispatch in `cmd/server/main.go`:
  ```go
  switch cfg.Transport {
  case "stdio":
      server.ServeStdio(s)
  case "http":
      sseServer := server.NewSSEServer(s, cfg.BaseURL)
      log.Fatal(sseServer.Start(cfg.Addr))
  }
  ```

### Phase 3 — MCP Server Core

- [x] `[C|K]` Design capability registration pattern — how tools and resources attach to `mcp.Server`
- [x] `[O|G]` Implement `internal/server/server.go`: create `mcp.NewServer(name, version)`, set capabilities
- [x] `[O|G]` Implement `internal/tools/tools.go`: `RegisterAll(s *mcp.Server)` calls each tool's register func
- [x] `[O|G]` Implement `internal/resources/resources.go`: `RegisterAll(s *mcp.Server)` calls each resource's register func
- [x] `[P|GF]` Add sample **echo tool** (`internal/tools/echo.go`) as reference implementation
- [x] `[P|GF]` Add sample **health resource** (`internal/resources/health.go`) as reference implementation

### Phase 4 — Configuration & Observability

- [x] `[O|G]` Implement `internal/config/config.go`: parse `MCP_TRANSPORT`, `HOST`, `PORT`, `LOG_LEVEL` from ENV; override with `--transport`, `--addr` flags
- [x] `[P|GF]` Wire `log/slog` with JSON output at startup, level controlled by `LOG_LEVEL`
- [x] `[P|O]` Add `GET /health` liveness endpoint (out-of-band, not part of MCP protocol)

> **Phases 5, 6, and 7 are independent and may be claimed and worked concurrently.**
> Each phase owns a distinct file area — no cross-phase file conflicts.
>
> | Phase | File area                              | Can start |
> |-------|----------------------------------------|-----------|
> | 5     | `Dockerfile`, `.dockerignore`, `docker-compose.yml` | now |
> | 6     | `internal/**/*_test.go`, `tests/`     | now       |
> | 7     | `.github/workflows/`                  | now       |

### Phase 5 — Docker Packaging

_File area: `Dockerfile`, `.dockerignore`, `docker-compose.yml`_

Steps 5.1–5.3 are independent of each other and may be claimed in parallel. Step 5.4 requires 5.1 to be complete.

- [ ] `[C|K]` **5.1** Write multi-stage `Dockerfile` (strategy decisions inline):
  1. **Builder**: `golang:1.23-alpine` — `CGO_ENABLED=0 go build -o /server ./cmd/server`
  2. **Runtime**: `gcr.io/distroless/static:nonroot` — copy `/server`; default `CMD` env `MCP_TRANSPORT=http`
- [ ] `[P|GF]` **5.2** Add `.dockerignore` (exclude `.git`, `tests/`, `docs/`)
- [O] `[P|O]` **5.3** Add `docker-compose.yml` for local development convenience
- [ ] `[D]` **5.4** Validate image size target: **< 20 MB**

> **stdio users** run the binary directly — no Docker required:
>
> ```bash
> MCP_TRANSPORT=stdio ./server
> # or: ./server --transport stdio
> ```

### Phase 6 — Tests

_File area: `internal/**/*_test.go`, `tests/`_

Steps 6.1 and 6.2 are independent and may be claimed in parallel.

- [ ] `[O|G]` **6.1** Write unit tests for each tool and resource handler (`_test.go` alongside source)
- [ ] `[K|C]` **6.2** Write integration tests — design harness inline (`tests/stdio_test.go`, `tests/http_test.go`): spin up binary over pipes (stdio) and `httptest.Server` (HTTP)

### Phase 7 — CI Pipeline

_File area: `.github/workflows/`_

- [ ] `[O|G]` **7.1** Write GitHub Actions workflow:
  - `gofmt -l .` (fail on unformatted files)
  - `go vet ./...`
  - `golangci-lint run`
  - `go test ./...`
  - `docker build` smoke-test
- [ ] `[D]` **7.2** Review and approve CI workflow before merging to `main`

---

## Open Questions — `[D]` to resolve

- [ ] `[D]` Authentication strategy (API key header? OAuth2 bearer token? mTLS?)
- [ ] `[D]` Rate-limiting requirements
- [ ] `[D]` Persistence layer needed? (e.g., tool outputs cached to Redis/SQLite)
- [ ] `[D]` Target deployment environment (bare Docker, Compose, Kubernetes, fly.io…)

---

## API Surface (MVP)

### HTTP mode

| Method | Endpoint   | Description                           |
| ------ | ---------- | ------------------------------------- |
| GET    | `/sse`     | Open SSE stream (MCP server → client) |
| POST   | `/message` | Send JSON-RPC message to MCP server   |
| GET    | `/health`  | Liveness probe (outside MCP)          |

### stdio mode

No HTTP endpoints. Reads newline-delimited JSON-RPC from `stdin`, writes responses to `stdout`.

---

## Key Design Decisions

1. **Dual transport, single binary** — same compiled binary serves both modes
2. **Runtime selection** — `MCP_TRANSPORT` env var or `--transport` CLI flag; CLI wins if both are set
3. **`mcp-go` SDK** — avoids reimplementing JSON-RPC framing, session state, and capability negotiation
4. **`net/http` stdlib** — no extra HTTP framework dependency; `mcp-go` wraps it cleanly
5. **Distroless static runtime** — Go produces truly static binaries (no libc); distroless image < 20 MB
6. **`log/slog`** — structured JSON logging with zero extra deps (Go 1.21+ stdlib)

---

## References

- [MCP Specification](https://spec.modelcontextprotocol.io)
- [`mcp-go` SDK (GitHub)](https://github.com/mark3labs/mcp-go)
- [Go stdlib `net/http`](https://pkg.go.dev/net/http)
- [distroless/static](https://github.com/GoogleContainerTools/distroless)
