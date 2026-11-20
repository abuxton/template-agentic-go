## Context

The repository starts empty. It needs a low-dependency Go baseline that can exercise Claude and OpenAI while remaining easy to test offline and easy to extend through spec-driven changes.

## Goals / Non-Goals

**Goals:**

- Provide one completion interface with Claude and OpenAI adapters.
- Keep provider credentials and endpoints configurable through environment variables.
- Use the standard library so setup works without downloading provider SDKs.
- Establish OpenSpec, agent guidance, tests, and CI as part of the baseline.

**Non-Goals:**

- Tool calling, autonomous loops, persistence, delegation, streaming, or production hosting.
- Secret management beyond documenting environment-based local configuration.
- Selecting a specific model on behalf of the user.

## Decisions

### Standard HTTP client over provider SDKs

The provider boundary uses `net/http` and private wire-format structs. This avoids an unnecessary dependency for a small starter, keeps tests deterministic with `httptest`, and leaves room to adopt SDKs later behind the same interface.

### Configuration-driven provider selection

`AGENT_PROVIDER`, `AGENT_MODEL`, and provider-specific API variables select the runtime behavior. Invalid configuration fails before any network request, making local failures explicit and preventing silent fallback between vendors.

### OpenSpec as the change contract

The initial setup is itself captured as an OpenSpec change. Future feature work should add observable scenarios before implementation, and agent tooling must not introduce external side effects without a reviewed specification.

## Risks / Trade-offs

- [Provider APIs evolve] -> Keep wire-format code isolated and cover headers, paths, and response parsing with local HTTP tests.
- [Environment variables can be misconfigured] -> Validate provider, model, and API key before requests and return contextual errors.
- [Future tools could create side effects] -> Require an OpenSpec change plus allowlists, audit events, and approval gates before adding them.

## Migration Plan

No migration is required for a new repository. Copy `.env.example`, configure one provider, run the quality gates, and then create a new OpenSpec change for application-specific behavior.

## Open Questions

- Which model defaults, if any, should be selected for a deployed application?
- Which tool runtime and audit store should be adopted when tool calling is introduced?

## Skills distribution

OpenSpec command files are repository configuration, but generated skill payloads are consumer-local installation artifacts. The repository commits the npm-managed `skills-lock.json` with an empty default set and ignores `.agents/skills/`, `.claude/skills/`, and `.codex/skills/`. Consumers can add and restore skills without changing the template's committed source.
