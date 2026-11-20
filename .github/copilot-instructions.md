# Project instructions

Use OpenSpec for non-trivial changes. If it is not initialized, follow `docs/openspec.md`. Read the active proposal, specs, design, and tasks before editing code, and keep implementation aligned with the scenarios.

Restore the pinned project skills with `npx skills experimental_install --frozen-lockfile`. The lockfile covers Go style, testing, security, linting, GitHub Actions hardening, and Taskfile workflows. Generated skill payloads are local-only and must not be committed.

For Go changes, prefer standard-library solutions, explicit error handling, concrete types, and `httptest` for external API boundaries. Run `task quality` and `task ci` before completing work.

The current agent capability is provider-neutral text completion for Claude and OpenAI. New tools, delegation, external side effects, persistence, or autonomous loops require an OpenSpec capability with explicit scenarios, fail-closed governance, approval boundaries, and audit behavior.
