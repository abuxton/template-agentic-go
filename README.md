# Template Agentic Go

A spec-driven Go starter for agentic applications with Claude and OpenAI support. It includes OpenSpec setup guidance, a provider-neutral HTTP client, and a small CLI.

## Quick start

```sh
cp .env.example .env
# Set AGENT_PROVIDER, AGENT_MODEL, and the matching API key in .env
task test
task build
PROMPT="Summarize the OpenSpec workflow in one sentence." task run
```

For OpenAI, use `AGENT_PROVIDER=openai` and `OPENAI_API_KEY`.

## OpenSpec workflow

Generated OpenSpec files are intentionally not included in this template. Follow [`docs/openspec.md`](docs/openspec.md) to initialize Claude Code and/or Codex integrations locally. Changes live under `openspec/changes/` after initialization and provide the durable record of why and how work was done.

The reusable skill lockfile is `skills-lock.json`. Add project skills with `npx skills add <package>`; the CLI updates the lockfile, and consumers can restore selected skills with `npx skills experimental_install`.

## Project layout

| Path | Purpose |
| --- | --- |
| `cmd/agent` | Example command-line entrypoint |
| `internal/agent` | Provider configuration and Claude/OpenAI adapters |
| `docs/openspec.md` | OpenSpec initialization and usage |
| `skills-lock.json` | npm-managed project skill restore contract |

## Quality gates

```sh
task quality
task ci
```

The lint and vulnerability checks use pinned tool versions through `go run`, so no global tool installation is required. Run `task --list` to discover all commands.

## Using this as a template

After creating a repository from this template, update the module path before publishing or importing it:

```sh
go mod edit -module github.com/<owner>/<repository>
go mod tidy
```
