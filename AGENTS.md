# Agentic Go Workspace

This repository is a spec-driven Go workspace for building small, testable agent systems.

## Development loop

1. If OpenSpec has not been initialized, follow `docs/openspec.md`; then start with `openspec new change <name>`.
2. Capture intent in `proposal.md`, behavior in `specs/`, decisions in `design.md`, and implementation work in `tasks.md`.
3. Implement only the tasks described by the active change.
4. Run `task quality` or, for the complete local CI workflow, `task ci`.
5. Validate the change with `openspec validate <name>` and archive it with `openspec archive <name>`.

OpenSpec command integrations are intentionally not included in the template. Follow `docs/openspec.md` to initialize Claude Code and/or Codex locally. Skill payloads are not committed; use the root `skills-lock.json` with `npx skills experimental_install --frozen-lockfile` to restore consumer-selected skills locally. Keep provider-specific behavior behind `internal/agent.Client`.

## Available project skills

The lockfile pins the project skills used by this template:

- Go code style, testing, security, and linting
- General secure-by-default guidance
- GitHub Actions hardening
- Taskfile workflow guidance

Restore them with `npx skills experimental_install --frozen-lockfile`. Treat installed skills as guidance, verify their recommendations against this repository, and never commit their generated payloads.

## Provider configuration

Copy `.env.example` to `.env` and set `AGENT_PROVIDER`, `AGENT_MODEL`, and the matching provider API key. Never commit `.env` or credentials. The sample CLI supports `claude` and `openai`.

## Go conventions

- Prefer concrete types and explicit error handling.
- Wrap errors with operation context using `%w`.
- Keep packages small and interfaces at the point of use.
- Use table-driven tests and `httptest` for provider boundaries.
- Do not add a provider SDK when the standard HTTP client is sufficient.
- Keep provider requests bounded by context, client timeouts, response-size limits, and retry policies.

## Agent safety

Agent tools must be allowlisted, fail closed, and auditable before they are exposed to production workflows. Keep destructive or external side effects behind explicit approval boundaries.

The current implementation supports text completion through Claude and OpenAI only. Tool calling, delegation, persistence, and autonomous side effects require a new OpenSpec capability with scenarios, governance controls, tests, and an audit design.
