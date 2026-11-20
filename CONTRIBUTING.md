# Contributing

Thank you for contributing to this spec-driven Go template. Contributions should be small, testable, and documented well enough for another contributor to reproduce the reasoning.

## Before you start

Install:

- Go 1.27 or newer
- [Task](https://taskfile.dev/)
- Git
- OpenSpec

Clone the repository and inspect the available development commands:

```sh
git clone <repository-url>
cd template-agentic-go
task --list
```

Do not commit API keys, `.env` files, generated skill directories, or other credentials. Project skills are restored locally from `skills-lock.json` with:

```sh
npx skills experimental_install --frozen-lockfile
```

The lockfile includes Go style, testing, security, linting, GitHub Actions hardening, and Taskfile workflow skills. Generated skill payloads are local-only and must not be committed.

## Development workflow

1. Create a branch from the default branch.
2. For non-trivial work, create an OpenSpec change:

   ```sh
   openspec new change <change-name>
   ```

3. If OpenSpec is not initialized, follow [`docs/openspec.md`](docs/openspec.md), then define the proposal, requirements, design, and tasks before implementation.
4. Implement only the agreed tasks and update their checkboxes as work is completed.
5. Run the quality gates:

   ```sh
   task quality
   task ci
   ```

6. Validate the OpenSpec change:

   ```sh
   openspec validate <change-name>
   ```

7. Archive completed changes with `openspec archive <change-name>`.

Small documentation or maintenance changes may not need a full OpenSpec change, but they should still be scoped clearly and checked locally.

## Go guidelines

- Prefer the standard library unless a dependency provides clear value.
- Keep provider-specific behavior behind `internal/agent.Client`.
- Return explicit, contextual errors; do not silently fall back.
- Use concrete types, table-driven tests, and `httptest` for HTTP boundaries.
- Keep provider requests bounded by context, client timeouts, response-size limits, and retry policies.
- Keep agent tools allowlisted and fail closed before adding external or destructive side effects.

## Pull requests

Before opening a pull request:

- Keep the branch focused on one logical change.
- Update documentation and specifications when behavior changes.
- Add or update focused tests for new behavior.
- Confirm `task quality` and `task ci` pass.
- Review the diff for secrets, generated files, unrelated edits, and accidental skill payloads.
- Use a clear, imperative pull request title and describe the motivation, implementation, and validation.

Pull requests should be ready for review, or clearly identify any remaining work or known limitations.

## Commit messages

Use concise [Conventional Commit](https://www.conventionalcommits.org/) messages, for example:

```text
feat: add provider retry policy
fix: reject empty agent prompts
docs: clarify OpenSpec workflow
```

## Security issues

Do not report suspected vulnerabilities in a public issue. Follow the repository's configured GitHub security reporting process, or contact the maintainers privately if no process is configured.
