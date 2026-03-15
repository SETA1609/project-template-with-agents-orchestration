# Workflow: commit

Squash, rebase, and fast-forward merge a completed step to `main`.
Agents run this independently — no locking or coordination needed.

## While working (on your task branch)

1. Verify you are on your task branch — never commit directly to `main`.
2. `git status` — review what has changed.
3. Stage only files relevant to this step.
4. Commit freely with `wip:` prefix while working:
   ```
   git commit -m "wip: ..."
   ```

## When the step is done

### Step A — squash
Collapse all wip commits into one clean commit:
```
git rebase -i main
```
Mark every commit except the first as `squash`. Write the final message:
```
<type>: <short description>

Co-Authored-By: <your agent trailer>
Co-Authored-By: SETA1609 <https://github.com/SETA1609>
```

### Step B — rebase onto latest origin/main
Bring your single squashed commit on top of whatever other agents merged while you worked:
```
git fetch origin
git rebase origin/main
```
- If clean (no conflict): continue to Step C.
- If conflict: resolve, `git add <file>`, `git rebase --continue`. If unresolvable, `git rebase --abort` and raise with `[D]`.

### Step C — fast-forward merge
```
git checkout main
git pull --ff-only
git merge --ff-only <your-branch>
git push
```

### Step D — delete the branch
```
git branch -d <your-branch>
```

### Step E — confirm
```
git log --oneline -3
```

## Branch naming

```
<agent-symbol-lowercase>/<kebab-description>
```

Examples: `c/transport-interface`, `o/stdio-impl`, `p/env-example`

## Notes

- Never use `--no-ff` — no merge commits on `main`.
- Never use `--no-verify` to skip hooks.
- If a pre-commit hook fails, fix the issue and create a new commit before squashing.
- `git pull --ff-only` on `main` in Step C will fail if someone pushed between your fetch and merge — just re-run from Step B.
