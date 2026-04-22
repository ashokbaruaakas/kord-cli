# kord-cli — Agent Instructions

## Project overview

`kord-cli` is a high-performance, open-source CLI tool built in Go using the [Cobra](https://github.com/spf13/cobra) framework. It acts as a digital "cord," seamlessly tying together Git development workflows, task management platforms, and LLM-powered automation. The project is in early development; the root command scaffolding is in place and new subcommands are the primary growth area.

## Build & run

```sh
go build -o kord-cli .   # compile binary
go run .                 # run without compiling
go test ./...            # run all tests
```

## Testing workflow

- Follow TDD when implementing new behavior: write failing tests first, then implement the minimum code to pass.
- Test locations:
  - `cmd/*_test.go` for unit tests around Cobra commands
  - `tests/*_test.go` for feature/integration tests across CLI workflows
- Use `github.com/stretchr/testify` (`assert`, `require`) for readable assertions.
- Useful commands:
  - `go test ./... -v` for full verbose test runs
  - `go test ./... -cover` for package coverage
  - `watchexec -c -- go test ./... -v` for continuous test runs
  - `watchexec -c -e go -- go test ./... -v` to trigger only on Go file changes

## Architecture

- `main.go` — entry point; calls `cmd.Execute()`
- `cmd/root.go` — defines `rootCmd` (the base `kord-cli` command)
- `cmd/<name>.go` — each subcommand lives in its own file inside `cmd/`
- `internal/git/` — Git operations abstraction
- `internal/taskmanager/` — Task platform adapter (currently Notion only; design for extensibility)
- `internal/llm/` — LLM integrations
- `internal/config/` — Loads `~/.kord/config.json` and environment variables

## Planned commands

The core workflow comprises these subcommands (in typical usage order):

1. **`kord setup`** — Interactive wizard to configure kord (Git, task management platform, LLM connections).
2. **`kord start <task-id/reference>`** — Automate task initiation: fetch task info, generate branch name, create & checkout branch, update task status.
3. **`kord commit`** — Auto-generate commit message from staged changes.
4. **`kord submit`** — Create PR with auto-detected title and description, submit to platform.
5. **`kord finish`** — Verify GitHub Actions pass, merge PR, update task status.
6. **`kord release`** — Auto-detect version tag, generate title and user-friendly release notes, create release draft.
7. **`kord publish`** — Finalize and publish release, update task statuses, sync with platforms.

## Adding a subcommand

1. Create `cmd/<name>.go` with a `var <name>Cmd = &cobra.Command{…}`.
2. Register it in `init()` with `rootCmd.AddCommand(<name>Cmd)`.
3. Follow the same file header pattern used in `cmd/root.go`.

Do **not** modify `main.go`; all command wiring belongs in `cmd/`.

### Integration considerations

Commands interact with external systems (Git, Notion, LLM, GitHub Actions). When implementing:

- Load configuration via `kord setup` state (`~/.kord/config.json` or environment variables).
- **Task management**: Use the `internal/taskmanager` adapter; currently only **Notion** is supported. Design the adapter interface generically so additional platforms (Jira, Linear, etc.) can be added later without changing command code.
- Use `RunE` to return errors; wrap external API calls with clear error messages.
- Use flags for optional overrides (e.g., `--branch-name`, `--no-pr-draft`).
- Log significant steps for user clarity (fetching task info, branch creation, API calls).

## Conventions

- Module path: `kord-cli` (see [go.mod](go.mod))
- Go version: 1.26.2
- Cobra version: v1.10.2 — use `cobra.Command` fields (`Use`, `Short`, `Long`, `RunE`) and `pflag` for flags.
- Prefer `RunE` over `Run` so commands can return errors to be handled by Cobra.
- Use `os.Exit(1)` only in `Execute()`; subcommands should return errors.
- **Configuration**: Store user config from `kord setup` in a standard location (e.g., `~/.kord/config.json`). Load it at command start.
- **External integrations**: Abstract Git, task management, and LLM APIs into separate packages (e.g., `internal/git`, `internal/taskmanager`, `internal/llm`) for testability and reuse.
- **Error messages**: Be specific about failures (e.g., "Failed to fetch task #123 from Notion: rate limit exceeded" vs. "Error").
- **User feedback**: Print status updates for long operations (fetch, API calls, branch creation) so users know progress.
