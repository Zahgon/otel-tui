package timeline

import (
	"github.com/rivo/tview"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/telemetry"
)

const (
	defaultGridProportion   = 29
	defaultDetailProportion = 21
)

type TimelinePage struct {
	switchToPageFn func()
	commands       *tview.TextView
	base           *tview.Flex
	container      *tview.Flex
	mainContainer  *tview.Flex
	store          *telemetry.Store
	onEscape       func()
	detail         *detail
	grid           *grid
	logPane        *logPane
	isLogCollapsed bool
	traceID        string
}

func NewTimelinePage(
	switchToPageFn func(),
	store *telemetry.Store,
	onEscape func(),
) *TimelinePage {
	_ = "STUB: not implemented"
	return nil
}

func (p *TimelinePage) GetPrimitive() tview.Primitive {
	_ = "STUB: not implemented"
	return *new(tview.Primitive)
}

func (p *TimelinePage) ShowTimelineByRow(row int) { _ = "STUB: not implemented"; return }

func (p *TimelinePage) DrawTimeline(traceID string) { _ = "STUB: not implemented"; return }

func (p *TimelinePage) registerCommands() { _ = "STUB: not implemented"; return }

func (p *TimelinePage) updateContainer() { _ = "STUB: not implemented"; return }
