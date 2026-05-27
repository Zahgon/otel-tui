package timeline

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/telemetry"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/tui/component/layout"
)

const (
	spanNameColumnWidthResizeUnit = 5
	spanNameColumnWidthDefalt     = 30
)

type spanTreeNode struct {
	span     *telemetry.SpanData
	label    string
	box      *tview.Box
	children []*spanTreeNode
	expand   bool
}

type grid struct {
	commands      *tview.TextView
	gridView      *tview.Grid
	tcache        *telemetry.TraceCache
	snameWidth    int
	totalRow      int
	currentRow    int
	tree          []*spanTreeNode
	duration      time.Duration
	nodes         []*spanTreeNode
	items         []*tview.TextView
	resizeManager *layout.ResizeManager
	detail        *detail
	logPane       *logPane
}

func newGrid(
	commands *tview.TextView,
	tcache *telemetry.TraceCache,
	resizeManager *layout.ResizeManager,
	detail *detail,
	logPane *logPane,
) *grid {
	_ = "STUB: not implemented"
	return nil
}

func (g *grid) updateGrid(traceID string) *telemetry.SpanData {
	_ = "STUB: not implemented"
	return nil
}

func (g *grid) prepareTimeline(duration time.Duration) { _ = "STUB: not implemented"; return }

// Draw a horizontal line across the middle of the box.

// Write some text along the horizontal line.

// Space for other content.

func (g *grid) placeSpans() { _ = "STUB: not implemented"; return }

func (g *grid) placeSpan(
	node *spanTreeNode,
	row, depth int,
	tvs *[]*tview.TextView,
	nodes *[]*spanTreeNode,
) int {
	_ = "STUB: not implemented"
	return 0
}

func (g *grid) newSpanTree(traceID string) (rootNodes []*spanTreeNode, duration time.Duration) {
	_ = "STUB: not implemented"
	return nil, *new(time.Duration)
}

// store memo and calculate start and end time of the trace

// color is assigned by the service name

// generate span tree

// sort root spans by start time

func (g *grid) newTextView(text string) *tview.TextView { _ = "STUB: not implemented"; return nil }

func (g *grid) updateCommands() { _ = "STUB: not implemented"; return }

// Remove default input capture to avoid conflict

func (g *grid) getCurrentSpan() *telemetry.SpanData { _ = "STUB: not implemented"; return nil }

func (g *grid) stepBy(step int) func(_ *tcell.EventKey) *tcell.EventKey {
	_ = "STUB: not implemented"
	return nil
}

func (g *grid) goToFirst(_ *tcell.EventKey) *tcell.EventKey { _ = "STUB: not implemented"; return nil }

func (g *grid) goToLast(_ *tcell.EventKey) *tcell.EventKey { _ = "STUB: not implemented"; return nil }

// g.gridView.SetOffset(g.currentRow, 0)

func (g *grid) updateCurrentSpan() { _ = "STUB: not implemented"; return }

func createSpan(color tcell.Color, total, start, end time.Duration) (span *tview.Box) {
	_ = "STUB: not implemented"
	return nil
}

// Draw a horizontal line across the middle of the box.

// Space for other content.

func getXByRatio(ratio float64, width int) int { _ = "STUB: not implemented"; return 0 }

func calculateTimelineUnit(duration time.Duration) (unit time.Duration, count int) {
	_ = "STUB: not implemented"
	// TODO: set count depends on the width
	return *new(time.Duration), 0
}

func roundDownDuration(d time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func narrowInLimit(step, curr, limit int) int { _ = "STUB: not implemented"; return 0 }

func widenInLimit(step, curr, limit int) int { _ = "STUB: not implemented"; return 0 }
