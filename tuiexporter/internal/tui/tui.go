package tui

import (
	"os"
	"time"

	"github.com/rivo/tview"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/telemetry"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/tui/component"
)

const refreshInterval = 500 * time.Millisecond

// TUIApp is the TUI application.
type TUIApp struct {
	initialInterval time.Duration
	app             *tview.Application
	pages           *component.TUIPages
	store           *telemetry.Store
	refreshedAt     time.Time
	logFile         *os.File
}

// NewTUIApp creates a new TUI application.
func NewTUIApp(store *telemetry.Store, initialInterval time.Duration, debugLogFilePath string) (*TUIApp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Disable logging if no file is specified

// Send SGITERM to self on Ctrl+C to ensure global signal handlers are triggered
// Prevents the need for pressing Ctrl+C twice due to tview consuming the first Ctrl+C

// NOTE: Do not return nil here, as we want to allow the event to propagate for Windows

// Store returns the store
func (t *TUIApp) Store() *telemetry.Store {
	_ = "STUB: not implemented"

	// Run starts the TUI application.
	return nil
}

func (t *TUIApp) Run() error { _ = "STUB: not implemented"; return nil }

// Stop stops the TUI application.
func (t *TUIApp) Stop() error { _ = "STUB: not implemented"; return nil }

func (t *TUIApp) refresh() { _ = "STUB: not implemented"; return }
