package telemetry

import (
	"sync"
	"time"

	"github.com/jonboulle/clockwork"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

const (
	MAX_SERVICE_SPAN_COUNT = 1000
	MAX_METRIC_COUNT       = 3000
	MAX_LOG_COUNT          = 1000
)

// SpanData is a struct to represent a span
type SpanData struct {
	Span         *ptrace.Span
	ResourceSpan *ptrace.ResourceSpans
	ScopeSpans   *ptrace.ScopeSpans
	ReceivedAt   time.Time
}

// IsRoot returns true if the span is a root span
func (sd *SpanData) IsRoot() bool { _ = "STUB: not implemented"; return false }

func (sd *SpanData) GetServiceName() string { _ = "STUB: not implemented"; return "" }

func (sd *SpanData) GetDurationText() string { _ = "STUB: not implemented"; return "" }

func (sd *SpanData) GetReceivedAtText(full bool) string { _ = "STUB: not implemented"; return "" }

func (sd *SpanData) GetSpanName() string { _ = "STUB: not implemented"; return "" }

// SvcSpans is a slice of service spans
// This is a slice of one span of a single service
type SvcSpans []*SpanData

func (ss *SvcSpans) replaceBySpanID(replaceSpanID string, data *SpanData) {
	_ = "STUB: not implemented"
	return
}

// MetricData is a struct to represent a metric
type MetricData struct {
	Metric         *pmetric.Metric
	ResourceMetric *pmetric.ResourceMetrics
	ScopeMetric    *pmetric.ScopeMetrics
	ReceivedAt     time.Time
}

// HasNumberDatapoints returns whether it has number datapoints
func (md *MetricData) HasNumberDatapoints() bool { _ = "STUB: not implemented"; return false }

func (md *MetricData) GetServiceName() string { _ = "STUB: not implemented"; return "" }

func (md *MetricData) GetMetricName() string { _ = "STUB: not implemented"; return "" }

func (md *MetricData) GetMetricTypeText() string { _ = "STUB: not implemented"; return "" }

func (md *MetricData) GetDataPointNum() string { _ = "STUB: not implemented"; return "" }

// LogData is a struct to represent a log
type LogData struct {
	Log         *plog.LogRecord
	ResourceLog *plog.ResourceLogs
	ScopeLog    *plog.ScopeLogs
	ReceivedAt  time.Time
}

func (l *LogData) GetResolvedBody() string { _ = "STUB: not implemented"; return "" }

func (l *LogData) GetTraceID() string { _ = "STUB: not implemented"; return "" }

func (l *LogData) GetServiceName() string { _ = "STUB: not implemented"; return "" }

func (l *LogData) GetTimestampText(full bool) string { _ = "STUB: not implemented"; return "" }

func (l *LogData) GetSeverity() string { _ = "STUB: not implemented"; return "" }

func (l *LogData) GetEventName() string {
	_ = "STUB: not implemented"
	// see: https://github.com/open-telemetry/semantic-conventions/blob/a4fc971e0c7ffa4b9572654f075d3cb8560db770/docs/general/events.md#event-definition
	return ""
}

func (l *LogData) GetRawData() string { _ = "STUB: not implemented"; return "" }

// Store is a store of trace spans
type Store struct {
	mut                 sync.Mutex
	clockwork           clockwork.Clock
	filterSvc           string
	filterMetric        string
	filterLog           string
	sortTrace           SortType
	svcspans            SvcSpans
	svcspansFiltered    SvcSpans
	tracecache          *TraceCache
	metrics             []*MetricData
	metricsFiltered     []*MetricData
	metriccache         *MetricCache
	logs                []*LogData
	logsFiltered        []*LogData
	logcache            *LogCache
	updatedAt           time.Time
	maxServiceSpanCount int
	maxMetricCount      int
	maxLogCount         int
	onSpanAdded         func()
	onMetricAdded       func()
	onLogAdded          func()
	onFlushed           []func()
}

// NewStore creates a new store
func NewStore(clock clockwork.Clock) *Store { _ = "STUB: not implemented"; return nil }

// TODO: make this configurable
// TODO: make this configurable
// TODO: make this configurable

// GetTraceCache returns the trace cache
func (s *Store) GetTraceCache() *TraceCache { _ = "STUB: not implemented"; return nil }

// GetMetricCache returns the metric cache
func (s *Store) GetMetricCache() *MetricCache { _ = "STUB: not implemented"; return nil }

// GetLogCache returns the log cache
func (s *Store) GetLogCache() *LogCache {
	_ = "STUB: not implemented"

	// GetSvcSpans returns the service spans in the store
	return nil
}

func (s *Store) GetSvcSpans() *SvcSpans {
	_ = "STUB: not implemented"

	// GetFilteredSvcSpans returns the filtered service spans in the store
	return nil
}

func (s *Store) GetFilteredSvcSpans() *SvcSpans { _ = "STUB: not implemented"; return nil }

// GetFilteredMetrics returns the filetered metrics in the store
func (s *Store) GetFilteredMetrics() *[]*MetricData { _ = "STUB: not implemented"; return nil }

// GetFilteredLogs returns the filtered logs in the store
func (s *Store) GetFilteredLogs() *[]*LogData { _ = "STUB: not implemented"; return nil }

// UpdatedAt returns the last updated time
func (s *Store) UpdatedAt() time.Time {
	_ = "STUB: not implemented"

	// SetOnSpanAdded sets the callback function to be called when a span is added
	return *new(time.Time)
}

func (s *Store) SetOnSpanAdded(f func()) {
	_ = "STUB: not implemented"

	// SetOnMetricAdded sets the callback function to be called when a metric is added
	return
}

func (s *Store) SetOnMetricAdded(f func()) { _ = "STUB: not implemented"; return }

// SetOnLogAdded sets the callback function to be called when a log is added
func (s *Store) SetOnLogAdded(f func()) {
	_ = "STUB: not implemented"

	// RegisterOnFlushed registers a callback function to be called when the store is flushed
	return
}

func (s *Store) RegisterOnFlushed(f func()) { _ = "STUB: not implemented"; return }

// ApplyFilterTraces applies a filter and sort to the traces
func (s *Store) ApplyFilterTraces(svc string, sortType SortType) { _ = "STUB: not implemented"; return }

func (s *Store) updateFilterService() { _ = "STUB: not implemented"; return }

// ApplyFilterMetrics applies a filter to the metrics
func (s *Store) ApplyFilterMetrics(filter string) { _ = "STUB: not implemented"; return }

func (s *Store) updateFilterMetrics() { _ = "STUB: not implemented"; return }

// ApplyFilterLogs applies a filter to the logs
func (s *Store) ApplyFilterLogs(filter string) { _ = "STUB: not implemented"; return }

func (s *Store) updateFilterLogs() { _ = "STUB: not implemented"; return }

// GetTraceIDByFilteredIdx returns the trace at the given index
func (s *Store) GetTraceIDByFilteredIdx(idx int) string { _ = "STUB: not implemented"; return "" }

// GetFilteredServiceSpansByIdx returns the spans for a given service at the given index
func (s *Store) GetFilteredServiceSpansByIdx(idx int) []*SpanData {
	_ = "STUB: not implemented"
	return nil
}

// RecalculateServiceRootSpanByIdx recalculates service root span of the specified index
func (s *Store) RecalculateServiceRootSpanByIdx(idx int) { _ = "STUB: not implemented"; return }

// TODO: Condider orphan span?

// GetFilteredMetricByIdx returns the metric at the given index
func (s *Store) GetFilteredMetricByIdx(idx int) *MetricData { _ = "STUB: not implemented"; return nil }

// GetFilteredLogByIdx returns the log at the given index
func (s *Store) GetFilteredLogByIdx(idx int) *LogData { _ = "STUB: not implemented"; return nil }

// AddSpan adds spans to the store
func (s *Store) AddSpan(traces *ptrace.Traces) { _ = "STUB: not implemented"; return }

// FIXME: More efficient logic is needed

// data rotation

// AddMetric adds metrics to the store
func (s *Store) AddMetric(metrics *pmetric.Metrics) { _ = "STUB: not implemented"; return }

// data rotation

// AddLog adds logs to the store
func (s *Store) AddLog(logs *plog.Logs) { _ = "STUB: not implemented"; return }

// data rotation

// Flush clears the store including the cache
func (s *Store) Flush() { _ = "STUB: not implemented"; return }
