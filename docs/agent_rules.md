# Agent Development Rules

These rules are non-negotiable for all AI agents (Antigravity, Gemini, Cursor, etc.) working on this repository.

## Linting & Quality
- `golangci-lint` **must** be run after any code changes are completed.
- All identified linting issues **must** be fixed before completing a task iteration.

## Branch Management
- The `main` branch is the primary development branch. All work is based off `main`.
- **Branch Creation**: When starting a new task, a new feature branch must be created from `main`.
- **Commits**: Commits to the feature branch must occur after each successful iteration/sub-task.
- **Conventional Commits**: All commit messages must follow the [Conventional Commits](https://www.conventionalcommits.org/) specification (e.g., `feat: ...`, `fix: ...`, `chore: ...`).

## Git Constraints
- **No Push**: Agents must **never** run `git push`.
- **No Merge**: Agents must **never** perform branch merges.
- Push and merge operations are strictly reserved for the human user.

## Workflow
1. Create feature branch from `main`.
2. Implement change/iteration.
3. Run tests and `golangci-lint`.
4. Fix issues.
5. Commit with conventional syntax.
6. Repeat until task is complete.
7. Notify user for review and manual merge.
