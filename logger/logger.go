// Package logger is just a wrapper over 3rd-party logger framework
package logger

import (
	"github.com/sirupsen/logrus"
	"io"
)

type Logger struct {
	logger *logrus.Logger
}

func NewLogger(output io.Writer) *Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(output)
	log.SetLevel(6)

	return &Logger{
		logger: log,
	}
}

// mapEntry maps passed arguments as key1, value1 ... keyN, valueN pairs.
// if an odd number of arguments are passed, the last one will be interpreted as a message
func (logger *Logger) mapEntry(args []interface{}) (*logrus.Entry, string) {
	entry := logrus.NewEntry(logger.logger)
	if len(args) > 1 {
		for i := 0; i < len(args)/2; i = i + 2 {
			entry = entry.WithField(args[i].(string), args[i+1])
		}
	}
	var message string
	if len(args)%2 == 1 {
		var val = args[len(args)-1]
		switch val.(type) {
		case string:
			message = val.(string)
		case error:
			entry = entry.WithError(val.(error))
		}
	}
	return entry, message
}

func (logger *Logger) Info(args ...interface{}) {
	entry, message := logger.mapEntry(args)
	entry.Info(message)
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	entry := logrus.NewEntry(logger.logger)
	entry.Infof(format, args...)
}

func (logger *Logger) Warn(args ...interface{}) {
	entry, message := logger.mapEntry(args)
	entry.Warn(message)
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	entry := logrus.NewEntry(logger.logger)
	entry.Warnf(format, args...)
}

func (logger *Logger) Error(args ...interface{}) {
	entry, message := logger.mapEntry(args)
	entry.Error(message)
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	entry := logrus.NewEntry(logger.logger)
	entry.Errorf(format, args...)
}

func (logger *Logger) Fatal(args ...interface{}) {
	entry, message := logger.mapEntry(args)
	entry.Fatal(message)
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	entry := logrus.NewEntry(logger.logger)
	entry.Fatalf(format, args...)
}

func (logger *Logger) Trace(args ...interface{}) {
	entry, message := logger.mapEntry(args)
	entry.Trace(message)
}

func (logger *Logger) Tracef(format string, args ...interface{}) {
	entry := logrus.NewEntry(logger.logger)
	entry.Tracef(format, args...)
}

func (logger *Logger) Debug(args ...interface{}) {
	entry, message := logger.mapEntry(args)
	entry.Debug(message)
}

func (logger *Logger) Debugf(format string,args ...interface{}) {
	entry := logrus.NewEntry(logger.logger)
	entry.Debugf(format, args...)
}
