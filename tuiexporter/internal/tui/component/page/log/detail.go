package log

import (
	"github.com/rivo/tview"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/telemetry"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/tui/component/layout"
)

type detail struct {
	commands       *tview.TextView
	view           *tview.Flex
	tree           *tview.TreeView
	drawTimelineFn func(traceID string)
	resizeManagers []*layout.ResizeManager
	tcache         *telemetry.TraceCache
}

func newDetail(
	commands *tview.TextView,
	drawTimelineFn func(traceID string),
	resizeManagers []*layout.ResizeManager,
	tcache *telemetry.TraceCache,
) *detail {
	_ = "STUB: not implemented"
	return nil
}

func (d *detail) flush() { _ = "STUB: not implemented"; return }

func (d *detail) update(l *telemetry.LogData) { _ = "STUB: not implemented"; return }

func (d *detail) getLogInfoTree(l *telemetry.LogData) *tview.TreeView {
	_ = "STUB: not implemented"
	return nil
}

// resource info

// scope info

// log body

func (d *detail) updateCommands() { _ = "STUB: not implemented"; return }
