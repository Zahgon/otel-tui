package trace

import (
	"github.com/rivo/tview"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/telemetry"
)

const (
	defaultTableProportion  = 30
	defaultDetailProportion = 20
)

type TracePage struct {
	view   *tview.Flex
	table  *table
	detail *detail
}

func NewTracePage(
	onSelectTableRow func(row, column int),
	store *telemetry.Store,
) *TracePage {
	_ = "STUB: not implemented"
	return nil
}

func (p *TracePage) GetPrimitive() tview.Primitive {
	_ = "STUB: not implemented"
	return *new(tview.Primitive)
}

func (p *TracePage) flush() { _ = "STUB: not implemented"; return }

func (p *TracePage) registerCommands() { _ = "STUB: not implemented"; return }
