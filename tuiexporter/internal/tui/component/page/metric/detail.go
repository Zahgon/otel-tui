package metric

import (
	"github.com/rivo/tview"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/telemetry"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/tui/component/layout"
)

type detail struct {
	commands       *tview.TextView
	view           *tview.Flex
	tree           *tview.TreeView
	resizeManagers []*layout.ResizeManager
}

func newDetail(
	commands *tview.TextView,
	resizeManagers []*layout.ResizeManager,
) *detail {
	_ = "STUB: not implemented"
	return nil
}

func (d *detail) flush() { _ = "STUB: not implemented"; return }

func (d *detail) update(m *telemetry.MetricData) { _ = "STUB: not implemented"; return }

func (d *detail) getMetricInfoTree(m *telemetry.MetricData) *tview.TreeView {
	_ = "STUB: not implemented"
	return nil
}

// resource info

// scope info

// metric

/// metadata

/// datapoints

// value

// flags

// exampler

// value

// filtered attributes

// attributes

// value

// flags

// exampler

// value

// filtered attributes

// attributes

// flags

// exampler

// value

// filtered attributes

// attributes

// flags

// exampler

// value

// filtered attributes

// attributes

// quantile

// flags

// attributes

func (d *detail) updateCommands() { _ = "STUB: not implemented"; return }
