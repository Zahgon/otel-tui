package telemetry

import (
	"sync"

	"go.opentelemetry.io/collector/pdata/ptrace"
)

// SpanDataMap is a map of span id to span data
// This is used to quickly look up a span by its id
type SpanDataMap map[string]*SpanData

// TraceSpanDataMap is a map of trace id to a slice of spans
// This is used to quickly look up all spans in a trace
type TraceSpanDataMap map[string][]*SpanData

// TraceServiceSpanDataMap is a map of trace id and service name to a slice of spans
// This is used to quickly look up all spans in a trace for a service
type TraceServiceSpanDataMap map[string]map[string][]*SpanData

// TraceServiceHasErrorMap is a map of trace id and service name to a flag whether
// the spans have any error status
type TraceServiceHasErrorMap map[string]map[string]bool

// TraceServiceParentIDMap is a map of trace id and service name to a parent span id
// This is used to update service root spans in the trace list.
type TraceServiceParentIDMap map[string]map[string]*SpanData

// TraceCache is a cache of trace spans
type TraceCache struct {
	mu                sync.RWMutex
	spanid2span       SpanDataMap
	traceid2spans     TraceSpanDataMap
	tracesvc2spans    TraceServiceSpanDataMap
	tracesvc2haserror TraceServiceHasErrorMap
	tracesvc2parent   TraceServiceParentIDMap
}

// NewTraceCache returns a new trace cache
func NewTraceCache() *TraceCache { _ = "STUB: not implemented"; return nil }

// UpdateCache updates the cache with a new span
func (c *TraceCache) UpdateCache(sname string, data *SpanData) (newtracesvc bool, replaceSpanID string) {
	_ = "STUB: not implemented"
	return false, ""
}

// This span is higher parent span
// NOTE: In this process, for performance reasons, only adjacent parent-child relationships
//   between spans are evaluated. For example, if the parent-child order of spans is 1, 2, 3, and
//   the arrival order is 3, 1, 2, span 2 will be recognized as the service root span. To recalculate
//   the specific parent-child relationship, use `R` key to trigger deep refreshing

// DeleteCache deletes a list of spans from the cache
func (c *TraceCache) DeleteCache(serviceSpans []*SpanData) { _ = "STUB: not implemented"; return }

// FIXME: more efficient way ?

// delete spans in traceid2spans only if there are no spans left in tracesvc2spans
// for better performance

// GetSpansByTraceID returns all spans for a given trace id
func (c *TraceCache) GetSpansByTraceID(traceID string) ([]*SpanData, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// GetSpansByTraceIDAndSvc returns all spans for a given trace id and service name
func (c *TraceCache) GetSpansByTraceIDAndSvc(traceID, svc string) ([]*SpanData, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// getSpansByTraceIDAndSvcLocked is the lock-free implementation. The caller
// must hold c.mu (read or write).
func (c *TraceCache) getSpansByTraceIDAndSvcLocked(traceID, svc string) ([]*SpanData, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// HasErrorByTraceIDAndSvc returns the flag whether the spans have any errors
func (c *TraceCache) HasErrorByTraceIDAndSvc(traceID, svc string) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

// GetSpanByID returns a span by its id
func (c *TraceCache) GetSpanByID(spanID string) (*SpanData, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *TraceCache) DrawSpanDependencies() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *TraceCache) flush() { _ = "STUB: not implemented"; return }

func spanHasError(span *ptrace.Span) bool { _ = "STUB: not implemented"; return false }

// TraceLogDataMap is a map of trace id to a slice of logs
// This is used to quickly look up all logs in a trace
type TraceLogDataMap map[string][]*LogData

// LogCache is a cache of logs
type LogCache struct {
	mu           sync.RWMutex
	traceid2logs TraceLogDataMap
}

// NewLogCache returns a new log cache
func NewLogCache() *LogCache { _ = "STUB: not implemented"; return nil }

// UpdateCache updates the cache with a new log
func (c *LogCache) UpdateCache(data *LogData) { _ = "STUB: not implemented"; return }

// DeleteCache deletes a list of logs from the cache
func (c *LogCache) DeleteCache(logs []*LogData) { _ = "STUB: not implemented"; return }

// GetLogsByTraceID returns all logs for a given trace id
func (c *LogCache) GetLogsByTraceID(traceID string) ([]*LogData, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *LogCache) flush() { _ = "STUB: not implemented"; return }

// MetricServiceMetricDataMap is a map of service name and metric name to a slice of metrics
// This is used to quickly look up datapoints in a service metric
type MetricServiceMetricDataMap map[string]map[string][]*MetricData

// MetricCache is a cache of metrics
type MetricCache struct {
	mu                sync.RWMutex
	svcmetric2metrics MetricServiceMetricDataMap
}

// NewMetricCache returns a new metric cache
func NewMetricCache() *MetricCache { _ = "STUB: not implemented"; return nil }

// UpdateCache updates the cache with a new metric
func (c *MetricCache) UpdateCache(sname string, data *MetricData) {
	_ = "STUB: not implemented"
	return
}

// DeleteCache deletes a list of metrics from the cache
func (c *MetricCache) DeleteCache(metrics []*MetricData) { _ = "STUB: not implemented"; return }

// GetMetricsBySvcAndMetricName returns all metrics for a given service name and metric name
func (c *MetricCache) GetMetricsBySvcAndMetricName(sname, mname string) ([]*MetricData, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *MetricCache) flush() { _ = "STUB: not implemented"; return }
