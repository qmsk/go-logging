package logging

import (
	"testing"
)

type testLogger struct {
	t      *testing.T
	prefix string
}

func (logger testLogger) Printf(format string, args ...interface{}) {
	logger.t.Logf(logger.prefix+format, args...)
}

func TestLogging(t *testing.T) Logging {
	return Logging{
		Debug: testLogger{t, "DEBUG: "},
		Info:  testLogger{t, "INFO: "},
		Warn:  testLogger{t, "WARN: "},
		Error: testLogger{t, "ERROR: "},
	}
}
