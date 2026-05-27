package timeline

import (
	"github.com/rivo/tview"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/telemetry"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/tui/component/table"
)

type logPane struct {
	commands  *tview.TextView
	tableView *tview.Table
	lcache    *telemetry.LogCache
	logData   *table.LogDataForTable
	allLogs   bool
}

func newLogPane(
	commands *tview.TextView,
	lcache *telemetry.LogCache,
) *logPane {
	_ = "STUB: not implemented"
	return nil
}

func (l *logPane) updateLog(traceID, spanID string) { _ = "STUB: not implemented"; return }

func (l *logPane) toggleAllLogs(traceID string, currentSpan *telemetry.SpanData) {
	_ = "STUB: not implemented"
	return
}

func (l *logPane) updateCommands() { _ = "STUB: not implemented"; return }
