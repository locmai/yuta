package common

import (
	"fmt"

	jaegerconfig "github.com/uber/jaeger-client-go/config"
)

type CommonConfig struct {
	Server    ServerConfig `yaml:"server"`
	JetStream JetStream    `yaml:"jetstream"`
	Metrics   Metrics      `yaml:"metrics"`
	Tracing   struct {
		// Set to true to enable tracer hooks. If false, no tracing is set up.
		Enabled bool `yaml:"enabled"`
		// The config for the jaeger opentracing reporter.
		Jaeger jaegerconfig.Configuration `yaml:"jaeger"`
	} `yaml:"tracing"`
	Logging []LogrusHook `yaml:"logging"`
}

type ServerConfig struct {
	// The host which the server run on
	Host string `yaml:"host"`
	// The port which the server run on
	Port string `yaml:"port"`
	// Timeout for both read and write operations
	Timeout int `yaml:"timeout"`
}

type JetStream struct {
	// A list of NATS addresses to connect to.
	Addresses   []string `yaml:"addresses"`
	TopicPrefix string   `yaml:"topic_prefix"`
}

func (c *JetStream) TopicFor(name string) string {
	return fmt.Sprintf("%s%s", c.TopicPrefix, name)
}

func (c *JetStream) Durable(name string) string {
	return c.TopicFor(name)
}

func (c *JetStream) Defaults(generate bool) {
	c.Addresses = []string{"0.0.0.0:4222"}
	c.TopicPrefix = "Yuta"
}

type Metrics struct {
	// Whether or not the metrics are enabled
	Enabled bool `yaml:"enabled"`
	// Use BasicAuth for Authorization
}

// LogrusHook represents a single logrus hook. At this point, only parsing and
// verification of the proper values for type and level are done.
// Validity/integrity checks on the parameters are done when configuring logrus.
type LogrusHook struct {
	// The type of hook, currently only "file" is supported.
	Type string `yaml:"type"`

	// The level of the logs to produce. Will output only this level and above.
	Level string `yaml:"level"`

	// The parameters for this hook.
	Params map[string]interface{} `yaml:"params"`
}
