package telemetry

const (
	SORT_TYPE_NONE         SortType = "none"
	SORT_TYPE_LATENCY_DESC SortType = "latency-desc"
	SORT_TYPE_LATENCY_ASC  SortType = "latency-asc"
)

// SortType is sort type
type SortType string

func (t SortType) IsNone() bool { _ = "STUB: not implemented"; return false }

func (t SortType) IsDesc() bool { _ = "STUB: not implemented"; return false }

func (t SortType) GetHeaderLabel() string { _ = "STUB: not implemented"; return "" }

func sortSvcSpans(svcSpans SvcSpans, sortType SortType) { _ = "STUB: not implemented"; return }

// default sort is received_at asc
