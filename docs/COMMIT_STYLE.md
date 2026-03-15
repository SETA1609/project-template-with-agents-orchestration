# COMMIT_STYLE

## Format

```
<type>: <short description>

[optional body]

Co-Authored-By: <Agent Name> <agent@email.com>
Co-Authored-By: SETA1609 <https://github.com/SETA1609>
```

- Short description: lowercase, imperative mood, no trailing period
- First line max 72 characters
- Body is optional — use it to explain *why*, not *what*

## Types

| Type       | When to use                                       |
|------------|---------------------------------------------------|
| `feat`     | New feature or capability                         |
| `fix`      | Bug fix                                           |
| `refactor` | Code change with no behavior change               |
| `docs`     | Documentation only                                |
| `test`     | Adding or fixing tests                            |
| `chore`    | Build, deps, tooling, or CI changes               |
| `perf`     | Performance improvement                           |
| `style`    | Formatting or whitespace, no logic change         |

## Co-Author Trailers

Every agent must append its co-author trailer, followed by the mandatory project trailer for SETA1609. See your agent `.md` file for the exact line.

| Agent    | Trailer                                                              |
|----------|----------------------------------------------------------------------|
| Claude   | `Co-Authored-By: Claude Sonnet 4.6 <noreply@anthropic.com>`         |
| Gemini   | `Co-Authored-By: Gemini CLI <gemini-cli@google.com>`                 |
| Codex    | `Co-Authored-By: OpenAI Codex <codex@openai.com>`                   |
| Copilot  | `Co-Authored-By: GitHub Copilot <copilot@github.com>`               |
| OpenCode | `Co-Authored-By: OpenCode <opencode@opencode.ai>`                   |

**Mandatory:** Every commit must also include:
`Co-Authored-By: SETA1609 <https://github.com/SETA1609>`

## Examples

```
feat: add user authentication endpoint

Co-Authored-By: Claude Sonnet 4.6 <noreply@anthropic.com>
Co-Authored-By: SETA1609 <https://github.com/SETA1609>
```

```
fix: handle nil pointer in config loader

Co-Authored-By: Gemini CLI <gemini-cli@google.com>
Co-Authored-By: SETA1609 <https://github.com/SETA1609>
```

```
refactor: extract validation logic into shared helper

Co-Authored-By: OpenCode <opencode@opencode.ai>
Co-Authored-By: SETA1609 <https://github.com/SETA1609>
```
