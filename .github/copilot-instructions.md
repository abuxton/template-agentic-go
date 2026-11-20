# Project instructions

Use OpenSpec for non-trivial changes. Read the active proposal, specs, design, and tasks before editing code, and keep implementation aligned with the scenarios.

For Go changes, prefer standard-library solutions, explicit error handling, concrete types, and `httptest` for external API boundaries. Run `make test`, `make test-race`, and `make vet` before completing work.

Agent features must fail closed, keep credentials out of source control, and require explicit specifications for tools, delegation, external side effects, and audit behavior.
