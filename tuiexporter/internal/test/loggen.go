package test

import (
	"testing"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

var (
	logTimestamp         = pcommon.NewTimestampFromTime(time.Date(2022, 10, 21, 7, 10, 2, 100000000, time.UTC))
	logObservedTimestamp = pcommon.NewTimestampFromTime(time.Date(2022, 10, 21, 7, 10, 2, 200000000, time.UTC))
)

type GeneratedLogs struct {
	Logs  []*plog.LogRecord
	RLogs []*plog.ResourceLogs
	SLogs []*plog.ScopeLogs
}

func GenerateOTLPLogsPayload(t *testing.T, traceID, resourceCount int, scopeCount []int, spanCount [][]int) (plog.Logs, *GeneratedLogs) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

// Create and populate resource data

// Create and populate instrumentation scope data

// Create and populate logs

// 2 logs per span

func fillLog(t *testing.T, l plog.LogRecord, traceID, resourceIndex, scopeIndex, spanIndex, logIndex, uniqueSpanIndex int) {
	_ = "STUB: not implemented"
	return
}

// #nosec G115

// #nosec G115
