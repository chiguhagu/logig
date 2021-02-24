package logig_zap

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogigZap struct {
	*zap.Logger
}

func NewLogigZap(logLevel string) *LogigZap {
	c := zap.NewProductionConfig()
	level := zapcore.InfoLevel
	err := level.Set(logLevel)
	if err != nil {
		return &LogigZap{zap.NewNop()}
	}
	c.Level = zap.NewAtomicLevelAt(level)
	c.DisableCaller = true
	c.EncoderConfig.TimeKey = "time"
	c.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	c.DisableStacktrace = true
	l, err := c.Build()
	if err != nil {
		return &LogigZap{zap.NewNop()}
	}
	return &LogigZap{l}
}

type zapFields map[string]interface{}

func (zf zapFields) zapOption() zap.Option {
	f := make([]zap.Field, len(zf))
	var c int
	for k, v := range zf {
		f[c] = zap.Any(k, v)
		c++
	}
	return zap.Fields(f...)
}

func (l *LogigZap) Debug(fields map[string]interface{}, str string) {
	l.Logger.WithOptions(zapFields(fields).zapOption()).Debug(str)
}

func (l *LogigZap) Info(fields map[string]interface{}, str string) {
	l.Logger.WithOptions(zapFields(fields).zapOption()).Info(str)
}

func (l *LogigZap) Warn(fields map[string]interface{}, str string) {
	l.Logger.WithOptions(zapFields(fields).zapOption()).Warn(str)
}

func (l *LogigZap) Error(fields map[string]interface{}, str string) {
	l.Logger.WithOptions(zapFields(fields).zapOption()).Error(str)
}

func (l *LogigZap) Fatal(fields map[string]interface{}, str string) {
	l.Logger.WithOptions(zapFields(fields).zapOption()).Fatal(str)
}

func (l *LogigZap) Debugf(fields map[string]interface{}, format string, args ...interface{}) {
	l.Logger.WithOptions(zapFields(fields).zapOption()).Debug(fmt.Sprintf(format, args...))
}

func (l *LogigZap) Infof(fields map[string]interface{}, format string, args ...interface{}) {
	l.Logger.WithOptions(zapFields(fields).zapOption()).Info(fmt.Sprintf(format, args...))
}

func (l *LogigZap) Warnf(fields map[string]interface{}, format string, args ...interface{}) {
	l.Logger.WithOptions(zapFields(fields).zapOption()).Warn(fmt.Sprintf(format, args...))
}

func (l *LogigZap) Errorf(fields map[string]interface{}, format string, args ...interface{}) {
	l.Logger.WithOptions(zapFields(fields).zapOption()).Error(fmt.Sprintf(format, args...))
}

func (l *LogigZap) Fatalf(fields map[string]interface{}, format string, args ...interface{}) {
	l.Logger.WithOptions(zapFields(fields).zapOption()).Fatal(fmt.Sprintf(format, args...))
}
