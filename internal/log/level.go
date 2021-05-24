package log

import (
	"github.com/rs/zerolog"
	"strings"
)

// DecodeLogLevel decodes a string loglevel to its constant
func DecodeLogLevel(loglevel string) zerolog.Level {
	switch strings.ToLower(loglevel) {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	default:
		return zerolog.InfoLevel
	}
}
