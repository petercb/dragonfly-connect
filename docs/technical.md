# Technical Specification: Dragonfly Bedrock Connect

## Core Principles

- **SOLID**: Maintain single responsibility for config, forms, and server logic.
- **KISS**: Keep the proxy logic simple; avoid unnecessary complexity.
- **DRY**: Reuse form generation and config parsing logic.
- **Agent Guidelines**: All AI agents must strictly follow [docs/agent_rules.md](file:///Users/pburns/git/dragonfly-connect/docs/agent_rules.md).

## Technology Stack

- **Language**: Go
- **Server Framework**: [Dragonfly](https://github.com/df-mc/dragonfly)
- **Configuration**: `github.com/sethvargo/go-envconfig` for environment variables.
- **Logging**: `log/slog` for structured logging.

## Directory Structure Convention

- Code uses the standard Go project layout (e.g., root directory for simple
  projects, or `cmd/` and `internal/` for larger applications).
- Tests reside alongside the code they test (e.g., `config_test.go`).

## Design Patterns

- **Configuration Pattern**: Layered configuration (Default -> JSON -> Environment).
- **Command/Form Pattern**: Separation of form definition from submission logic.
