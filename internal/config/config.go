package config

import "github.com/jinzhu/configor"

// Config holds the application config
type Config struct {
	HTTP struct {
		Listen string `default:":8080"`
	}

	Prometheus bool `default:"true"`

	Backend struct {
		Distributor   string
		QueryFrontend string `yaml:"query_frontend"`
		Querier       string
	}

	Authenticator struct {
		Name   string
		Config map[string]string
	}

	Log struct {
		Level string `default:"info"`
	}
}

var config *Config

// Get returns a existing config or creates a new
func Get() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}

// Load loads a config from filesystem
func (c *Config) Load(path string) error {
	return configor.New(&configor.Config{ENVPrefix: "LOKI_AUTH"}).Load(c, path)
}
