package test

import (
	"bytes"
	"testing"

	"github.com/gdamore/tcell/v2"
)

func projectRoot(t *testing.T) string { _ = "STUB: not implemented"; return "" }

func LoadTestdata(t *testing.T, name string) string { _ = "STUB: not implemented"; return "" }

//gosec:disable G304 -- This is a test code

func GetScreenContent(t *testing.T, screen tcell.SimulationScreen) bytes.Buffer {
	_ = "STUB: not implemented"
	return *new(bytes.Buffer)
}

// Replace SimulationScreen's default fill character ('X') with space.
// Since tcell v2.12.0, the Put API changes (PR #846, #848) caused the fill
// character to become visible in test environments when tview components don't
// fully occupy their allocated space. This doesn't occur in actual terminals,
// so we normalize the test environment to match real-world behavior.
// See https://github.com/gdamore/tcell/pull/848

// To avoid replacing intentional 'X' characters (e.g., in text content or
// control characters like Ctrl-X), we only replace 'X' when it appears as
// part of a consecutive fill pattern.
