// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package sharedcomponent exposes functionality for components
// to register against a shared key, such as a configuration object, in order to be reused across signal types.
// This is particularly useful when the component relies on a shared resource such as os.File or http.Server.
package sharedcomponent // import "go.opentelemetry.io/collector/internal/sharedcomponent"

import (
	"context"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componentstatus"
)

func NewMap[K comparable, V component.Component]() *Map[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Map keeps reference of all created instances for a given shared key such as a component configuration.
type Map[K comparable, V component.Component] struct {
	lock       sync.Mutex
	components map[K]*Component[V]
}

// LoadOrStore returns the already created instance if exists, otherwise creates a new instance
// and adds it to the map of references.
func (m *Map[K, V]) LoadOrStore(key K, create func() (V, error), telemetrySettings *component.TelemetrySettings) (*Component[V], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Component ensures that the wrapped component is started and stopped only once.
// When stopped it is removed from the Map.
type Component[V component.Component] struct {
	component V

	startOnce  sync.Once
	stopOnce   sync.Once
	removeFunc func()

	telemetry *component.TelemetrySettings

	hostWrapper *hostWrapper
}

// Unwrap returns the original component.
func (c *Component[V]) Unwrap() V {
	_ = "STUB: not implemented"

	// Start starts the underlying component if it never started before.
	return *new(V)
}

func (c *Component[V]) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// It's important that status for a shared component is reported through its
// telemetry settings to keep status in sync and avoid race conditions. This logic duplicates
// and takes priority over the automated status reporting that happens in graph, making the
// status reporting in graph a no-op.

var _ component.Host = (*hostWrapper)(nil)
var _ componentstatus.Reporter = (*hostWrapper)(nil)

type hostWrapper struct {
	host           component.Host
	sources        []componentstatus.Reporter
	previousEvents []*componentstatus.Event
	lock           sync.Mutex
}

func (h *hostWrapper) GetExtensions() map[component.ID]component.Component {
	_ = "STUB: not implemented"
	return nil
}

func (h *hostWrapper) Report(e *componentstatus.Event) {
	_ = "STUB: not implemented"
	// Only remember an event if it will be emitted and it has not been sent already.
	return
}

func (h *hostWrapper) addSource(s componentstatus.Reporter) { _ = "STUB: not implemented"; return }

// Shutdown shuts down the underlying component.
func (c *Component[V]) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// It's important that status for a shared component is reported through its
// telemetry settings to keep status in sync and avoid race conditions. This logic duplicates
// and takes priority over the automated status reporting that happens in graph, making the
// status reporting in graph a no-op.
