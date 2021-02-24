package logig_default

import "fmt"

type LogigDefault struct {
}

func NewLogigDefault() *LogigDefault {
	return &LogigDefault{}
}

func (l *LogigDefault) Debug(fields map[string]interface{}, str string) {
	fmt.Println(str)
}

func (l *LogigDefault) Info(fields map[string]interface{}, str string) {
	fmt.Println(str)
}

func (l *LogigDefault) Warn(fields map[string]interface{}, str string) {
	fmt.Println(str)
}

func (l *LogigDefault) Error(fields map[string]interface{}, str string) {
	fmt.Println(str)
}

func (l *LogigDefault) Fatal(fields map[string]interface{}, str string) {
	fmt.Println(str)
}

func (l *LogigDefault) Debugf(fields map[string]interface{}, format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (l *LogigDefault) Infof(fields map[string]interface{}, format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (l *LogigDefault) Warnf(fields map[string]interface{}, format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (l *LogigDefault) Errorf(fields map[string]interface{}, format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (l *LogigDefault) Fatalf(fields map[string]interface{}, format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
