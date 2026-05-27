package topology

import (
	"github.com/rivo/tview"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/telemetry"
)

type TopologyPage struct {
	view  *tview.Flex
	topo  *tview.TextView
	cache *telemetry.TraceCache
}

func NewTopologyPage(cache *telemetry.TraceCache) *TopologyPage {
	_ = "STUB: not implemented"
	return nil
}

func (p *TopologyPage) GetPrimitive() tview.Primitive {
	_ = "STUB: not implemented"
	return *new(tview.Primitive)
}

func (p *TopologyPage) registerCommands(commands *tview.TextView) {
	_ = "STUB: not implemented"
	return
}

func (p *TopologyPage) UpdateTopology() { _ = "STUB: not implemented"; return }
