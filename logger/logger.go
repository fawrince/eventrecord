// Package logger is just a wrapper over 'logrus' framework
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

func (logger *Logger) Infof(format string, args ...interface{}) {
	entry := logrus.NewEntry(logger.logger)
	entry.Infof(format, args...)
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	entry := logrus.NewEntry(logger.logger)
	entry.Warnf(format, args...)
}

func (logger *Logger) Error(err error) {
	entry := logrus.NewEntry(logger.logger)
	entry.WithError(err)
	entry.Error()
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	entry := logrus.NewEntry(logger.logger)
	entry.Errorf(format, args...)
}

func (logger *Logger) Fatal(err error) {
	entry := logrus.NewEntry(logger.logger)
	entry.WithError(err)
	entry.Fatal()
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	entry := logrus.NewEntry(logger.logger)
	entry.Fatalf(format, args...)
}

func (logger *Logger) Tracef(format string, args ...interface{}) {
	entry := logrus.NewEntry(logger.logger)
	entry.Tracef(format, args...)
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	entry := logrus.NewEntry(logger.logger)
	entry.Debugf(format, args...)
}
