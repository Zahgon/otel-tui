package log

import (
	"github.com/rivo/tview"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/telemetry"
)

const (
	defaultTableProportion  = 30
	defaultDetailProportion = 20
	defaultMainProportion   = 15
	defaultBodyProportion   = 3
)

type LogPage struct {
	view   *tview.Flex
	table  *table
	detail *detail
	body   *body
}

func NewLogPage(
	drawTimelineFn func(traceID string),
	store *telemetry.Store,
) *LogPage {
	_ = "STUB: not implemented"
	return nil
}

func (p *LogPage) GetPrimitive() tview.Primitive {
	_ = "STUB: not implemented"
	return *new(tview.Primitive)
}

func (p *LogPage) flush() { _ = "STUB: not implemented"; return }

func (p *LogPage) registerCommands() { _ = "STUB: not implemented"; return }
