# Template Agentic Go

A spec-driven Go starter for agentic applications with Claude and OpenAI support. It includes OpenSpec workflows, Claude Code and Codex integrations, a provider-neutral HTTP client, and a small CLI.

## Quick start

```sh
cp .env.example .env
# Set AGENT_PROVIDER, AGENT_MODEL, and the matching API key in .env
make test
make build
AGENT_PROVIDER=claude AGENT_MODEL=<model> ANTHROPIC_API_KEY=<key> \
  PROMPT="Summarize the OpenSpec workflow in one sentence." make run
```

For OpenAI, use `AGENT_PROVIDER=openai` and `OPENAI_API_KEY`.

## Spec-driven workflow

OpenSpec command wiring is initialized for both Claude Code and Codex. Skill payloads are intentionally not committed; consumers restore any selected project skills with the npm `skills` CLI:

```text
/opsx:new <change-name>
/opsx:continue <change-name>
/opsx:apply <change-name>
/opsx:verify <change-name>
/opsx:archive <change-name>
```

The reusable lockfile is `skills-lock.json`. Add project skills with `npx skills add <package>`; the CLI updates the lockfile, and consumers can restore the selected skills with `npx skills experimental_install`. Installed skill files remain local and are ignored by Git.

Changes live under `openspec/changes/` and provide the durable record of why and how work was done. Use `AGENTS.md` for repository conventions and `.env.example` for runtime configuration.

## Project layout

| Path | Purpose |
| --- | --- |
| `cmd/agent` | Example command-line entrypoint |
| `internal/agent` | Provider configuration and Claude/OpenAI adapters |
| `openspec` | Specifications and change history |
| `.claude` | Claude Code OpenSpec commands and skills |
| `.codex` | Codex OpenSpec skills |

## Quality gates

```sh
make test
make test-race
make vet
```

`make lint` is available when `golangci-lint` is installed.
