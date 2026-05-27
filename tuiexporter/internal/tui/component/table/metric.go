package table

import (
	"github.com/rivo/tview"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/telemetry"
)

var defaultMetricCellMappers = cellMappers[telemetry.MetricData]{
	0: {
		header: "Service Name",
		getTextRowFn: func(data *telemetry.MetricData) string {
			return data.GetServiceName()
		},
	},
	1: {
		header: "Metric Name",
		getTextRowFn: func(data *telemetry.MetricData) string {
			return data.GetMetricName()
		},
	},
	2: {
		header: "Metric Type",
		getTextRowFn: func(data *telemetry.MetricData) string {
			return data.GetMetricTypeText()
		},
	},
	3: {
		header: "Data Point Count",
		getTextRowFn: func(data *telemetry.MetricData) string {
			return data.GetDataPointNum()
		},
	},
}

type MetricDataForTable struct {
	tview.TableContentReadOnly
	metrics *[]*telemetry.MetricData
	mapper  cellMappers[telemetry.MetricData]
}

func NewMetricDataForTable(metrics *[]*telemetry.MetricData) MetricDataForTable {
	_ = "STUB: not implemented"
	return *new(MetricDataForTable)
}

// implementations for tview Virtual Table
// see: https://github.com/rivo/tview/wiki/VirtualTable
func (m MetricDataForTable) GetCell(row, column int) *tview.TableCell {
	_ = "STUB: not implemented"
	return nil
}

func (m MetricDataForTable) GetRowCount() int { _ = "STUB: not implemented"; return 0 }

func (m MetricDataForTable) GetColumnCount() int { _ = "STUB: not implemented"; return 0 }

func (m MetricDataForTable) getHeaderCell(column int) *tview.TableCell {
	_ = "STUB: not implemented"
	return nil
}
