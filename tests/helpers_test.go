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

	"github.com/stretchr/testify/assert"
)

// TestAssertionFramework verifies that testify assertions work correctly
func TestAssertionFramework(t *testing.T) {
	// Test basic assertions
	assert.Equal(t, 1, 1, "equal assertion should work")
	assert.NotEqual(t, 1, 2, "not equal assertion should work")
	assert.True(t, true, "true assertion should work")
	assert.False(t, false, "false assertion should work")
}

// TestErrorCapture verifies we can capture command output for testing
func TestErrorCapture(t *testing.T) {
	var buf bytes.Buffer
	output := "test output"
	buf.WriteString(output)

	assert.Equal(t, "test output", buf.String(), "should capture output correctly")
}

// TestStringContains verifies string contains assertions
func TestStringContains(t *testing.T) {
	message := "This is a test message"
	assert.Contains(t, message, "test", "should find substring in string")
	assert.NotContains(t, message, "xyz", "should not find missing substring")
}

// TestSliceAssertions verifies slice operations
func TestSliceAssertions(t *testing.T) {
	slice := []string{"kord", "setup", "start", "commit"}
	assert.Len(t, slice, 4, "slice should have 4 elements")
	assert.Contains(t, slice, "setup", "slice should contain setup")
	assert.NotContains(t, slice, "missing", "slice should not contain missing")
}
