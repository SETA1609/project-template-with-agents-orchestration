# AGENTS

This is the master coordination document. Every agent working in this repository **must** read this file before touching any code.

---

## Agent Roster

| Symbol  | Agent          | Tool                  | Detail file                                      | Role summary                                              |
|---------|----------------|-----------------------|--------------------------------------------------|-----------------------------------------------------------|
| `[C]`   | Claude         | Claude Code CLI       | [docs/agents/CLAUDE.md](docs/agents/CLAUDE.md)   | Architecture, planning, new abstractions, framework design |
| `[G]`   | Gemini         | Gemini CLI            | [docs/agents/GEMINI.md](docs/agents/GEMINI.md)   | Refactoring, extending abstractions, UI/UX, minor tasks   |
| `[K]`   | Codex          | OpenAI Codex          | [docs/agents/CODEX.md](docs/agents/CODEX.md)     | Planning, complex algorithms, new systems                 |
| `[P]`   | Copilot        | GitHub Copilot        | [docs/agents/COPILOT.md](docs/agents/COPILOT.md) | Boilerplate, pattern completion, simple utilities         |
| `[O]`   | OpenCode       | OpenCode CLI          | [docs/agents/OPENCODE.md](docs/agents/OPENCODE.md) | Full-stack implementation, code generation              |
| `[GF]`  | Grok Code Fast | Grok (fast mode)      | —                                                | Speed-first code generation, quick single-file fixes      |
| `[GR]`  | Grok 4.20      | Grok 4.20             | —                                                | Reasoning-heavy tasks, deep code analysis                 |
| `[PK]`  | Pickle         | OpenCode / Pickle     | —                                                | Experimental, rapid prototyping                           |
| `[D]`   | Human          | —                     | —                                                | Final decisions, unblocking, planning approval            |

---

## Task-Type Assignment Table

Use this to quickly find which agent(s) should own a given type of work.

| Task type                          | Primary   | Fallback   | Notes                                                         |
|------------------------------------|-----------|------------|---------------------------------------------------------------|
| Architecture / new abstractions    | `[C]`     | `[K]`      | Must produce an interface before others implement             |
| New algorithm / data structure     | `[K]`     | `[C]`      | Accompanies with a proposal for `[D]` review                  |
| Full-stack feature implementation  | `[O]`     | `[G]`      | Only after `[C]`/`[K]` defines the interface                  |
| Boilerplate / repetitive patterns  | `[P]`     | `[O]`      | No novel design; follow existing patterns only                |
| Refactoring existing code          | `[G]`     | `[O]`      | Must not change public API without `[C]`/`[K]` sign-off       |
| UI / view-layer features           | `[G]`     | `[O]`      | Must not alter shared-state interfaces                        |
| Quick single-file fix or patch     | `[GF]`    | `[P]`      | In-and-out; no design changes                                 |
| Deep reasoning / analysis          | `[GR]`    | `[K]`      | Use when problem requires multi-step logical deduction        |
| Rapid prototype / throwaway spike  | `[PK]`    | `[GF]`     | Not for production; produces a working sketch only            |
| Tests (unit + integration)         | All       | —          | Each agent writes tests for its own changes                   |
| `docs/` updates                    | All       | —          | Keep in sync with code; agent that changed code updates docs  |
| `PLAN.md` writes / updates         | `[C]`     | `[K]`, `[D]` | Only after `[D]` approves the plan                          |
| Irreversible architectural decision| `[D]`     | —          | Human only; no fallback                                       |

---

## Before You Start

1. Read `PLAN.md` — understand the active work manifest before touching any files.
2. Read your detail file in `docs/agents/` for role boundaries, co-author trailer, and context files.
3. Read the relevant `docs/` files for the layer you will be working on.
4. Do **not** start work on a step that is already claimed by another agent symbol.

---

## Step Claiming Protocol

`PLAN.md` lists all work as numbered steps. Each step starts unclaimed `[ ]` and lists the agents eligible to work it.

### Eligibility notation

Steps use `[Primary|Fallback]` to show which agents may claim them, in priority order:

```
- [ ] [C|K]     → C should take this; K may take it if C is unavailable
- [ ] [O|G]     → O should take this; G may take it if O is unavailable
- [ ] [GF|P]    → Grok Code Fast should take this; Copilot may fall back
- [ ] [D]       → human only, no fallback
```

### Claiming a step

1. **Check eligibility** — only claim a step if your symbol appears in the step's eligibility list.
2. **Claim** — replace `[ ]` with *your own symbol*, then commit that change alone before doing any other work.
3. **Work** the step.
4. **Complete** — replace your symbol with `[x]` and commit before moving to the next step.
5. Only one agent may hold a symbol on any given step at a time.
6. Never skip a step unless it is explicitly marked optional.

Example progression:
```
- [ ] [C|K]   →  - [C] [C|K]   →  - [x] [C|K]    (Claude claimed and completed)
- [ ] [O|G]   →  - [G] [O|G]   →  - [x] [O|G]    (Gemini took it as fallback)
- [ ] [GF|P]  →  - [GF] [GF|P] →  - [x] [GF|P]   (Grok Code Fast claimed and completed)
```

---

## Role Boundaries

### `[C]` Claude
- Designs new abstractions, base classes, and framework-level systems
- Writes and maintains `PLAN.md`
- Unblocks other agents by defining interfaces they can build against
- **Cannot** be bypassed — `[G]`, `[P]`, and `[O]` must not introduce new abstractions without a `[C]` or `[K]` handoff step first
- See [docs/agents/CLAUDE.md](docs/agents/CLAUDE.md) for full detail

### `[G]` Gemini
- Refactors and extends existing abstractions
- Implements UI/view-layer features
- Carries out clearly scoped minor tasks
- **Cannot** introduce new abstractions or design framework systems without a prior `[C]`/`[K]` step
- See [docs/agents/GEMINI.md](docs/agents/GEMINI.md) for full detail

### `[K]` Codex
- Designs and implements complex algorithms and data structures
- Produces detailed implementation proposals for human review
- Plans new subsystems with clear interfaces for other agents to build against
- See [docs/agents/CODEX.md](docs/agents/CODEX.md) for full detail

### `[P]` Copilot
- Fills in boilerplate and repeat patterns already established
- Implements simple, self-contained utilities (single file)
- **Cannot** design APIs, introduce abstractions, or touch shared/critical code without prior `[C]`/`[K]` authorization
- See [docs/agents/COPILOT.md](docs/agents/COPILOT.md) for full detail

### `[O]` OpenCode
- Implements features end-to-end once interfaces are defined by `[C]` or `[K]`
- Handles code generation and scaffolding within existing patterns
- **Cannot** change public interfaces without a prior `[C]`/`[K]` design step
- Must use its **actual running model name** in the co-author trailer — not the generic "OpenCode" label
- See [docs/agents/OPENCODE.md](docs/agents/OPENCODE.md) for full detail

### `[GF]` Grok Code Fast
- Speed-first: writes or fixes single files quickly with minimum overhead
- Best for quick patches, filling in function bodies, one-off scripts
- **Cannot** introduce new abstractions or design cross-file patterns

### `[GR]` Grok 4.20
- Reasoning-heavy tasks: multi-step deduction, root-cause analysis, algorithm design
- Can propose design changes but must hand off to `[C]`/`[K]` for implementation contracts
- Use when a task requires careful logical reasoning, not just code generation

### `[PK]` Pickle
- Rapid prototyping and experimental spikes
- Output is a working sketch — not intended for production without review by `[C]` or `[K]`
- **Cannot** merge prototype code to `main` without `[D]` approval

### `[D]` Human
- Approves or rejects plans before implementation starts
- Resolves conflicts between agents
- Makes all irreversible architectural decisions

---

## File Ownership

| Path pattern              | Responsible agents        | Notes                                         |
|---------------------------|---------------------------|-----------------------------------------------|
| Core / shared libraries   | `[C]`, `[K]`              | Breaking changes require `[D]` review         |
| UI / view layer           | `[G]`, `[O]`              | Must not alter shared-state interfaces        |
| Quick fixes / patches     | `[GF]`, `[P]`             | Single-file scope only                        |
| Prototypes / spikes       | `[PK]`                    | Must be in a throwaway branch                 |
| Tests                     | All                       | Each agent writes tests for its own changes   |
| `PLAN.md`                 | `[C]`, `[K]`, `[D]`       | Updated before and after each work session    |
| `AGENTS.md`               | `[C]`, `[D]`              | Structural changes need `[D]` approval        |
| `docs/agents/`            | All (own file only)       | Each agent maintains their own detail file    |
| `docs/`                   | All                       | Keep in sync with code changes                |

---

## Workflows

Reusable workflows live in `.agents/workflows/`. Run them as needed:

- [`build.md`](.agents/workflows/build.md) — build and report errors
- [`commit.md`](.agents/workflows/commit.md) — stage and commit with correct format and co-author trailers
- [`review.md`](.agents/workflows/review.md) — review changed files for quality and architecture violations
- [`update-docs.md`](.agents/workflows/update-docs.md) — sync `docs/PROJECT_TREE.md` and other docs after code changes

---

## Git Rules

- Commit format: `<type>: <description>` — see [`docs/COMMIT_STYLE.md`](docs/COMMIT_STYLE.md)
- Every commit must include your agent's co-author trailer (see your detail file in `docs/agents/`)
- Never force-push to `main`
- Open a PR for any change that touches shared or framework-level code

### Branch & merge strategy (squash + rebase + fast-forward)

Every task is worked in its own short-lived branch. Agents work simultaneously — no locking required. The rebase step before merge is what keeps `main` linear without coordination.

```
# 1. Branch off main before touching any files
git checkout main && git pull --ff-only
git checkout -b <agent>/<short-description>   # e.g. c/transport-interface

# 2. Do your work — commit as often as needed while working
git add <files>
git commit -m "wip: ..."

# 3. Squash all branch commits into one
git rebase -i main   # mark every commit except the first as "squash"
# Write the final commit message following COMMIT_STYLE.md (with co-author trailers)

# 4. Rebase onto the latest origin/main (handles any merges by other agents while you worked)
git fetch origin
git rebase origin/main

# 5. Fast-forward main — guaranteed clean because you just rebased
git checkout main
git pull --ff-only
git merge --ff-only <agent>/<short-description>

# 6. Push and delete the branch
git push
git branch -d <agent>/<short-description>
```

**Why this works without locking:**
Steps 3–5 are the key. Squash first so the rebase in step 4 moves a single clean commit — not a pile of wip commits. If another agent merged while you were working, `rebase origin/main` replays your one commit on top of their work. As long as agents respect file ownership, rebases are conflict-free.

**If the rebase in step 4 has a conflict:**
1. Resolve the conflict in the affected file
2. `git add <file> && git rebase --continue`
3. If unresolvable, `git rebase --abort` and coordinate with `[D]`

**Rules:**
- One branch per PLAN.md step — do not batch multiple steps into one branch
- Branch name format: `<agent-symbol-lowercase>/<kebab-description>` (e.g. `c/transport-interface`, `o/stdio-impl`)
- Always `git fetch origin` before rebasing — never rebase onto stale local main
- The squashed commit message must follow `docs/COMMIT_STYLE.md` and include all co-author trailers
- Never merge with `--no-ff` — no merge commits on `main`
- Never force-push `main`
