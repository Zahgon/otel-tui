package test

import (
	"testing"

	"go.opentelemetry.io/collector/pdata/pmetric"
)

type GeneratedMetrics struct {
	Metrics  []*pmetric.Metric
	RMetrics []*pmetric.ResourceMetrics
	SMetrics []*pmetric.ScopeMetrics
}

func GenerateOTLPGaugeMetricsPayload(t *testing.T, resourceCount int, scopeCount []int, dpCount [][]int) (pmetric.Metrics, *GeneratedMetrics) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// Create and populate resource data

// Create and populate instrumentation scope data

// Create and populate metrics
// 1 metric per scope

func GenerateOTLPHistogramMetricsPayload(t *testing.T, resourceCount int, scopeCount []int, dpCount [][]int) (pmetric.Metrics, *GeneratedMetrics) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// Create and populate resource data

// Create and populate instrumentation scope data

// Create and populate metrics
// 1 metric per scope

func fillMetric(t *testing.T, m pmetric.Metric, resourceIndex, scopeIndex int) {
	_ = "STUB: not implemented"
	return
}

func fillNumberDataPoint(t *testing.T, dp pmetric.NumberDataPoint, dpIndex int) {
	_ = "STUB: not implemented"
	return
}

// TODO: examplers

func fillHistogramDataPoint(t *testing.T, dp pmetric.HistogramDataPoint, dpIndex int) {
	_ = "STUB: not implemented"
	return
}

// #nosec G115

// TODO: examplers
