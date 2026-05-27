package trace

import (
	"github.com/rivo/tview"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/telemetry"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/tui/component/filter"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/tui/component/layout"
	ctable "github.com/ymtdzzz/otel-tui/tuiexporter/internal/tui/component/table"
)

type table struct {
	store    *telemetry.Store
	view     *tview.Flex
	table    *tview.Table
	spanData *ctable.SpanDataForTable
	filter   *filter.Filter
	detail   *detail
}

func newTable(
	commands *tview.TextView,
	onSelectTableRow func(row, column int),
	store *telemetry.Store,
	detail *detail,
	resizeManager *layout.ResizeManager,
) *table {
	_ = "STUB: not implemented"
	return nil
}

// Select the first data row (row 1), not the header (row 0)

func (t *table) registerCommands(commands *tview.TextView, resizeManager *layout.ResizeManager) {
	_ = "STUB: not implemented"
	return
}

func (t *table) onSelectionChangedFunc() func(row, col int) { _ = "STUB: not implemented"; return nil }
