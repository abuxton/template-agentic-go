# Agentic Go Workspace

This repository is a spec-driven Go workspace for building small, testable agent systems.

## Development loop

1. Start with an OpenSpec change: `openspec new change <name>`.
2. Capture intent in `proposal.md`, behavior in `specs/`, decisions in `design.md`, and implementation work in `tasks.md`.
3. Implement only the tasks described by the active change.
4. Run `make test`, `make test-race`, and `make vet`.
5. Verify the change with `openspec verify <name>` and archive it with `openspec archive <name>`.

Claude Code and Codex command integrations are initialized under `.claude/` and `.codex/`. Skill payloads are not committed. Use the root `skills-lock.json` with `npx skills experimental_install` to restore consumer-selected skills locally. Keep provider-specific behavior behind `internal/agent.Client`.

## Provider configuration

Copy `.env.example` to `.env` and set `AGENT_PROVIDER`, `AGENT_MODEL`, and the matching provider API key. Never commit `.env` or credentials. The sample CLI supports `claude` and `openai`.

## Go conventions

- Prefer concrete types and explicit error handling.
- Wrap errors with operation context using `%w`.
- Keep packages small and interfaces at the point of use.
- Use table-driven tests and `httptest` for provider boundaries.
- Do not add a provider SDK when the standard HTTP client is sufficient.

## Agent safety

Agent tools must be allowlisted, fail closed, and auditable before they are exposed to production workflows. Keep destructive or external side effects behind explicit approval boundaries.
