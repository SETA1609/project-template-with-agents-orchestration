# CODE_STYLE

> **Note:** This is a template. Replace these guidelines with your project's actual conventions.

## General

- Prefer clarity over cleverness
- Each function/method does one thing
- Keep files focused — split when a file grows beyond ~300 lines
- No commented-out code in commits (use `git stash` or a branch instead)

## Naming

- Use conventions idiomatic to your language (`camelCase` for JS/TS, `snake_case` for Rust/Python, `PascalCase` for Go/types/classes)
- Boolean names start with `is`, `has`, or `can`
- Avoid abbreviations unless universally known (`id`, `url`, `http`)

## Error Handling

- Never silently swallow errors
- Surface errors at the boundary closest to the caller
- Prefer explicit error returns over exceptions where the language allows

## Tests

- Every new public function gets at least one test
- Test files live next to the source or in a sibling `tests/` directory
- Test names describe the scenario: `test_returns_error_when_config_missing`

## Imports / Dependencies

- No circular dependencies
- Group imports: standard library → third-party → internal
- Remove unused imports before committing

## Comments

- Comment *why*, not *what*
- Keep comments current — stale comments are worse than none
