package metric

import (
	"math"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/telemetry"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/tui/component/layout"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

const nullValueFloat64 = math.MaxFloat64

type chart struct {
	commands       *tview.TextView
	view           *tview.Flex
	ch             *tview.Flex
	focusTargets   []layout.FocusableBox
	store          *telemetry.Store
	resizeManagers []*layout.ResizeManager
}

func newChart(
	commands *tview.TextView,
	store *telemetry.Store,
	resizeManagers []*layout.ResizeManager,
) *chart {
	_ = "STUB: not implemented"
	return nil
}

func (c *chart) flush() { _ = "STUB: not implemented"; return }

func (c *chart) update(m *telemetry.MetricData) { _ = "STUB: not implemented"; return }

type ByTimestamp []*pmetric.NumberDataPoint

func (a ByTimestamp) Len() int           { _ = "STUB: not implemented"; return 0 }
func (a ByTimestamp) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (a ByTimestamp) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (c *chart) drawMetricChartByRow(m *telemetry.MetricData) layout.KeyMaps {
	_ = "STUB: not implemented"
	return *new(layout.KeyMaps)
}

func (c *chart) drawMetricHistogramChart(m *telemetry.MetricData) layout.KeyMaps {
	_ = "STUB: not implemented"
	return *new(layout.KeyMaps)
}

func (c *chart) drawMetricNumberChart(m *telemetry.MetricData) layout.KeyMaps {
	_ = "STUB: not implemented"
	return *new(layout.KeyMaps)
}

// attribute name and value map

// sort keys

// TODO: Delete it after implementing drawMetric* for all types

// Draw a chart of the first attribute

func (c *chart) getDataToDraw(dataMap map[string]map[string][]*pmetric.NumberDataPoint, attrkey string, start, end time.Time) ([][]float64, *tview.TextView) {
	_ = "STUB: not implemented"
	// Sort keys
	return nil, nil
}

// Count datapoints

// Set null value

// Set values to timestamp relative position.
// Note that this process keeps values between corresponding positions null value.
// ex: [1.2 1.3 null 1.6 1.1 null null 2.5]

// Get timestamp and locate it to relative position

// Replace null value with appropriate value for smooth line
// ex: [1.2 1.3 1.45 1.6 1.1 1.56 2.02 2.5]

// Fill after the last element

// Fill before the first element

func (c *chart) updateCommands(keyMaps layout.KeyMaps) { _ = "STUB: not implemented"; return }

// uint64ToInt converts uint64 into int. When the input is larger than math.MaxInt, it returns math.MaxInt.
func uint64ToInt(u uint64) int { _ = "STUB: not implemented"; return 0 }

// lineColors returns a color slice for the plot widget.
// If n <= len(layout.Colors), returns the original slice (no allocation).
// Otherwise, returns a new slice with colors cycling via modulo.
func lineColors(n int) []tcell.Color { _ = "STUB: not implemented"; return nil }
