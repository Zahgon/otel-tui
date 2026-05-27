package main

import (
	_ "embed"
)

//go:embed config.yml.tpl
var configYmlTpl string

type Param struct {
	Key    string
	Values []string
}

type PromScrapeConfig struct {
	JobName     string
	Scheme      string
	MetricsPath string
	Target      string
	Params      []Param
}

type Config struct {
	OTLPHost               string
	OTLPHTTPPort           int
	OTLPGRPCPort           int
	EnableZipkin           bool
	EnableDatadog          bool
	FromJSONFile           string
	PromTarget             []string
	PromScrapeConfigs      []*PromScrapeConfig
	DebugLogFilePath       string
	DisableInternalMetrics bool
	AuthToken              string // #nosec G117
}

func NewConfig(
	otlpHost string,
	otlpHTTPPort int,
	otlpGRPCPort int,
	enableZipkin bool,
	enableDatadog bool,
	fromJSONFile string,
	promTarget []string,
	debugLogFilePath string,
	disableInternalMetrics bool,
	authToken string,
) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// buildPromScrapeConfigs parses PromTarget entries and builds PromScrapeConfig objects
func (c *Config) buildPromScrapeConfigs() error { _ = "STUB: not implemented"; return nil }

func (c *Config) RenderYml() (string, error) { _ = "STUB: not implemented"; return "", nil }

func structToMap(s any) (map[string]any, error) { _ = "STUB: not implemented"; return nil, nil }

// validate checks if the otel-tui configuration is valid
func (c *Config) validate() error { _ = "STUB: not implemented"; return nil }
