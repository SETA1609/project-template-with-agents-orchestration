# Workflow: build

Run the project's build command and report the result.

## Steps

1. Run the build command for this project (e.g. `cargo build`, `npm run build`, `go build ./...`, `dotnet build`).
2. If the build **succeeds**: report success and the output summary.
3. If the build **fails**: report each error with file path, line number, and error message. Do not attempt to fix errors silently — surface them so the responsible agent can address them.

## Notes

- Do not modify any source files during this workflow.
- If the build command is not yet defined for this project, ask `[D]` (the human) to specify it before running.
