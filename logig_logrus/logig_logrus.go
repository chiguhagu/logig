package logig_logrus

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type LogigLogrus struct {
	logrus.Entry
}

func NewLogigLogrus(logLevel string) *LogigLogrus {
	l, err := logrus.ParseLevel(logLevel)
	if err != nil {
		l = logrus.InfoLevel
	}
	return &LogigLogrus{*logrus.NewEntry(&logrus.Logger{
		Out: os.Stderr,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat: time.RFC3339Nano,
		},
		Hooks:        nil,
		Level:        l,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	})}
}

func (l *LogigLogrus) Debug(fields map[string]interface{}, str string) {
	l.Logger.WithFields(fields).Debug(str)
}

func (l *LogigLogrus) Info(fields map[string]interface{}, str string) {
	l.Logger.WithFields(fields).Info(str)
}

func (l *LogigLogrus) Warn(fields map[string]interface{}, str string) {
	l.Logger.WithFields(fields).Warn(str)
}

func (l *LogigLogrus) Error(fields map[string]interface{}, str string) {
	l.Logger.WithFields(fields).Error(str)
}

func (l *LogigLogrus) Fatal(fields map[string]interface{}, str string) {
	l.Logger.WithFields(fields).Fatal(str)
}

func (l *LogigLogrus) Debugf(fields map[string]interface{}, format string, args ...interface{}) {
	l.Logger.WithFields(fields).Debugf(format, args...)
}

func (l *LogigLogrus) Infof(fields map[string]interface{}, format string, args ...interface{}) {
	l.Logger.WithFields(fields).Infof(format, args...)
}

func (l *LogigLogrus) Warnf(fields map[string]interface{}, format string, args ...interface{}) {
	l.Logger.WithFields(fields).Warnf(format, args...)
}

func (l *LogigLogrus) Errorf(fields map[string]interface{}, format string, args ...interface{}) {
	l.Logger.WithFields(fields).Errorf(format, args...)
}

func (l *LogigLogrus) Fatalf(fields map[string]interface{}, format string, args ...interface{}) {
	l.Logger.WithFields(fields).Fatalf(format, args...)
}
