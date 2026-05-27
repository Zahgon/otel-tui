package metric

import (
	"github.com/rivo/tview"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/telemetry"
)

const (
	defaultTableProportion = 25
	defaultSideProportion  = 25
	defaultDetaiProportion = 25
	defaultChartProportion = 25
)

type MetricPage struct {
	view   *tview.Flex
	table  *table
	detail *detail
	chart  *chart
}

func NewMetricPage(
	store *telemetry.Store,
) *MetricPage {
	_ = "STUB: not implemented"
	return nil
}

func (p *MetricPage) GetPrimitive() tview.Primitive {
	_ = "STUB: not implemented"
	return *new(tview.Primitive)
}

func (p *MetricPage) flush() { _ = "STUB: not implemented"; return }

func (p *MetricPage) registerCommands() { _ = "STUB: not implemented"; return }
