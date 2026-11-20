## Why

This empty folder needs a reproducible Go foundation for building agentic applications without coupling the project to one model vendor. Establishing provider support, OpenSpec workflows, and documented safety boundaries now makes future work testable and reviewable from the first change.

## What Changes

- Add a Go module with a small CLI entrypoint.
- Add provider-neutral completion support for Claude and OpenAI.
- Add environment-based configuration with no committed credentials.
- Add unit tests around both provider HTTP boundaries.
- Add repository guidance, quality commands, CI, and agent governance expectations.
- Initialize OpenSpec command integrations for Claude Code and Codex without committing skill payloads.
- Commit the npm `skills-lock.json` so consumers can restore selected skills reproducibly.

## Capabilities

### New Capabilities

- `agent-completion`: Send a prompt through the configured Claude or OpenAI provider.
- `spec-driven-workspace`: Provide repeatable OpenSpec and Go development workflows.

### Modified Capabilities

<!-- No existing capabilities are being modified. -->

## Impact

- New Go source under `cmd/agent` and `internal/agent`.
- New developer and agent guidance in `AGENTS.md`, `README.md`, and `docs/`.
- New CI workflow under `.github/workflows/ci.yml`.
- OpenSpec-generated Claude Code and Codex files under `.claude/` and `.codex/`.
- No runtime dependency is added beyond the Go standard library.
