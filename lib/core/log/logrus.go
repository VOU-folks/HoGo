package log

import (
	"github.com/sirupsen/logrus"
)

func Init(level Level, formatter Formatter) {
	SetLevel(level)
	SetFormatter(formatter)
}

const (
	PanicLevel   = "panic"
	FatalLevel   = "fatal"
	ErrorLevel   = "error"
	WarningLevel = "warning"
	InfoLevel    = "info"
	DebugLevel   = "debug"
	TraceLevel   = "trace"
)

type Level string

// SetLevel
// sets minimal verbosity level
// Available level names are:
// "panic"
// "fatal"
// "error"
// "warn" or "warning"
// "info"
// "debug"
// "trace"
func SetLevel(level Level) {
	switch level {
	case PanicLevel:
		logrus.SetLevel(logrus.PanicLevel)
		break

	case FatalLevel:
		logrus.SetLevel(logrus.FatalLevel)
		break

	case ErrorLevel:
		logrus.SetLevel(logrus.ErrorLevel)
		break

	case WarningLevel:
		logrus.SetLevel(logrus.WarnLevel)
		break

	case InfoLevel:
		logrus.SetLevel(logrus.InfoLevel)
		break

	case DebugLevel:
		logrus.SetLevel(logrus.DebugLevel)
		break

	case TraceLevel:
		logrus.SetLevel(logrus.TraceLevel)
		break

	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func Print(args ...interface{}) {
	logrus.Print(args...)
}

func Println(args ...interface{}) {
	logrus.Println(args...)
}

func Info(args ...interface{}) {
	logrus.Infoln(args...)
}

func Warn(args ...interface{}) {
	logrus.Warnln(args...)
}

func Error(args ...interface{}) {
	logrus.Errorln(args...)
}

func Fatal(args ...interface{}) {
	logrus.Fatalln(args...)
}

// Panic
// logs information and exits application
func Panic(args ...interface{}) {
	logrus.Panicln(args...)
}

func Debug(args ...interface{}) {
	logrus.Debugln(args...)
}

func Trace(args ...interface{}) {
	logrus.Traceln(args...)
}

func WithFields(fields map[string]interface{}) *logrus.Entry {
	return logrus.WithFields(fields)
}

var jsonFormatter = &logrus.JSONFormatter{
	DisableTimestamp:  false,
	DisableHTMLEscape: false,
	PrettyPrint:       true,
}

var textFormatter = &logrus.TextFormatter{
	ForceColors:               false,
	DisableColors:             false,
	ForceQuote:                false,
	DisableQuote:              false,
	EnvironmentOverrideColors: false,
	DisableTimestamp:          false,
	FullTimestamp:             false,
	TimestampFormat:           "",
	DisableSorting:            false,
	SortingFunc:               nil,
	DisableLevelTruncation:    false,
	PadLevelText:              false,
	QuoteEmptyFields:          false,
	FieldMap:                  nil,
	CallerPrettyfier:          nil,
}

type Formatter string

const (
	JSONFormatter = "json"
	TextFormatter = "text"
)

func SetFormatter(formatter Formatter) {
	switch formatter {
	case JSONFormatter:
		logrus.SetFormatter(jsonFormatter)
		break

	case TextFormatter:
		logrus.SetFormatter(textFormatter)
		break

	default:
		logrus.SetFormatter(textFormatter)
	}
}

type Fields = logrus.Fields
