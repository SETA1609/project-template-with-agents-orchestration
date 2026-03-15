# Workflow: update-docs

Sync documentation after code changes.

## Steps

1. Run `git diff --name-only HEAD~3` (or adjust range to cover the relevant commits).
2. For each changed file, check whether any of the following docs need updating:
   - **`docs/PROJECT_TREE.md`** — if directories or key files were added, moved, or removed
   - **`docs/DEPENDENCIES.md`** — if dependencies were added, removed, or upgraded
   - **`docs/INDEX.md`** — if new top-level docs or major sections were added
3. Update only the docs that are actually out of date. Do not rewrite docs that are still accurate.
4. Commit any doc changes using the `docs` commit type:
   ```
   docs: update PROJECT_TREE for new src/handlers directory

   Co-Authored-By: <your trailer here>
   ```

## Notes

- Do not update `PLAN.md` here — that is managed directly by `[C]`, `[K]`, or `[D]`.
- If you are unsure whether a doc change is needed, err on the side of updating — stale docs mislead other agents.
