package reolink

import (
	"bytes"
	"strings"
	"testing"
)

func TestNoOpLogger(t *testing.T) {
	logger := &NoOpLogger{}

	// Should not panic
	logger.Debug("test")
	logger.Info("test")
	logger.Warn("test")
	logger.Error("test")

	// With args
	logger.Debug("test %s", "arg")
	logger.Info("test %s", "arg")
	logger.Warn("test %s", "arg")
	logger.Error("test %s", "arg")
}

func TestStdLogger(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := NewStdLogger(buf)

	logger.Debug("debug message")
	if !strings.Contains(buf.String(), "[DEBUG]") {
		t.Error("expected [DEBUG] prefix")
	}
	if !strings.Contains(buf.String(), "debug message") {
		t.Error("expected debug message")
	}

	buf.Reset()
	logger.Info("info message")
	if !strings.Contains(buf.String(), "[INFO]") {
		t.Error("expected [INFO] prefix")
	}
	if !strings.Contains(buf.String(), "info message") {
		t.Error("expected info message")
	}

	buf.Reset()
	logger.Warn("warn message")
	if !strings.Contains(buf.String(), "[WARN]") {
		t.Error("expected [WARN] prefix")
	}
	if !strings.Contains(buf.String(), "warn message") {
		t.Error("expected warn message")
	}

	buf.Reset()
	logger.Error("error message")
	if !strings.Contains(buf.String(), "[ERROR]") {
		t.Error("expected [ERROR] prefix")
	}
	if !strings.Contains(buf.String(), "error message") {
		t.Error("expected error message")
	}
}

func TestStdLoggerWithArgs(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := NewStdLogger(buf)

	logger.Debug("debug %s %d", "test", 123)
	if !strings.Contains(buf.String(), "debug test 123") {
		t.Errorf("expected formatted message, got: %s", buf.String())
	}

	buf.Reset()
	logger.Info("info %s", "test")
	if !strings.Contains(buf.String(), "info test") {
		t.Errorf("expected formatted message, got: %s", buf.String())
	}
}

func TestStdLoggerNilWriter(t *testing.T) {
	// Should not panic with nil writer (defaults to os.Stderr)
	logger := NewStdLogger(nil)
	if logger == nil {
		t.Error("expected non-nil logger")
	}

	// Should not panic when logging
	logger.Debug("test")
	logger.Info("test")
	logger.Warn("test")
	logger.Error("test")
}

func TestLogLevelString(t *testing.T) {
	tests := []struct {
		level    LogLevel
		expected string
	}{
		{LogLevelDebug, "DEBUG"},
		{LogLevelInfo, "INFO"},
		{LogLevelWarn, "WARN"},
		{LogLevelError, "ERROR"},
		{LogLevelNone, "NONE"},
		{LogLevel(999), "UNKNOWN(999)"},
	}

	for _, tt := range tests {
		if got := tt.level.String(); got != tt.expected {
			t.Errorf("LogLevel(%d).String() = %s, want %s", tt.level, got, tt.expected)
		}
	}
}

func TestLevelLogger(t *testing.T) {
	buf := &bytes.Buffer{}
	stdLogger := NewStdLogger(buf)

	tests := []struct {
		name          string
		level         LogLevel
		logFunc       func(Logger)
		shouldContain string
		shouldBeEmpty bool
	}{
		{
			name:          "debug level logs debug",
			level:         LogLevelDebug,
			logFunc:       func(l Logger) { l.Debug("debug") },
			shouldContain: "debug",
		},
		{
			name:          "info level skips debug",
			level:         LogLevelInfo,
			logFunc:       func(l Logger) { l.Debug("debug") },
			shouldBeEmpty: true,
		},
		{
			name:          "info level logs info",
			level:         LogLevelInfo,
			logFunc:       func(l Logger) { l.Info("info") },
			shouldContain: "info",
		},
		{
			name:          "warn level skips info",
			level:         LogLevelWarn,
			logFunc:       func(l Logger) { l.Info("info") },
			shouldBeEmpty: true,
		},
		{
			name:          "warn level logs warn",
			level:         LogLevelWarn,
			logFunc:       func(l Logger) { l.Warn("warn") },
			shouldContain: "warn",
		},
		{
			name:          "error level skips warn",
			level:         LogLevelError,
			logFunc:       func(l Logger) { l.Warn("warn") },
			shouldBeEmpty: true,
		},
		{
			name:          "error level logs error",
			level:         LogLevelError,
			logFunc:       func(l Logger) { l.Error("error") },
			shouldContain: "error",
		},
		{
			name:          "none level skips error",
			level:         LogLevelNone,
			logFunc:       func(l Logger) { l.Error("error") },
			shouldBeEmpty: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.Reset()
			levelLogger := NewLevelLogger(tt.level, stdLogger)
			tt.logFunc(levelLogger)

			output := buf.String()
			if tt.shouldBeEmpty {
				if output != "" {
					t.Errorf("expected empty output, got: %s", output)
				}
			} else if !strings.Contains(output, tt.shouldContain) {
				t.Errorf("expected output to contain %q, got: %s", tt.shouldContain, output)
			}
		})
	}
}

func TestLevelLoggerWithArgs(t *testing.T) {
	buf := &bytes.Buffer{}
	stdLogger := NewStdLogger(buf)
	levelLogger := NewLevelLogger(LogLevelDebug, stdLogger)

	levelLogger.Debug("debug %s %d", "test", 123)
	if !strings.Contains(buf.String(), "debug test 123") {
		t.Errorf("expected formatted message, got: %s", buf.String())
	}

	buf.Reset()
	levelLogger.Info("info %s", "test")
	if !strings.Contains(buf.String(), "info test") {
		t.Errorf("expected formatted message, got: %s", buf.String())
	}
}
