# Agentic workflow

## Change lifecycle

Every non-trivial behavior change should use an OpenSpec change:

```text
explore → proposal → specs → design → tasks → apply → verify → archive
```

Requirements should describe observable behavior with `WHEN`, `THEN`, and `AND` scenarios. Designs should record provider, tool, data, and failure-boundary decisions. Tasks should be small enough to verify independently.

## Provider boundary

`internal/agent.Client` presents one `Complete` operation while keeping Claude and OpenAI wire formats private. This makes provider selection configuration-driven and keeps tests local with `httptest`.

## Governance baseline

Before adding tools, define:

- an explicit tool allowlist;
- blocked content and dangerous-intent checks;
- rate limits and approval requirements for external or destructive actions;
- append-only audit events;
- fail-closed behavior when policy evaluation fails.

The current CLI intentionally performs text completion only. Tool execution, persistence, delegation, and autonomous side effects are out of scope until they have an OpenSpec change with corresponding tests and governance controls.
