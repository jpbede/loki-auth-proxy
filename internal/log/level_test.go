package log_test

import (
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"go.bnck.me/loki-auth-proxy/internal/log"
	"testing"
)

func TestDecodeLogLevel(t *testing.T) {
	assert.Equal(t, zerolog.InfoLevel, log.DecodeLogLevel("bla"), "Unknown level should match InfoLevel")
	assert.Equal(t, zerolog.DebugLevel, log.DecodeLogLevel("debug"), "String debug does not match DebugLevel")
	assert.Equal(t, zerolog.InfoLevel, log.DecodeLogLevel("info"), "String info does not match InfoLevel")
	assert.Equal(t, zerolog.WarnLevel, log.DecodeLogLevel("warn"), "String warn does not match WarnLevel")
	assert.Equal(t, zerolog.ErrorLevel, log.DecodeLogLevel("error"), "String error does not match ErrorLevel")
	assert.Equal(t, zerolog.FatalLevel, log.DecodeLogLevel("fatal"), "String fatal does not match FatalLevel")
}
