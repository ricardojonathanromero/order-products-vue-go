package logger

import (
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/logger/formatter"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

const (
	app = "app-name"
)

type Tag map[string]string

type Opts struct {
	LogLevel  string
	AppName   string
	BlackList []string
	Tags      Tag
}

func getLogLevel(lvl string) logrus.Level {
	switch strings.ToLower(lvl) {
	case "info":
		return logrus.InfoLevel
	case "debug":
		return logrus.DebugLevel
	case "error":
		return logrus.ErrorLevel
	case "trace":
		return logrus.TraceLevel
	case "warning", "warn":
		return logrus.WarnLevel
	default:
		return logrus.ErrorLevel
	}
}

func New(opts Opts) Logger {
	log := logrus.New()

	log.SetLevel(getLogLevel(opts.LogLevel))
	log.SetOutput(os.Stdout)
	log.SetFormatter(formatter.NewFormatter(opts.BlackList...))

	e := log.WithField(app, opts.AppName)

	for k, v := range opts.Tags {
		e = e.WithField(k, v)
	}

	return &loggerImpl{entry: e}
}

// Logger defines the interface for logging operations
type Logger interface {
	AddFields(fields map[string]interface{})
	Error(msg string)
	Errorf(msg string, args ...interface{})
	Debug(msg string)
	Debugf(msg string, args ...interface{})
	Info(msg string)
	Trace(msg string)
	Warn(msg string)
	Warnf(msg string, args ...interface{})
	Fatalf(msg string, args ...interface{})
	Writer() io.Writer
}

type loggerImpl struct {
	entry *logrus.Entry
}

func (l *loggerImpl) AddFields(fields map[string]interface{}) {
	for field, value := range fields {
		l.entry = l.entry.WithField(field, value)
	}
}

func (l *loggerImpl) Error(msg string) {
	l.entry.Error(msg)
}

func (l *loggerImpl) Errorf(msg string, args ...interface{}) {
	l.entry.Errorf(msg, args...)
}

func (l *loggerImpl) Debug(msg string) {
	l.entry.Debug(msg)
}

func (l *loggerImpl) Debugf(msg string, args ...interface{}) {
	l.entry.Debugf(msg, args...)
}

func (l *loggerImpl) Info(msg string) {
	l.entry.Info(msg)
}

func (l *loggerImpl) Trace(msg string) {
	l.entry.Trace(msg)
}

func (l *loggerImpl) Warn(msg string) {
	l.entry.Warn(msg)
}

func (l *loggerImpl) Warnf(msg string, args ...interface{}) {
	l.entry.Warnf(msg, args...)
}

func (l *loggerImpl) Fatalf(msg string, args ...interface{}) {
	l.entry.Fatalf(msg, args...)
}

func (l *loggerImpl) Writer() io.Writer {
	return l.entry.Writer()
}
