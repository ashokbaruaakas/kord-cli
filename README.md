# kord-cli

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/go-1.26.2-00add8.svg)](https://golang.org/)
[![Status](https://img.shields.io/badge/status-early%20development-orange.svg)](#status)

A high-performance CLI tool that acts as a digital "cord," seamlessly tying together Git development workflows, task management platforms, and LLM-powered automation.

## Overview

`kord-cli` streamlines your development lifecycle by automating the repetitive tasks between Git, your task management system, and AI-powered tooling. From task initialization to release publication, kord handles the orchestration, reducing context switching and human error. Currently supports **Notion** as the task management adapter, with more platforms planned.

## Features

- 🚀 **Automated Task Initiation** — Fetch task details, auto-generate branch names, create branches with a single command
- ✍️ **Intelligent Commit Messages** — Generate contextual commit messages from staged changes
- 🔄 **Pull Request Management** — Auto-detect task context, create PRs with rich descriptions
- ✅ **Workflow Automation** — Verify CI/CD status, merge PRs, update task statuses automatically
- 🏷️ **Release Orchestration** — Auto-detect version tags, generate release notes, draft releases
- 🔗 **Platform Integration** — Connect to Git, GitHub, and Notion (more task platforms coming soon)
- 🤖 **LLM-Powered** — Leverage AI for intelligent suggestions and automations
- 🧪 **TDD-Ready Test Setup** — Unit and feature test scaffolding with `testify` assertions

## Quick Start

### Installation

```bash
# Clone the repository
git clone https://github.com/ashokbaruaakas/kord-cli.git
cd kord-cli

# Build the binary
go build -o kord-cli .

# Optionally, install to your PATH
sudo mv kord-cli /usr/local/bin/
```

### First Command

```bash
# Configure kord with your platforms (Git, task management, LLM)
kord setup

# Start work on a task
kord start TASK-123

# Write changes, then auto-generate and commit
git add .
kord commit

# Create and submit a pull request
kord submit

# After review and CI/CD passes, merge and finish
kord finish
```

## Usage

### Commands

| Command                | Purpose                                                        |
| ---------------------- | -------------------------------------------------------------- |
| `kord setup`           | Interactive configuration wizard for platforms and credentials |
| `kord start <task-id>` | Initialize work on a task (branch creation, status updates)    |
| `kord commit`          | Generate and apply a commit message from staged changes        |
| `kord submit`          | Create and submit a pull request with auto-detected metadata   |
| `kord finish`          | Verify CI/CD status, merge PR, update task status              |
| `kord release`         | Create a release draft with auto-detected version and notes    |
| `kord publish`         | Finalize and publish release, update all connected platforms   |

### Command Details

#### `kord setup`

Runs an interactive wizard to connect kord with your platforms. Stores credentials and preferences in `~/.kord/config.json`.

```bash
kord setup
```

#### `kord start <task-id>`

Fetches task details from your task management platform, generates a descriptive branch name using LLM, creates and checks out the branch, and updates the task status to _In Progress_.

```bash
kord start PROJ-42
kord start PROJ-42 --branch-name my-custom-branch
```

#### `kord commit`

Inspects staged changes, generates a concise and conventional commit message using LLM, and commits them.

```bash
git add .
kord commit
```

#### `kord submit`

Detects the current task from the active branch, creates a pull request with an auto-generated title and description, and submits it.

```bash
kord submit
kord submit --no-pr-draft
```

#### `kord finish`

Checks the latest PR for the current task, waits for all GitHub Actions to pass, merges the PR, and updates the task status to _Done_.

```bash
kord finish
```

#### `kord release`

Auto-detects the next version tag, generates a developer changelog and a user-friendly release note using LLM, and creates a release draft on GitHub.

```bash
kord release
```

#### `kord publish`

Publishes the latest release draft, syncs the release status to connected platforms, and marks related tasks as released.

```bash
kord publish
```

## Architecture

```
kord-cli/
├── main.go              # Entry point
├── cmd/
│   ├── root.go          # Root command definition
│   ├── setup.go         # kord setup subcommand
│   ├── start.go         # kord start subcommand
│   └── ...              # Additional subcommands
├── internal/
│   ├── git/             # Git operations
│   ├── taskmanager/     # Task platform integrations
│   ├── llm/             # LLM integrations
│   └── config/          # Configuration management
├── go.mod              # Go module definition
└── LICENSE             # License file
```

**Key Design Principles:**

- Each subcommand lives in `cmd/<name>.go`
- External integrations are abstracted into `internal/` packages for reusability and testability
- Configuration is loaded from `~/.kord/config.json` and environment variables
- Commands return errors via `RunE` for proper Cobra error handling

## Configuration

kord stores its configuration at `~/.kord/config.json`. This file is created and populated during `kord setup`. You can also override settings using environment variables.

```json
// ~/.kord/config.json (example)
{
  "git": {
    "provider": "github",
    "token": "<your-github-token>"
  },
  "task_manager": {
    "provider": "notion",
    "token": "<your-notion-integration-token>",
    "database_id": "<your-notion-database-id>"
  },
  "llm": {
    "provider": "openai",
    "model": "gpt-4o",
    "api_key": "<your-api-key>"
  }
}
```

> **Note:** Never commit `~/.kord/config.json` to version control. kord will never read credentials from the project directory.

## Roadmap

### Phase 1 (Current)

- [x] Root command scaffolding
- [x] Baseline test setup (unit + feature tests)
- [ ] `kord setup` — Interactive configuration wizard
- [ ] `kord start` — Task initialization automation
- [ ] `kord commit` — Smart commit message generation

### Phase 2

- [ ] `kord submit` — PR creation and submission
- [ ] `kord finish` — Workflow completion and status updates
- [ ] Support for additional task platforms (Jira, Linear, GitHub Issues)

### Phase 3

- [ ] `kord release` — Release draft generation
- [ ] `kord publish` — Release publication and synchronization
- [ ] Advanced LLM features (release notes generation, context summarization)

## Development

### Prerequisites

- Go 1.26.2 or later
- Git

### Build and Test

```bash
# Build the binary
go build -o kord-cli .

# Run the CLI
go run . <command>

# Run tests
go test ./...

# Run tests with coverage
go test ./... -cover
```

## Testing

This project follows TDD and includes both unit and feature/integration tests. For full testing commands, watch-mode setup, conventions, and examples, see [TESTING.md](TESTING.md).

### Code Conventions

- Prefer `RunE` over `Run` for proper error handling
- Use `os.Exit(1)` only in `Execute()`; subcommands return errors
- Wrap external API calls with clear, specific error messages
- Log significant steps for user feedback
- Use flags for optional overrides (e.g., `--branch-name`, `--no-pr-draft`)

## Contributing

Contributions are welcome! Here's how to get started:

1. **Fork the repository** and clone it locally
2. **Create a feature branch** — `git checkout -b feat/your-feature`
3. **Follow the code conventions** described in the [Development](#development) section
4. **Write tests** for new features and ensure `go test ./...` passes
5. **Commit your changes** — consider using `kord commit` once it's available 😄
6. **Open a pull request** with a clear title and description

For bug reports and feature requests, please [open an issue](https://github.com/ashokbaruaakas/kord-cli/issues) first to discuss before submitting a PR.

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.

## Status

🏗️ **Current Status:** Early Development
The root command scaffolding is complete and stable. Primary focus is implementing the planned subcommands. Expect breaking changes and API refinements as the project evolves.

## Support

- **Bug reports & feature requests** — [Open an issue](https://github.com/ashokbaruaakas/kord-cli/issues)
- **Questions & discussions** — [GitHub Discussions](https://github.com/ashokbaruaakas/kord-cli/discussions)
- **Security vulnerabilities** — Please email [ashokbaruaakas@gmail.com](mailto:ashokbaruaakas@gmail.com) directly instead of opening a public issue

---

**Made with ❤️ by the kord-cli team**
