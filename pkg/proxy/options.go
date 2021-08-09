package proxy

import "github.com/rs/zerolog"

// Option function that enables a feature on Proxy
type Option func(proxy *Proxy)

// WithPrometheus enables prometheus metric endpoint
func WithPrometheus() Option {
	return func(proxy *Proxy) {
		proxy.Prometheus = true
	}
}

// WithLogger adds a logger to log requests
func WithLogger(logger *zerolog.Logger) Option {
	return func(proxy *Proxy) {
		proxy.logger = logger
	}
}
