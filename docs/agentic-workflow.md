# Agentic workflow

## Change lifecycle

Every non-trivial behavior change should use an OpenSpec change:

```text
explore → proposal → specs → design → tasks → apply → verify → archive
```

Requirements should describe observable behavior with `WHEN`, `THEN`, and `AND` scenarios. Designs should record provider, tool, data, and failure-boundary decisions. Tasks should be small enough to verify independently.

## Provider boundary

`internal/agent.Client` presents one `Complete` operation while keeping Claude and OpenAI wire formats private. This makes provider selection configuration-driven and keeps tests local with `httptest`.

Provider calls use caller cancellation, a bounded HTTP client timeout, capped response bodies, and bounded retries for rate-limit and server failures. New provider adapters should preserve those boundaries and add focused tests for their wire format and failure modes.

## Governance baseline

Before adding tools, define:

- an explicit tool allowlist;
- blocked content and dangerous-intent checks;
- rate limits and approval requirements for external or destructive actions;
- append-only audit events;
- fail-closed behavior when policy evaluation fails.

The current CLI intentionally performs text completion only. Tool execution, persistence, delegation, and autonomous side effects are out of scope until they have an OpenSpec change with corresponding tests and governance controls.

## Verification

Use `task quality` for the fast quality workflow and `task ci` to include the build. Both use pinned versions of `golangci-lint` and `govulncheck` through the Taskfile.

## Project skills

The root `skills-lock.json` pins reusable guidance for Go style, testing, security, linting, GitHub Actions hardening, and Taskfile workflows. Restore it with `npx skills experimental_install --frozen-lockfile`; generated skill directories remain local and ignored.
