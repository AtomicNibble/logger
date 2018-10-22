package logger

import (
	"github.com/gobuffalo/envy"
	"github.com/sirupsen/logrus"
)

// FieldLogger interface
type FieldLogger interface {
	Logger
	WithField(string, interface{}) Logger
	WithFields(map[string]interface{}) Logger
}

// Logger interface is used throughout Buffalo
// apps to log a whole manner of things.
type Logger interface {
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Printf(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Fatalf(string, ...interface{})
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Fatal(...interface{})
	Panic(...interface{})
}

// NewLogger based on the specified log level, defaults to "debug".
// See `New` for more details.
func NewLogger(level string) Logger {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = logrus.DebugLevel
	}
	return New(lvl)
}

// New based on the specified log level, defaults to "debug".
// This logger will log to the STDOUT in a human readable,
// but parseable form.
/*
	Example: time="2016-12-01T21:02:07-05:00" level=info duration=225.283µs human_size="106 B" method=GET path="/" render=199.79µs request_id=2265736089 size=106 status=200
*/
func New(lvl Level) Logger {
	dev := envy.Get("GO_ENV", "development") == "development"
	l := logrus.New()
	l.Level = lvl
	l.Formatter = &textFormatter{
		ForceColors: dev,
	}
	return Logrus{l}
}
