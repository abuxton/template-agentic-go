# Template Agentic Go

Spec-driven Go starter for agentic applications with Claude and OpenAI
support. This repository provides a small provider-neutral CLI, reproducible
development tooling, governance guidance, and OpenSpec onboarding. It is a
template rather than an application framework: create a repository from it,
then add the domain code, specifications, tests, and deployment configuration
for the new project.

## What is included

| Path | Purpose |
| --- | --- |
| `cmd/agent/` | Example CLI entrypoint |
| `internal/agent/` | Claude/OpenAI configuration and HTTP adapters |
| `Taskfile.yml` | Build, test, lint, vulnerability, and CI commands |
| `.devcontainer/` | Reproducible Go, Task, and Node/npm environment |
| `.vscode/` | Workspace settings, tasks, debugging, and extensions |
| `.github/copilot-instructions.md` | Copilot operating instructions |
| `.github/workflows/ci.yml` | SHA-pinned GitHub Actions CI |
| `AGENTS.md` | Agent operating contract |
| `CONTRIBUTING.md` | Contribution and pull request guidance |
| `docs/agentic-workflow.md` | Capabilities, governance, and verification |
| `docs/openspec.md` | OpenSpec initialization and change lifecycle |
| `.env.example` | Provider configuration names and defaults |
| `skills-lock.json` | Pinned, reproducible project skill sources |

Generated OpenSpec integrations, installed skill payloads, and credentials are
not part of the template's source. Only `skills-lock.json` is committed for
project skills.

## Create a project from the template

### GitHub web flow

1. Open [template-agentic-go](https://github.com/abuxton/template-agentic-go).
2. Select **Use this template**, then **Create a new repository**.
3. Choose the owner, repository name, visibility, and description.
4. Clone the new repository:

   ```sh
   git clone git@github.com:<owner>/<repository>.git
   cd <repository>
   ```

### GitHub CLI flow

```sh
gh repo create <owner>/<repository> --public --template \
  abuxton/template-agentic-go
gh repo clone <owner>/<repository>
cd <repository>
```

Use `--private` instead of `--public` when appropriate.

## First-time setup

### 1. Open the development environment

Open the repository in VS Code and choose **Reopen in Container**. The
development container provides Go 1.27, Task, and Node/npm. It also installs
the recommended Go, Task, GitHub Copilot, Claude Code, and OpenAI extensions.

For a host setup, install Go 1.27 or newer, [Task](https://taskfile.dev/),
Node.js/npm, Git, and an OpenSpec-capable agent client.

### 2. Configure a provider

```sh
cp .env.example .env
```

Set `AGENT_PROVIDER`, `AGENT_MODEL`, and the matching API key:

| Provider | `AGENT_PROVIDER` | Required key |
| --- | --- | --- |
| Claude | `claude` | `ANTHROPIC_API_KEY` |
| OpenAI | `openai` | `OPENAI_API_KEY` |

Base URLs can be overridden with `ANTHROPIC_BASE_URL` and
`OPENAI_BASE_URL`. Keep `.env` private; never put credentials in Git, VS Code
settings, OpenSpec artifacts, issue descriptions, or prompts.

### 3. Restore the locked skills

From the repository root:

```sh
npx skills experimental_install --frozen-lockfile
```

The installer restores local skill payloads for supported agent clients.
Generated directories such as `.agents/skills/`, `.claude/skills/`, and
`.codex/skills/` are ignored by Git. To inspect the lockfile without changing
it, use the skills CLI's dry-run option.

The current lockfile includes:

| Skill | Use it for |
| --- | --- |
| `golang-code-style` | Go naming, structure, and implementation style |
| `golang-testing` | Go unit, HTTP, integration, and coverage practices |
| `golang-security` | Go security review and threat modeling |
| `golang-lint` | `golangci-lint` configuration and linting |
| `security-best-practices` | Secure-by-default application patterns |
| `github-actions-hardening` | CI workflow permissions and supply-chain hardening |
| `taskfile-automation` | Taskfile design and automation workflows |

Skills provide guidance, not authority. Verify recommendations against the
repository and the active OpenSpec requirements.

### 4. Initialize OpenSpec

OpenSpec is intentionally not initialized in the template. Install it and
initialize only the integrations required by the new project:

```sh
npm install --global @fission-ai/openspec@latest
openspec --version
openspec init --tools claude,codex .
```

Use `--tools claude` or `--tools codex` for a single integration. Restart the
relevant agent client after initialization. Read
[`docs/openspec.md`](docs/openspec.md) for the complete lifecycle.

### 5. Validate the environment

```sh
task --list
task test
task ci
git diff --check
```

## Available capabilities

The starter currently provides:

- Provider-neutral text completion through Claude and OpenAI.
- Configuration-driven provider selection through environment variables.
- Context cancellation, a 60-second default HTTP timeout, bounded response
  bodies, and bounded retries for rate limits and server failures.
- A CLI that joins prompt arguments and handles SIGINT/SIGTERM cleanly.
- Local quality gates using tests, race detection, `go vet`, formatting,
  pinned `golangci-lint`, and pinned `govulncheck`.
- VS Code tasks and debugging for the example CLI.
- A dev container suitable for Go development and Node-based OpenSpec/skills
  tooling.
- Agent governance guidance covering allowlists, approvals, audit events, and
  fail-closed behavior.

The starter does **not** provide tool calling, delegation, persistence,
autonomous loops, database access, or external side effects. Add each new
capability through an OpenSpec change with observable scenarios, explicit
governance controls, approval boundaries, failure behavior, and focused tests.

## Daily development workflow

For non-trivial work:

1. Start or continue an OpenSpec change.
2. Explore the repository and record requirements and design decisions.
3. Implement only the tasks defined by the active change.
4. Run focused tests during development.
5. Run `task quality` and `task ci`.
6. Validate with `openspec validate <change-name>`.
7. Review the diff, then archive with `openspec archive <change-name>`.

Create a change with:

```sh
openspec new change <change-name>
```

Run the example CLI with:

```sh
PROMPT="Summarize the OpenSpec workflow in one sentence." task run
```

Run `task --list` to see every available command.

## Repository layout after initialization

```text
.
├── .devcontainer/          # Reproducible VS Code development container
├── .github/                # Copilot instructions and CI workflow
├── .vscode/                # Workspace settings, tasks, and debugging
├── cmd/agent/              # Example CLI
├── docs/                   # Agentic workflow and OpenSpec guidance
├── internal/agent/         # Provider configuration and adapters
├── openspec/                # Project changes/specs after local initialization
├── AGENTS.md               # Agent operating contract
├── CONTRIBUTING.md         # Contributor workflow
├── Taskfile.yml            # Development command surface
├── go.mod                  # Go module definition
└── skills-lock.json        # Skill source and revision lockfile
```

## Using the template as a new module

Update the module path before publishing or importing the new project:

```sh
go mod edit -module github.com/<owner>/<repository>
go mod tidy
```

Also update provider defaults, project documentation, and any OpenSpec
requirements that are specific to the derived application.

## Troubleshooting

**Skills are missing.** Run the frozen install command again from the
repository root and restart the agent client. Do not commit copied skill
directories.

**OpenSpec commands are missing.** Run `openspec --version`, initialize from
the repository root, and restart the relevant client. Do not copy generated
commands from another project.

**The CLI reports missing configuration.** Confirm `.env` contains a supported
`AGENT_PROVIDER`, a non-empty `AGENT_MODEL`, and the matching provider API key.

**A change needs external or destructive actions.** Pause and add an OpenSpec
capability with explicit approval, least privilege, audit behavior, and
fail-closed handling before implementing it.

## Contributing

See [`CONTRIBUTING.md`](CONTRIBUTING.md) for setup, testing, OpenSpec, pull
request, and security guidance.
