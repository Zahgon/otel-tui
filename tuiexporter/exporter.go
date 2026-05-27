package tuiexporter

import (
	"context"

	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/tui"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type tuiExporter struct {
	app *tui.TUIApp
}

func newTuiExporter(config *Config) (*tuiExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FIXME: When reading telemetry from a JSON file on startup, the UI will break
//        if it runs at the same time as the UI drawing. As a workaround, wait for a second.

func (e *tuiExporter) pushTraces(_ context.Context, traces ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tuiExporter) pushMetrics(_ context.Context, metrics pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tuiExporter) pushLogs(_ context.Context, logs plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// Start runs the TUI exporter
func (e *tuiExporter) Start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Shutdown stops the TUI exporter
func (e *tuiExporter) Shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }
