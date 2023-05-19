package api

import (
	"time"
)

type Configuration struct {
	basePath  string `json:"basePath,omitempty"`
	host      string `json:"host,omitempty"`
	userAgent string `json:"userAgent,omitempty"`
	username  string
	password  string
	timeout   time.Duration
}

// ConfigurationOption configures the Aurora API Client.
type ConfigurationOption func(*Configuration)

func NewConfiguration(opts ...ConfigurationOption) *Configuration {
	cfg := &Configuration{
		basePath:  "https://api.visionular.com",
		userAgent: "Aurora Go | 0.0.1",
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

func WithBasicAuth(username, password string) ConfigurationOption {
	return func(c *Configuration) {
		c.username = username
		c.password = password
	}
}

func WithTimeout(t time.Duration) ConfigurationOption {
	return func(c *Configuration) {
		c.timeout = t
	}
}

func WithBasePath(basePath string) ConfigurationOption {
	return func(c *Configuration) {
		c.basePath = basePath
	}
}
