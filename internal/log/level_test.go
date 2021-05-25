package log_test

import (
	"github.com/rs/zerolog"
	"go.bnck.me/loki-auth-proxy/internal/log"
	"testing"
)

func TestDecodeLogLevel(t *testing.T) {
	bla := log.DecodeLogLevel("bla")
	if bla != zerolog.InfoLevel {
		t.Error("Unknown level should match InfoLevel")
	}

	debug := log.DecodeLogLevel("debug")
	if debug != zerolog.DebugLevel {
		t.Error("String debug does not match DebugLevel")
	}

	info := log.DecodeLogLevel("info")
	if info != zerolog.InfoLevel {
		t.Error("String info does not match InfoLevel")
	}

	warn := log.DecodeLogLevel("warn")
	if warn != zerolog.WarnLevel {
		t.Error("String warn does not match WarnLevel")
	}

	error := log.DecodeLogLevel("error")
	if error != zerolog.ErrorLevel {
		t.Error("String error does not match ErrorLevel")
	}

	fatal := log.DecodeLogLevel("fatal")
	if fatal != zerolog.FatalLevel {
		t.Error("String fatal does not match FatalLevel")
	}
}
