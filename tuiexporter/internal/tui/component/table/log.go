package table

import (
	"github.com/rivo/tview"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/telemetry"
)

var defaultLogCellMappers = cellMappers[telemetry.LogData]{
	0: {
		header: "Trace ID",
		getTextRowFn: func(log *telemetry.LogData) string {
			return log.GetTraceID()
		},
	},
	1: {
		header: "Service Name",
		getTextRowFn: func(log *telemetry.LogData) string {
			return log.GetServiceName()
		},
	},
	2: {
		header: "Timestamp",
		getTextRowFn: func(log *telemetry.LogData) string {
			panic("Timestamp column should be overridden")
		},
	},
	3: {
		header: "Severity",
		getTextRowFn: func(log *telemetry.LogData) string {
			return log.GetSeverity()
		},
	},
	4: {
		header: "Event Name",
		getTextRowFn: func(log *telemetry.LogData) string {
			return log.GetEventName()
		},
	},
	5: {
		header: "RawData",
		getTextRowFn: func(log *telemetry.LogData) string {
			return log.GetRawData()
		},
	},
}

var logCellMappersForTimeline = cellMappers[telemetry.LogData]{
	0: {
		header: "Service Name",
		getTextRowFn: func(log *telemetry.LogData) string {
			return log.GetServiceName()
		},
	},
	1: {
		header: "Timestamp",
		getTextRowFn: func(log *telemetry.LogData) string {
			panic("Timestamp column should be overridden")
		},
	},
	2: {
		header: "Severity",
		getTextRowFn: func(log *telemetry.LogData) string {
			return log.GetSeverity()
		},
	},
	3: {
		header: "Event Name",
		getTextRowFn: func(log *telemetry.LogData) string {
			return log.GetEventName()
		},
	},
	4: {
		header: "RawData",
		getTextRowFn: func(log *telemetry.LogData) string {
			return log.GetRawData()
		},
	},
}

// LogDataForTable is a wrapper for logs to be displayed in a table
type LogDataForTable struct {
	tview.TableContentReadOnly
	logs           *[]*telemetry.LogData
	mapper         cellMappers[telemetry.LogData]
	isFullDatetime bool
}

// NewLogDataForTable creates a new LogDataForTable.
func NewLogDataForTable(logs *[]*telemetry.LogData) LogDataForTable {
	_ = "STUB: not implemented"
	return *new(LogDataForTable)
}

// NewLogDataForTableForTimeline creates a new LogDataForTable for timeline page.
func NewLogDataForTableForTimeline(logs *[]*telemetry.LogData) LogDataForTable {
	_ = "STUB: not implemented"
	return *new(LogDataForTable)
}

// SetFullDatetime sets whether to display full datetime or not
func (l *LogDataForTable) SetFullDatetime(full bool) { _ = "STUB: not implemented"; return }

// IsFullDatetime returns whether to display full datetime or not
func (l LogDataForTable) IsFullDatetime() bool { _ = "STUB: not implemented"; return false }

func (l *LogDataForTable) updateTimestampMapper() { _ = "STUB: not implemented"; return }

// implementation for tableModalMapper interface
func (l *LogDataForTable) GetColumnIdx() int { _ = "STUB: not implemented"; return 0 }

// implementations for tview Virtual Table
// see: https://github.com/rivo/tview/wiki/VirtualTable
func (l LogDataForTable) GetCell(row, column int) *tview.TableCell {
	_ = "STUB: not implemented"
	return nil
}

func (l LogDataForTable) GetRowCount() int { _ = "STUB: not implemented"; return 0 }

func (l LogDataForTable) GetColumnCount() int { _ = "STUB: not implemented"; return 0 }

func (l LogDataForTable) getHeaderCell(column int) *tview.TableCell {
	_ = "STUB: not implemented"
	return nil
}
