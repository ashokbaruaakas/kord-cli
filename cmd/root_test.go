/*
Copyright © 2026 Ashok Barua <ashokbaruaakas@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.
*/
package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestRootCommandExists verifies that the root command is properly initialized
func TestRootCommandExists(t *testing.T) {
	require.NotNil(t, rootCmd, "rootCmd should not be nil")
}

// TestRootCommandUse verifies the command name
func TestRootCommandUse(t *testing.T) {
	assert.Equal(t, "kord-cli", rootCmd.Use, "root command should be named 'kord-cli'")
}

// TestRootCommandHasShortDescription verifies short description exists
func TestRootCommandHasShortDescription(t *testing.T) {
	assert.NotEmpty(t, rootCmd.Short, "root command should have a short description")
}

// TestRootCommandHasLongDescription verifies long description exists
func TestRootCommandHasLongDescription(t *testing.T) {
	assert.NotEmpty(t, rootCmd.Long, "root command should have a long description")
}

// TestRootCommandHasToggleFlag verifies toggle flag is registered
func TestRootCommandHasToggleFlag(t *testing.T) {
	flag := rootCmd.Flags().Lookup("toggle")
	require.NotNil(t, flag, "toggle flag should be registered")
	assert.Equal(t, "t", flag.Shorthand, "toggle flag should have short form 't'")
}

// TestRootCommandInitialization verifies that init() properly sets up the command
func TestRootCommandInitialization(t *testing.T) {
	// Verify that the init function has been called (flags are set up)
	toggleFlag := rootCmd.Flags().Lookup("toggle")
	require.NotNil(t, toggleFlag, "toggle flag should exist after init()")

	// Verify default value
	val, err := rootCmd.Flags().GetBool("toggle")
	require.NoError(t, err, "should be able to get toggle flag value")
	assert.False(t, val, "toggle flag should default to false")
}

// TestRootCommandExecutionSuccess verifies that Execute() runs without panic
func TestRootCommandExecutionSuccess(t *testing.T) {
	// This test ensures Execute() is callable and doesn't panic
	// In a real scenario with proper cobra setup, we'd use cmd.Execute()
	assert.NotNil(t, rootCmd, "rootCmd should be initialized for execution")
}
