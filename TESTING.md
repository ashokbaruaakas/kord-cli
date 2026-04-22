# Testing Guide for kord-cli

This guide covers testing setup, conventions, and TDD workflow for the kord-cli project.

## Overview

This project uses **TDD (Test-Driven Development)** with the following testing framework:
- **Testing Framework**: Go's built-in `testing` package
- **Assertion Library**: [testify](https://github.com/stretchr/testify) for expressive assertions
- **Test Structure**: Unit tests + Feature/Integration tests

## Test Organization

```
kord-cli/
├── cmd/
│   ├── root.go
│   └── root_test.go              # Unit tests for root command
│   ├── <command>.go              # Subcommands (to be added)
│   └── <command>_test.go         # Unit tests for subcommands
├── tests/
│   ├── integration_test.go       # Integration/feature tests
│   └── helpers_test.go           # Test utilities and helpers
└── TESTING.md
```

## Running Tests

### Run all tests
```bash
go test ./...
```

### Run tests with verbose output
```bash
go test ./... -v
```

### Run tests with coverage
```bash
go test ./... -cover
```

### Generate coverage report
```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Run specific test file
```bash
go test ./cmd -v
```

### Run specific test function
```bash
go test ./cmd -run TestRootCommandUse -v
```

### Run tests in watch mode (requires external tool)
```bash
# Using watchexec (install with: brew install watchexec)
watchexec -c -- go test ./... -v

# Optional: only trigger on Go file changes
watchexec -c -e go -- go test ./... -v
```

## TDD Workflow

Follow these steps when implementing a new command or feature:

### 1. Write the Test First
Create a `*_test.go` file with failing tests before implementation:

```go
package cmd

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestSetupCommandExists(t *testing.T) {
    require.NotNil(t, setupCmd, "setupCmd should not be nil")
    assert.Equal(t, "setup", setupCmd.Use)
}

func TestSetupCommandDescription(t *testing.T) {
    assert.NotEmpty(t, setupCmd.Short)
    assert.NotEmpty(t, setupCmd.Long)
}
```

### 2. Run Tests (They Should Fail)
```bash
go test ./cmd -v
```

### 3. Implement the Minimal Feature
Create `setup.go` with minimal implementation to pass the test:

```go
package cmd

import "github.com/spf13/cobra"

var setupCmd = &cobra.Command{
    Use:   "setup",
    Short: "Configure kord with your Git, task platform, and LLM settings",
    Long:  "A longer description for the setup command...",
    RunE: func(cmd *cobra.Command, args []string) error {
        // Implementation
        return nil
    },
}

func init() {
    rootCmd.AddCommand(setupCmd)
}
```

### 4. Run Tests Again (They Should Pass)
```bash
go test ./cmd -v
```

### 5. Refactor if Needed
Improve code quality while keeping tests passing.

### 6. Repeat for Next Feature
Write tests for the next piece of functionality.

## Test Types

### Unit Tests
Test individual components in isolation.

**Location**: `cmd/<command>_test.go`

**Example**:
```go
func TestRootCommandUse(t *testing.T) {
    assert.Equal(t, "kord-cli", rootCmd.Use)
}
```

### Integration/Feature Tests
Test interactions between components and workflows.

**Location**: `tests/integration_test.go`

**Example**:
```go
func TestCLISubcommandExecution(t *testing.T) {
    rootCmd := &cobra.Command{Use: "kord-cli"}
    var ran bool

    setupCmd := &cobra.Command{
        Use: "setup",
        RunE: func(cmd *cobra.Command, args []string) error {
            ran = true
            return nil
        },
    }

    rootCmd.AddCommand(setupCmd)
    rootCmd.SetArgs([]string{"setup"})
    rootCmd.Execute()

    assert.True(t, ran, "setup command should execute")
}
```

## Testify Assertions Cheat Sheet

### Require (Stops test on failure)
```go
require.NotNil(t, value)
require.NoError(t, err)
require.Equal(t, expected, actual)
```

### Assert (Continues test on failure)
```go
assert.NotNil(t, value)
assert.NoError(t, err)
assert.Equal(t, expected, actual)
assert.NotEqual(t, unexpected, actual)
assert.True(t, condition)
assert.False(t, condition)
assert.Contains(t, str, substring)
assert.Len(t, slice, 5)
assert.Error(t, err)
```

### Message Formatting
```go
assert.Equal(t, expected, actual, "custom failure message")
```

## Best Practices

1. **Test Names**: Use descriptive names starting with `Test`
   ```go
   TestRootCommandExists()       // ✓ Good
   TestSetup()                   // ✗ Too vague
   ```

2. **One Assertion Per Test (When Possible)**
   ```go
   // ✓ Good: Focused test
   func TestCommandUse(t *testing.T) {
       assert.Equal(t, "setup", setupCmd.Use)
   }

   // ✗ Avoid: Multiple assertions
   func TestCommand(t *testing.T) {
       assert.Equal(t, "setup", setupCmd.Use)
       assert.NotEmpty(t, setupCmd.Short)
       assert.NotEmpty(t, setupCmd.Long)
   }
   ```

3. **Use Require for Critical Failures**
   ```go
   require.NotNil(t, setupCmd)           // Must not be nil
   assert.Equal(t, "setup", setupCmd.Use) // Can continue if fails
   ```

4. **Test Error Paths**
   ```go
   func TestCommandWithInvalidArgs(t *testing.T) {
       cmd := &cobra.Command{
           Use: "test",
           RunE: func(cmd *cobra.Command, args []string) error {
               return errors.New("invalid argument")
           },
       }

       cmd.SetArgs([]string{"invalid"})
       err := cmd.Execute()

       assert.Error(t, err, "should return error for invalid args")
   }
   ```

5. **Document Complex Tests**
   ```go
   // TestStartCommandWorkflow verifies the complete workflow:
   // 1. Fetch task info from Notion
   // 2. Generate branch name
   // 3. Create Git branch
   // 4. Update task status
   func TestStartCommandWorkflow(t *testing.T) {
       // Test implementation
   }
   ```

## Coverage Goals

- **Unit Tests**: Aim for >80% coverage
- **Integration Tests**: Focus on critical workflows (>60% coverage is acceptable)
- **Overall**: Target >70% total project coverage

Check coverage for specific packages:
```bash
go test ./cmd -cover
go test ./tests -cover
```

## Common Test Patterns

### Testing Flag Parsing
```go
func TestCommandFlags(t *testing.T) {
    cmd := &cobra.Command{Use: "test"}
    cmd.Flags().StringP("config", "c", "", "config file")

    cmd.SetArgs([]string{"--config", "test.json"})
    cmd.Execute()

    config, _ := cmd.Flags().GetString("config")
    assert.Equal(t, "test.json", config)
}
```

### Testing Command Execution
```go
func TestCommandExecution(t *testing.T) {
    var executed bool

    cmd := &cobra.Command{
        Use: "test",
        RunE: func(cmd *cobra.Command, args []string) error {
            executed = true
            return nil
        },
    }

    cmd.SetArgs([]string{})
    err := cmd.Execute()

    require.NoError(t, err)
    assert.True(t, executed)
}
```

### Testing with Mock Data
```go
func TestCommandWithMockData(t *testing.T) {
    mockData := map[string]string{
        "task_id": "123",
        "title":   "Test Task",
    }

    cmd := &cobra.Command{
        Use: "test",
        RunE: func(cmd *cobra.Command, args []string) error {
            assert.Equal(t, "123", mockData["task_id"])
            return nil
        },
    }

    cmd.SetArgs([]string{})
    err := cmd.Execute()
    require.NoError(t, err)
}
```

## Continuous Testing

To continuously run tests as you develop:

```bash
# Install entr (macOS: brew install entr)
find . -name "*.go" | entr go test ./... -v

# Or use watchexec
watchexec -c -- go test ./... -v
```

## Debugging Tests

### Print Debug Information
```go
t.Logf("Command use: %s", setupCmd.Use)
t.Logf("Args: %v", args)
```

### Run Single Test with Extra Verbosity
```bash
go test ./cmd -run TestRootCommandUse -v
```

### Set Build Tags for Test Variants
```go
//go:build testing
// +build testing

func helperFunction() string {
    return "test value"
}
```

## Next Steps

1. ✅ Unit test framework set up (cmd/*_test.go)
2. ✅ Integration test framework set up (tests/)
3. ✅ Testify assertions configured
4. **TODO**: Add tests for each new command as you implement them
5. **TODO**: Add tests for internal packages (git/, taskmanager/, llm/, config/)

## Resources

- [Go Testing Package Docs](https://pkg.go.dev/testing)
- [Testify Documentation](https://github.com/stretchr/testify)
- [Cobra Testing Patterns](https://cobra.dev/)
- [TDD Best Practices](https://en.wikipedia.org/wiki/Test-driven_development)
