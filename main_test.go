package main

import (
	"testing"
)

func TestVersion(t *testing.T) {
	// Test that the version variable exists and can be set
	originalVersion := Version
	Version = "test-version"

	if Version != "test-version" {
		t.Errorf("Expected Version to be 'test-version', got '%s'", Version)
	}

	// Restore original version
	Version = originalVersion
}

func TestMain(t *testing.T) {
	// This test ensures the main package compiles without errors
	// The actual main() function requires stdin/stdout interaction
	// so we just verify the package structure is correct
	if Version == "" {
		Version = "dev"
	}

	if Version != "dev" && Version != "test-version" {
		t.Logf("Version is set to: %s", Version)
	}
}
