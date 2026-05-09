# Agent Development Rules

These rules are non-negotiable for all AI agents (Antigravity, Gemini, Cursor,
etc.) working on this repository.

## Validation & Formatting Rules

Agents must run the following tools based on the type of files modified:

### 1. General (All Text Files)

- **Tool**: `codespell`
- **Action**: Run on all modified text files.
- **Ignores**: Add to `.codespellrc` for false positives.

### 2. Go Code

- **Tool**: `golangci-lint`
- **Action**: Run **only** after Go code changes. Fix all issues.

### 3. YAML Files

- **Tools**: `yamlfmt` (formatting) and `yamllint` (linting).
- **Action**: Run on any modified `.yml` or `.yaml` files. Correct all issues.

### 4. CircleCI Configuration

- **Tool**: `circleci config validate`
- **Action**: Run after any changes to `.circleci/config.yml`. Fix all issues.

### 5. Goreleaser Configuration

- **Tool**: `goreleaser check`
- **Action**: Run after any changes to `.goreleaser.yaml`. Fix all issues.

### 6. Markdown Files

- **Tool**: `markdownlint-cli2`
- **Action**: Run on any modified `.md` files. Correct all issues.

## Branch Management

- The `main` branch is the primary development branch. All work is based off `main`.
- **Branch Creation**: When starting a new task, a new feature branch must be
  created from `main`.
- **Commits**: Commits to the feature branch must occur after each successful iteration/sub-task.
- **Conventional Commits**: All commit messages must follow the
  [Conventional Commits](https://www.conventionalcommits.org/) specification
  (e.g., `feat: ...`, `fix: ...`, `chore: ...`).

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
