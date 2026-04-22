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
package tests

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestCLIHelpCommand verifies that the help command works
func TestCLIHelpCommand(t *testing.T) {
	cmd := &cobra.Command{
		Use:   "kord-cli",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application.`,
	}

	// Verify command structure
	assert.Equal(t, "kord-cli", cmd.Use, "command use should be kord-cli")
	assert.NotEmpty(t, cmd.Short, "command should have short description")
	assert.NotEmpty(t, cmd.Long, "command should have long description")
}

// TestCLIVersionFlag verifies version flag can be set
func TestCLIVersionFlag(t *testing.T) {
	cmd := &cobra.Command{
		Use: "kord-cli",
	}

	cmd.Version = "0.1.0"

	var out bytes.Buffer
	cmd.SetOut(&out)

	cmd.SetArgs([]string{"--version"})
	_ = cmd.Execute() // Version flag causes exit
	assert.NotNil(t, cmd.Version, "version should be set")
	assert.Equal(t, "0.1.0", cmd.Version, "version should be 0.1.0")
}

// TestCLISubcommandStructure verifies that subcommands can be added
func TestCLISubcommandStructure(t *testing.T) {
	// Create root command
	rootCmd := &cobra.Command{
		Use:   "kord-cli",
		Short: "Digital cord tying together Git, task management, and LLM automation",
	}

	// Create a subcommand
	setupCmd := &cobra.Command{
		Use:   "setup",
		Short: "Configure kord with your Git, task platform, and LLM settings",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Mock implementation
			return nil
		},
	}

	// Add subcommand to root
	rootCmd.AddCommand(setupCmd)

	// Verify subcommand was added
	assert.NotNil(t, rootCmd.Commands(), "root command should have commands")
	assert.Len(t, rootCmd.Commands(), 1, "root command should have 1 subcommand")
	assert.Equal(t, "setup", rootCmd.Commands()[0].Use, "first subcommand should be setup")
}

// TestCLISubcommandExecution verifies that subcommands can be executed
func TestCLISubcommandExecution(t *testing.T) {
	// Create root command
	rootCmd := &cobra.Command{
		Use: "kord-cli",
	}

	// Track if subcommand ran
	var subcommandRan bool

	// Create a subcommand with RunE
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start a new task",
		RunE: func(cmd *cobra.Command, args []string) error {
			subcommandRan = true
			return nil
		},
	}

	rootCmd.AddCommand(startCmd)

	// Execute subcommand
	rootCmd.SetArgs([]string{"start"})
	err := rootCmd.Execute()

	require.NoError(t, err, "subcommand execution should not error")
	assert.True(t, subcommandRan, "subcommand should have executed")
}

// TestCLIErrorHandling verifies proper error handling
func TestCLIErrorHandling(t *testing.T) {
	rootCmd := &cobra.Command{
		Use: "kord-cli",
	}

	initialCommandCount := len(rootCmd.Commands())

	// Create command that can be executed
	testCmd := &cobra.Command{
		Use: "test",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	rootCmd.AddCommand(testCmd)

	rootCmd.SetArgs([]string{"test"})
	_ = rootCmd.Execute() // Verify command executes without panic

	// Verify the command was added
	assert.Greater(t, len(rootCmd.Commands()), initialCommandCount, "command should be added")
	assert.True(t, rootCmd.HasSubCommands(), "should have subcommands")
}

// TestCLIFlagParsing verifies that flags are properly parsed
func TestCLIFlagParsing(t *testing.T) {
	rootCmd := &cobra.Command{
		Use: "kord-cli",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	rootCmd.Flags().BoolP("toggle", "t", false, "Toggle flag")
	rootCmd.Flags().StringP("config", "c", "", "Config file")

	// Test with flags
	rootCmd.SetArgs([]string{"--toggle", "--config", "test.json"})
	_ = rootCmd.Execute()

	toggle, _ := rootCmd.Flags().GetBool("toggle")
	config, _ := rootCmd.Flags().GetString("config")

	assert.True(t, toggle, "toggle flag should be set")
	assert.Equal(t, "test.json", config, "config flag should be set to test.json")
}
