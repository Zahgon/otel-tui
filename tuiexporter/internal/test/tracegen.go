package test

import (
	"testing"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// This is written referencing following code: https://github.com/CtrlSpice/otel-desktop-viewer/blob/af38ec47a37564e5f03b6d9cefa20b2422033e03/desktopexporter/testdata/trace.go
var (
	spanStartTimestamp = pcommon.NewTimestampFromTime(time.Date(2022, 10, 21, 7, 10, 2, 100000000, time.UTC))
	spanEventTimestamp = pcommon.NewTimestampFromTime(time.Date(2022, 10, 21, 7, 10, 2, 150000000, time.UTC))
	spanEndTimestamp   = pcommon.NewTimestampFromTime(time.Date(2022, 10, 21, 7, 10, 2, 300000000, time.UTC))
)

type GeneratedSpans struct {
	Spans  []*ptrace.Span
	RSpans []*ptrace.ResourceSpans
	SSpans []*ptrace.ScopeSpans
}

// This is written referencing following code: https://github.com/CtrlSpice/otel-desktop-viewer/blob/af38ec47a37564e5f03b6d9cefa20b2422033e03/desktopexporter/testdata/trace.go
func GenerateOTLPTracesPayload(t *testing.T, traceID, resourceCount int, scopeCount []int, spanCount [][]int) (ptrace.Traces, *GeneratedSpans) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

// Create and populate resource data

// Create and populate instrumentation scope data

//Create and populate spans

// This is written referencing following code: https://github.com/CtrlSpice/otel-desktop-viewer/blob/af38ec47a37564e5f03b6d9cefa20b2422033e03/desktopexporter/testdata/trace.go
func fillResource(t *testing.T, resource pcommon.Resource, resourceIndex int) {
	_ = "STUB: not implemented"
	return
}

// This is written referencing following code: https://github.com/CtrlSpice/otel-desktop-viewer/blob/af38ec47a37564e5f03b6d9cefa20b2422033e03/desktopexporter/testdata/trace.go
func fillScope(t *testing.T, scope pcommon.InstrumentationScope, resourceIndex, scopeIndex int) {
	_ = "STUB: not implemented"
	return
}

// This is written referencing following code: https://github.com/CtrlSpice/otel-desktop-viewer/blob/af38ec47a37564e5f03b6d9cefa20b2422033e03/desktopexporter/testdata/trace.go
func fillSpan(t *testing.T, span ptrace.Span, traceID, resourceIndex, scopeIndex, spanIndex, uniqueSpanIndex int) {
	_ = "STUB: not implemented"
	return
}

// #nosec G115

// #nosec G115

// GenerateSpanWithDuration returns a span with specified span name and duration.
func GenerateSpanWithDuration(t *testing.T, spanName string, duration time.Duration) *ptrace.Span {
	_ = "STUB: not implemented"
	return nil
}
