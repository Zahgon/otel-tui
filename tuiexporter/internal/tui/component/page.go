package component

import (
	"github.com/rivo/tview"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/telemetry"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/tui/component/page/timeline"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/tui/component/page/topology"
)

type TUIPages struct {
	store    *telemetry.Store
	pages    *tview.Pages
	traces   tview.Primitive
	timeline *timeline.TimelinePage
	topology *topology.TopologyPage
	metrics  tview.Primitive
	logs     tview.Primitive
	modal    tview.Primitive
	current  string
}

func NewTUIPages(store *telemetry.Store, setFocusFn func(tview.Primitive)) *TUIPages {
	_ = "STUB: not implemented"
	return nil
}

// GetPages returns the pages
func (p *TUIPages) GetPages() *tview.Pages {
	_ = "STUB: not implemented"

	// TogglePage toggles Traces & Logs page.
	return nil
}

func (p *TUIPages) TogglePage() { _ = "STUB: not implemented"; return }

func (p *TUIPages) TogglePageReverse() { _ = "STUB: not implemented"; return }

func (p *TUIPages) switchToPage(name string) { _ = "STUB: not implemented"; return }

func (p *TUIPages) registerPages(store *telemetry.Store, setFocusFn func(tview.Primitive)) {
	_ = "STUB: not implemented"
	return
}
