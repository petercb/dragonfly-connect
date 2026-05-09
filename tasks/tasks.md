# Implementation Tasks

## TASK-001: Project Onboarding & Structure
- **Status**: Done
- **Priority**: High
- **Description**: Setup core documentation and project structure.
- **Acceptance Criteria**:
    - `docs/` directory populated with PRD, technical, architecture, and testing docs.
    - `tasks/tasks.md` and `docs/status.md` created.

## TASK-002: Improve Error Handling & Environment Configuration
- **Status**: Planned
- **Priority**: High
- **PRD Reference**: Features & Requirements -> Configuration, Robust error handling.
- **Description**: Fix `srv.Listen()` error handling and implement `github.com/sethvargo/go-envconfig` for environment variable support.
- **Dependencies**: TASK-001
- **Checklist**:
    - [ ] Create feature branch `feature/task-002-error-handling-env-config` from `main`.
    - [ ] Install `github.com/sethvargo/go-envconfig`.
    - [ ] Update `Config` struct with `envconfig` tags.
    - [ ] Implement failing test for `loadConfig` environment variable overrides.
    - [ ] Update `loadConfig` to use `go-envconfig`.
    - [ ] Implement failing test for `srv.Listen()` error handling (if mockable) or verify logic.
    - [ ] Update `main.go` to handle `srv.Listen()` error.
    - [ ] Run `golangci-lint` and fix issues.
    - [ ] Verify all tests pass.
- **Acceptance Criteria**:
    - `srv.Listen()` errors are logged and the program exits gracefully.
    - Configuration can be overridden by environment variables (e.g., `SERVERS_JSON_PATH`, `LOG_LEVEL`).

## TASK-003: Formalize Agent Development Rules
- **Status**: Done
- **Priority**: High
- **Description**: Persist specific agent constraints (linting, branch naming, git push/merge restrictions) in repository documentation.
- **Acceptance Criteria**:
    - `docs/agent_rules.md` created.
    - `docs/technical.md` and `tasks/tasks.md` updated to reflect these rules.
