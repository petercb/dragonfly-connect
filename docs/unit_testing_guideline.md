# Unit Testing Guideline

## Testing Strategy
- Use Go's built-in `testing` package.
- Focus on unit testing configuration parsing and form generation logic.
- Mock network interactions where possible.

## Test Commands
- Run all tests: `go test -race ./...`
- Run tests with coverage: `go test -race -cover ./...`

## TDD Workflow
1. Write a failing test for the desired behavior.
2. Run the test to confirm failure.
3. Write the minimal code to make the test pass.
4. Refactor and ensure all tests still pass.
