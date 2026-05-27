package main

import (
	"log"

	"github.com/spf13/cobra"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/otelcol"

	// Force dependency on main module to ensure it is unambiguous during
	// module resolution.
	// See: https://github.com/googleapis/google-api-go-client/issues/2613.
	// TODO: move to other file such as doc.go ?
	_ "google.golang.org/genproto/googleapis/type/datetime"
)

var version = "unknown"

func main() {
	info := component.BuildInfo{
		Command:     "otel-tui",
		Description: "OpenTelemetry Collector with TUI viewer",
		Version:     version,
	}

	if err := run(otelcol.CollectorSettings{BuildInfo: info, Factories: components}); err != nil {
		log.Fatal(err)
	}
}

func runInteractive(params otelcol.CollectorSettings) error {
	_ = "STUB: not implemented"
	// cmd := otelcol.NewCommand(params)
	return nil
}

type collectorCommand struct {
	*cobra.Command
	params                 otelcol.CollectorSettings
	httpPort               int
	grpcPort               int
	host                   string
	zipkinEnabled          bool
	datadogEnabled         bool
	promTargets            []string
	fromJSONFile           string
	debugLog               bool
	disableInternalMetrics bool
}

func (c *collectorCommand) preRunE(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *collectorCommand) runE(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func newCommand(params otelcol.CollectorSettings) *collectorCommand {
	_ = "STUB: not implemented"
	return nil
}

func setLoggingOptions(params *otelcol.CollectorSettings, debugLogFlag bool) (logPath string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}
