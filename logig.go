package logig

import (
	"context"

	"github.com/chiguhagu/logig/logig_default"
)

type Logger interface {
	Debug(fields map[string]interface{}, str string)
	Info(fields map[string]interface{}, str string)
	Warn(fields map[string]interface{}, str string)
	Error(fields map[string]interface{}, str string)
	Fatal(fields map[string]interface{}, str string)

	Debugf(fields map[string]interface{}, format string, args ...interface{})
	Infof(fields map[string]interface{}, format string, args ...interface{})
	Warnf(fields map[string]interface{}, format string, args ...interface{})
	Errorf(fields map[string]interface{}, format string, args ...interface{})
	Fatalf(fields map[string]interface{}, format string, args ...interface{})
}

type logigMaker struct{}

type logig struct {
	logger *Logger
	fields map[string]interface{}
}

var logigKey = &logigMaker{}

func getField(ctx context.Context) map[string]interface{} {
	l, ok := ctx.Value(logigKey).(*logig)
	if !ok || l == nil {
		return map[string]interface{}{}
	}
	return l.fields
}

func extract(ctx context.Context) Logger {
	l, ok := ctx.Value(logigKey).(*logig)
	if !ok || l == nil {
		return logig_default.NewLogigDefault()
	}
	return *l.logger
}

func ToContext(ctx context.Context, logger Logger) context.Context {
	ctx = context.WithValue(ctx, logigKey, &logig{
		logger: &logger,
		fields: map[string]interface{}{},
	})
	return ctx
}

func AddField(ctx context.Context, fields map[string]interface{}) {
	l, ok := ctx.Value(logigKey).(*logig)
	if !ok || l == nil {
		return
	}
	for k, v := range fields {
		l.fields[k] = v
	}
}

func Debug(ctx context.Context, str string) {
	extract(ctx).Debug(getField(ctx), str)
}

func Info(ctx context.Context, str string) {
	extract(ctx).Info(getField(ctx), str)
}

func Warn(ctx context.Context, str string) {
	extract(ctx).Warn(getField(ctx), str)
}

func Error(ctx context.Context, str string) {
	extract(ctx).Error(getField(ctx), str)
}

func Fatal(ctx context.Context, str string) {
	extract(ctx).Fatal(getField(ctx), str)
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	extract(ctx).Debugf(getField(ctx), format, args...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	extract(ctx).Infof(getField(ctx), format, args...)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	extract(ctx).Warnf(getField(ctx), format, args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	extract(ctx).Errorf(getField(ctx), format, args...)
}

func Fatalf(ctx context.Context, format string, args ...interface{}) {
	extract(ctx).Fatalf(getField(ctx), format, args...)
}
