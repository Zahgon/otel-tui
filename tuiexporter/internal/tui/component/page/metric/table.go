package metric

import (
	"github.com/rivo/tview"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/telemetry"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/tui/component/filter"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/tui/component/layout"
	ctable "github.com/ymtdzzz/otel-tui/tuiexporter/internal/tui/component/table"
)

type table struct {
	store      *telemetry.Store
	view       *tview.Flex
	table      *tview.Table
	metricData *ctable.MetricDataForTable
	filter     *filter.Filter
	detail     *detail
	chart      *chart
}

func newTable(
	commands *tview.TextView,
	store *telemetry.Store,
	detail *detail,
	chart *chart,
	resizeManagers []*layout.ResizeManager,
) *table {
	_ = "STUB: not implemented"
	return nil
}

// Select the first data row (row 1), not the header (row 0)

func (t *table) registerCommands(commands *tview.TextView, resizeManagers []*layout.ResizeManager) {
	_ = "STUB: not implemented"
	return
}

func (t *table) onSelectionChangedFunc() func(row, col int) { _ = "STUB: not implemented"; return nil }
