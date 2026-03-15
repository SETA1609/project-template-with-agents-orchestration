# Workflow: commit

Stage, commit, and merge a completed step back to `main` using the squash + fast-forward strategy.

## Steps

### While working (on your task branch)

1. Verify you are on your task branch — never commit directly to `main`.
2. Run `git status` to review what has changed.
3. Stage only the files relevant to the completed step — do not stage unrelated changes.
4. Commit with a `wip:` prefix as often as needed while working:
   ```
   git commit -m "wip: add stdio handler loop"
   ```

### When the step is done (ready to merge)

5. Squash all commits on the branch into one with `git rebase -i main`:
   - Mark every commit except the first as `squash` (or `s`).
   - Write the final commit message following [`docs/COMMIT_STYLE.md`](../../docs/COMMIT_STYLE.md):
     ```
     <type>: <short description>

     Co-Authored-By: <your agent trailer>
     Co-Authored-By: SETA1609 <https://github.com/SETA1609>
     ```
6. Fast-forward `main` onto the squashed branch:
   ```
   git checkout main
   git merge --ff-only <your-branch>
   ```
7. Delete the branch:
   ```
   git branch -d <your-branch>
   ```
8. Confirm with `git log --oneline -3`.

## Branch naming

```
<agent-symbol-lowercase>/<kebab-description>
```

Examples: `c/transport-interface`, `o/stdio-impl`, `p/env-example`

## Notes

- Never use `--no-ff` — no merge commits on `main`.
- If `--ff-only` fails (main moved ahead), rebase your branch onto the latest `main` first, then retry.
- Never use `--no-verify` to skip hooks.
- If a pre-commit hook fails, fix the underlying issue and create a new commit.
