package trace

import (
	"github.com/rivo/tview"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/telemetry"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/tui/component/layout"
)

type detail struct {
	commands      *tview.TextView
	view          *tview.Flex
	tree          *tview.TreeView
	resizeManager *layout.ResizeManager
}

func newDetail(
	commands *tview.TextView,
	resizeManager *layout.ResizeManager,
) *detail {
	_ = "STUB: not implemented"
	return nil
}

func (d *detail) flush() { _ = "STUB: not implemented"; return }

func (d *detail) update(spans []*telemetry.SpanData) { _ = "STUB: not implemented"; return }

func (d *detail) getTraceInfoTree(spans []*telemetry.SpanData) *tview.TreeView {
	_ = "STUB: not implemented"
	return nil
}

// statistics

// resource info

// scope info

func (d *detail) updateCommands() { _ = "STUB: not implemented"; return }
