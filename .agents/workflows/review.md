# Workflow: review

Review recently changed files for code quality and architectural compliance.

## Steps

1. Run `git diff --name-only HEAD~3` (or adjust the range to cover the relevant commits).
2. For each changed source file, review for:
   - **Correctness**: logic errors, off-by-one, unhandled edge cases
   - **Quality**: dead code, unnecessary complexity, missing error handling
   - **Style**: deviations from [`docs/CODE_STYLE.md`](../../docs/CODE_STYLE.md)
   - **Architecture**: violations of role boundaries (e.g. `[P]` introducing a new abstraction without `[C]` authorisation)
3. Write findings to `REVIEW.md` at the project root (overwrite if it already exists).
4. Group findings by file. For each issue include: file path, line number (if applicable), severity (`error` / `warning` / `suggestion`), and a brief description.
5. If no issues are found, write a single line: `No issues found.`

## Notes

- Do not modify source files during this workflow — report only.
- `REVIEW.md` is ephemeral; it is not committed to the repository.
