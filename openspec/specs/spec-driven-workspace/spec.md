# spec-driven-workspace Specification

## Purpose
TBD - created by archiving change setup-agentic-go. Update Purpose after archive.
## Requirements
### Requirement: OpenSpec-first development

The repository SHALL document an OpenSpec lifecycle for non-trivial changes.

#### Scenario: Starting a change

- **WHEN** a developer begins a non-trivial change
- **THEN** the repository guidance directs them to create an OpenSpec change with proposal, specs, design, and tasks

#### Scenario: Verifying a change

- **WHEN** implementation work is complete
- **THEN** the repository guidance directs the developer to run Go quality gates and OpenSpec verification before archiving

### Requirement: Safe agent configuration

The repository SHALL keep credentials out of source control and document explicit boundaries for future tools and side effects.

#### Scenario: Local credentials

- **WHEN** a developer configures a provider locally
- **THEN** they use environment variables from `.env.example`
- **AND** `.gitignore` excludes local environment files

#### Scenario: Adding agent tools

- **WHEN** a future change adds agent tools or side effects
- **THEN** repository guidance requires allowlisting, fail-closed behavior, auditability, and approval boundaries

### Requirement: Reproducible local skills

The repository SHALL commit the npm skills lockfile while keeping installed skill payloads out of source control.

#### Scenario: Restoring skills

- **WHEN** a consumer runs `npx skills experimental_install`
- **THEN** the CLI uses `skills-lock.json` as the source of truth
- **AND** installed skill directories remain local and ignored by Git

#### Scenario: Empty default lock

- **WHEN** no project skills have been selected
- **THEN** `skills-lock.json` contains an empty skills map
- **AND** the repository contains no committed skill payloads

