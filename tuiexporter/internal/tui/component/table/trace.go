package table

import (
	"github.com/rivo/tview"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/telemetry"
)

var defaultSpanCellMappers = cellMappers[telemetry.SpanData]{
	1: {
		header: "Service Name",
		getTextRowFn: func(data *telemetry.SpanData) string {
			return data.GetServiceName()
		},
	},
	2: {
		header: "Latency",
		getTextRowFn: func(data *telemetry.SpanData) string {
			return data.GetDurationText()
		},
	},
	3: {
		header: "Received At",
		getTextRowFn: func(data *telemetry.SpanData) string {
			panic("Received At column should be overridden")
		},
	},
	4: {
		header: "Span Name",
		getTextRowFn: func(data *telemetry.SpanData) string {
			return data.GetSpanName()
		},
	},
}

// SpanDataForTable is a wrapper for spans to be displayed in a table.
type SpanDataForTable struct {
	tview.TableContentReadOnly
	tcache         *telemetry.TraceCache
	spans          *telemetry.SvcSpans
	sortType       *telemetry.SortType
	mapper         cellMappers[telemetry.SpanData]
	isFullDatetime bool
}

// NewSpanDataForTable creates a new SpanDataForTable.
func NewSpanDataForTable(tcache *telemetry.TraceCache, spans *telemetry.SvcSpans, sortType *telemetry.SortType) SpanDataForTable {
	_ = "STUB: not implemented"
	return *new(SpanDataForTable)
}

// SetFullDatetime sets the full datetime flag for the table.
func (s *SpanDataForTable) SetFullDatetime(full bool) { _ = "STUB: not implemented"; return }

// IsFullDatetime returns the full datetime flag for the table.
func (s SpanDataForTable) IsFullDatetime() bool { _ = "STUB: not implemented"; return false }

func (s *SpanDataForTable) updateReceivedAtMapper() { _ = "STUB: not implemented"; return }

// implementations for tview Virtual Table
// see: https://github.com/rivo/tview/wiki/VirtualTable
func (s SpanDataForTable) GetCell(row, column int) *tview.TableCell {
	_ = "STUB: not implemented"
	return nil
}

func (s SpanDataForTable) GetRowCount() int { _ = "STUB: not implemented"; return 0 }

func (s SpanDataForTable) GetColumnCount() int { _ = "STUB: not implemented"; return 0 }

// including error indicator

func (s SpanDataForTable) getErrorIndicator(span *telemetry.SpanData) *tview.TableCell {
	_ = "STUB: not implemented"
	return nil
}

func (s SpanDataForTable) getHeaderCell(column int, sortType telemetry.SortType) *tview.TableCell {
	_ = "STUB: not implemented"
	return nil
}

// Error indicator
