## ADDED Requirements

### Requirement: Provider-neutral completion

The CLI SHALL send a non-empty prompt to the configured provider and return the provider's text response.

#### Scenario: Claude completion

- **WHEN** `AGENT_PROVIDER=claude` and valid Claude configuration is present
- **THEN** the client sends a request to `/v1/messages` with the configured model and API key
- **AND** returns the first response text

#### Scenario: OpenAI completion

- **WHEN** `AGENT_PROVIDER=openai` and valid OpenAI configuration is present
- **THEN** the client sends a request to `/v1/chat/completions` with the configured model and bearer token
- **AND** returns the first assistant message

#### Scenario: Invalid configuration

- **WHEN** the provider, API key, model, or prompt is missing or unsupported
- **THEN** the client returns an explicit error
- **AND** it does not make a provider request

### Requirement: Provider failures are visible

The client SHALL return non-success HTTP responses and malformed provider responses as errors containing operation context.

#### Scenario: Provider API failure

- **WHEN** a provider returns a non-success HTTP status
- **THEN** the client returns an error containing the HTTP status

#### Scenario: Empty provider response

- **WHEN** a provider response contains no usable text
- **THEN** the client returns an error rather than a success-shaped empty result
