# AGENTS.md

## Project overview

This repository is a reusable, spec-driven Go starter for small agentic
applications. It contains:

- `cmd/agent`: the example CLI entrypoint.
- `internal/agent`: provider configuration and standard-library HTTP adapters
  for Claude and OpenAI.
- `docs/`: OpenSpec initialization and agent workflow guidance.
- `Taskfile.yml`: the local development and CI command surface.

The current implementation performs provider-neutral text completion only.
Tool calling, delegation, persistence, and autonomous side effects are not
implemented capabilities.

## Development environment

The preferred environment is the checked-in VS Code dev container:

1. Open the repository in VS Code.
2. Select **Reopen in Container**.
3. Run `task --list`.

The container provides Go 1.27, Task, and Node/npm for OpenSpec and skills.
The checked-in `.vscode/` configuration recommends Go, Task, GitHub Copilot,
Claude Code, and OpenAI extensions. Credentials must still be supplied through
the normal local or extension authentication flows.

For a host setup, install Go 1.27 or newer, Task, Node.js/npm, Git, and
OpenSpec. Copy the environment template before running the CLI:

```sh
cp .env.example .env
```

Never commit `.env`, API keys, or generated skill payloads.

## Agent and specification workflow

Use OpenSpec for non-trivial behavior changes:

1. If OpenSpec is not initialized, follow `docs/openspec.md`.
2. Create a change with `openspec new change <change-name>`.
3. Read the active proposal, requirements, design, and tasks before editing.
4. Implement only the agreed tasks and update task checkboxes.
5. Validate with `openspec validate <change-name>`.
6. Archive completed work with `openspec archive <change-name>`.

OpenSpec command integrations and generated specifications are intentionally
not part of this template. Consumers initialize them locally. Restore the
locked project skills with:

```sh
npx skills experimental_install --frozen-lockfile
```

Only `skills-lock.json` belongs in version control. Do not commit `.agents/skills`,
`.claude/skills`, or `.codex/skills` payloads.

## Common commands

Run `task --list` for the complete list. The primary commands are:

```sh
task test       # go test ./...
task test-race  # race-enabled tests
task vet        # go vet ./...
task format     # fail if files need gofmt
task lint       # pinned golangci-lint
task vuln       # pinned govulncheck
task quality    # tests, race, vet, format, lint, vulnerability scan
task build      # writes bin/agent
task ci         # quality checks followed by build
```

Run the example CLI with a prompt:

```sh
PROMPT="Summarize this project." task run
```

`task run` loads `.env` and requires `AGENT_PROVIDER`, `AGENT_MODEL`, and the
matching `ANTHROPIC_API_KEY` or `OPENAI_API_KEY`.

## Testing and validation

Tests live beside the package they exercise, for example
`internal/agent/client_test.go`. Use table-driven tests and `httptest` for
provider HTTP boundaries. Add focused tests for new behavior and preserve
context cancellation, response limits, provider error handling, and retry
behavior.

Before completing a code change, run:

```sh
task quality
task ci
```

For documentation-only changes, run at least `git diff --check`. CI runs
`task ci` on pushes and pull requests with read-only repository permissions.

## Code style and architecture

- Use standard-library Go unless a dependency provides clear value.
- Keep provider-specific behavior behind `internal/agent.Client`.
- Prefer concrete types, explicit validation, and contextual `%w` error
  wrapping.
- Use Go naming conventions: `ID`, `URL`, and `HTTP`, not `Id`, `Url`, or
  `Http`.
- Keep interfaces at the point of use; inject HTTP clients for testability.
- Run `gofmt`; `task format` enforces formatting.
- Bound external calls with context cancellation, client timeouts, response
  size limits, and bounded retries.
- Do not silently fall back after invalid configuration or provider errors.

The module path is `github.com/abuxton/template-agentic-go`. Template users
should update it with `go mod edit -module ...` and then run `go mod tidy`.

## Security and governance

- Keep credentials in `.env` or the provider's supported secret store.
- Do not log API keys, authorization headers, prompts containing secrets, or
  complete provider error bodies.
- Validate provider URLs and configuration before making requests.
- Keep external or destructive agent actions allowlisted, fail-closed, and
  auditable behind explicit approval boundaries.
- Any new agent capability needs an OpenSpec change with scenarios,
  governance controls, tests, and an audit design.
- Review changes for accidental secrets, generated files, and skill payloads.

## Pull requests and commits

Keep changes focused and update documentation and OpenSpec artifacts when
behavior changes. Before opening a pull request, run `task ci`, validate the
active OpenSpec change, and inspect the complete diff.

Use Conventional Commits with an imperative subject, for example:

```text
feat: add provider retry policy
fix: reject empty agent prompts
docs: clarify OpenSpec workflow
```
