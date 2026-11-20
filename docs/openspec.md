# OpenSpec setup

This template does not include generated OpenSpec command files, skills, active changes, or archived changes. Initialize OpenSpec locally when you begin using the template.

## Initialize OpenSpec

From the repository root, configure the integrations you use:

```sh
openspec init --tools claude,codex .
```

Use `--tools claude` or `--tools codex` when only one integration is required. Restart the relevant IDE or agent CLI after initialization so its commands become available.

OpenSpec creates local command and skill files for the selected integrations. Skill payloads are not source files for this template and remain ignored by Git. The project `skills-lock.json` can be used to restore deliberately selected npm-managed skills:

```sh
npx skills experimental_install --frozen-lockfile
```

## Start a change

Create a change before making non-trivial behavior changes:

```sh
openspec new change <change-name>
```

Use the generated workflow to create:

1. `proposal.md` for motivation and scope
2. `specs/<capability>/spec.md` for testable requirements
3. `design.md` for technical decisions
4. `tasks.md` for implementation steps

When the artifacts are complete, implement the tasks and validate the change:

```sh
openspec status --change <change-name> --json
openspec validate <change-name>
openspec archive <change-name>
```

If the selected agent integration provides slash commands, the equivalent workflow is typically `/opsx:new`, `/opsx:continue`, `/opsx:apply`, `/opsx:verify`, and `/opsx:archive`.

## Current capabilities

The template's initial capability is provider-neutral text completion through Claude and OpenAI. It does not include tool calling, delegation, persistence, autonomous loops, or external side effects.

Add any of those capabilities through a new OpenSpec change. The change should specify observable scenarios, provider and tool boundaries, governance policy, approval requirements, audit events, failure behavior, and focused tests.

## Repository expectations

- Keep OpenSpec changes focused and observable.
- Use `WHEN` / `THEN` / `AND` scenarios for requirements.
- Run `task quality` and `task ci` before archiving.
- Do not commit credentials, generated skill payloads, or unrelated local agent configuration.
