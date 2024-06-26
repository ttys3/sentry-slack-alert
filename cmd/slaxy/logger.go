package main

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/innogames/slaxy"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func init() {
	// Log as JSON instead of the default ASCII formatter.
	if os.Getenv("SLAXY_LOG_FORMAT") == "json" {
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05Z07:00",
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				filename := path.Base(f.File)
				return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
			},
		})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05Z07:00",
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				filename := path.Base(f.File)
				return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
			},
		})
	}

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logger.SetOutput(os.Stdout)

	logger.SetReportCaller(true)

	level := new(logrus.Level)
	if err := level.UnmarshalText([]byte(os.Getenv("SLAXY_LOG_LEVEL"))); err != nil {
		*level = logrus.InfoLevel
		logger.Infof("using default log level=%v", *level)
		logger.SetLevel(*level)
	} else {
		logger.Infof("set log level=%v", *level)
		logger.SetLevel(*level)
	}
}

// logrusLogger wraps a logrus logger for compatibility with the slaxy library
type logrusLogger struct {
	slaxy.Logger
	l *logrus.Logger
}

// Debug logs debug messages
func (l *logrusLogger) Debug(args ...interface{}) {
	l.l.Debug(args...)
}

// Debugf logs debug messages
func (l *logrusLogger) Debugf(msg string, args ...interface{}) {
	l.l.Debugf(msg, args...)
}

// Info logs debug messages
func (l *logrusLogger) Info(args ...interface{}) {
	l.l.Info(args...)
}

// Infof logs debug messages
func (l *logrusLogger) Infof(msg string, args ...interface{}) {
	l.l.Infof(msg, args...)
}

// Warn logs debug messages
func (l *logrusLogger) Warn(args ...interface{}) {
	l.l.Warn(args...)
}

// Warnf logs debug messages
func (l *logrusLogger) Warnf(msg string, args ...interface{}) {
	l.l.Warnf(msg, args...)
}

// Error logs debug messages
func (l *logrusLogger) Error(args ...interface{}) {
	l.l.Error(args...)
}

// Errorf logs debug messages
func (l *logrusLogger) Errorf(msg string, args ...interface{}) {
	l.l.Errorf(msg, args...)
}
