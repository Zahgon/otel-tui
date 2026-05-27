package telemetry

func (m SpanDataMap) getDependencyGraph() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (m SpanDataMap) getDependencies() *dependencyInfo {
	_ = "STUB: not implemented"
	// TODO: should we take an exclusive lock?
	return nil
}

// service ID to node map

// new dependency

// The nodes that do not have a parent are the head nodes

type dependencyInfo struct {
	HeadNodes  []*node
	CallCounts map[string]int
}

func (d *dependencyInfo) getMermaid() string { _ = "STUB: not implemented"; return "" }

// Start new line

type node struct {
	Service  string
	Parent   *node
	Children []*node
	Depth    int
}

func (n *node) updateDepth() int {
	_ = "STUB: not implemented"

	// If depth is already calculated, return it
	return 0
}

func getSortedMermaid(input string) string { _ = "STUB: not implemented"; return "" }

func getDepKey(parentsn, childsn string) string { _ = "STUB: not implemented"; return "" }
