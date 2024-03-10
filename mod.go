package slf4go

import (
	"fmt"
)

type LogLevel int

const (
	Fatal LogLevel = iota
	Error
	Warn
	Info
	Debug
	Trace
)

type Logger interface {
	Name() string
	Level() LogLevel
	Trace(format string, args ...any)
	TraceIf(expression bool, format string, args ...any)
	TraceUnless(expression bool, format string, args ...any)
	IsTraceEnabled() bool
	Debug(format string, args ...any)
	DebugIf(expression bool, format string, args ...any)
	DebugUnless(expression bool, format string, args ...any)
	IsDebugEnabled() bool
	Info(format string, args ...any)
	InfoIf(expression bool, format string, args ...any)
	InfoUnless(expression bool, format string, args ...any)
	IsInfoEnabled() bool
	Warn(format string, args ...any)
	WarnIf(expression bool, format string, args ...any)
	WarnUnless(expression bool, format string, args ...any)
	IsWarnEnabled() bool
	Error(format string, args ...any)
	ErrorIf(expression bool, format string, args ...any)
	ErrorUnless(expression bool, format string, args ...any)
	IsErrorEnabled() bool
	Fatal(format string, args ...any)
	FatalIf(expression bool, format string, args ...any)
	FatalUnless(expression bool, format string, args ...any)
	IsFatalEnabled() bool
}

var installedDriver Driver = nop{}
var isInitialized bool

func UseDriver(driver Driver) {
	if isInitialized {
		panic(fmt.Sprintf("Driver already installed: %s", installedDriver.Name()))
	}
	isInitialized = true
	installedDriver = driver
}

type Driver interface {
	Name() string
	GetLogger() Logger
	MdcClear()
	MdcPut(key, value string)
	MdcGet(key string) (string, bool)
	MdcRemove(key string)
	MdcCopy() map[string]string
}

func GetLogger() Logger {
	return installedDriver.GetLogger()
}

func MdcPut(key, value string) {
	installedDriver.MdcPut(key, value)
}

func MdcGet(key string) (string, bool) {
	return installedDriver.MdcGet(key)
}

func MdcRemove(key string) {
	installedDriver.MdcRemove(key)
}

func MdcClear() {
	installedDriver.MdcClear()
}

func MdcCopy() map[string]string {
	return installedDriver.MdcCopy()
}

type nop struct{}
type nopLogger struct{}

func (n nopLogger) Name() string                           { return "nop" }
func (n nopLogger) Level() LogLevel                        { return Fatal }
func (n nopLogger) Trace(_ string, _ ...any)               {}
func (n nopLogger) TraceIf(_ bool, _ string, _ ...any)     {}
func (n nopLogger) TraceUnless(_ bool, _ string, _ ...any) {}
func (n nopLogger) IsTraceEnabled() bool                   { return false }
func (n nopLogger) Debug(_ string, _ ...any)               {}
func (n nopLogger) DebugIf(_ bool, _ string, _ ...any)     {}
func (n nopLogger) DebugUnless(_ bool, _ string, _ ...any) {}
func (n nopLogger) IsDebugEnabled() bool                   { return false }
func (n nopLogger) Info(_ string, _ ...any)                {}
func (n nopLogger) InfoIf(_ bool, _ string, _ ...any)      {}
func (n nopLogger) InfoUnless(_ bool, _ string, _ ...any)  {}
func (n nopLogger) IsInfoEnabled() bool                    { return false }
func (n nopLogger) Warn(_ string, _ ...any)                {}
func (n nopLogger) WarnIf(_ bool, _ string, _ ...any)      {}
func (n nopLogger) WarnUnless(_ bool, _ string, _ ...any)  {}
func (n nopLogger) IsWarnEnabled() bool                    { return false }
func (n nopLogger) Error(_ string, _ ...any)               {}
func (n nopLogger) ErrorIf(_ bool, _ string, _ ...any)     {}
func (n nopLogger) ErrorUnless(_ bool, _ string, _ ...any) {}
func (n nopLogger) IsErrorEnabled() bool                   { return false }
func (n nopLogger) Fatal(_ string, _ ...any)               {}
func (n nopLogger) FatalIf(_ bool, _ string, _ ...any)     {}
func (n nopLogger) FatalUnless(_ bool, _ string, _ ...any) {}
func (n nopLogger) IsFatalEnabled() bool                   { return false }

func (n nop) Name() string                   { return "nop" }
func (n nop) GetLogger() Logger              { return nopLogger{} }
func (n nop) MdcClear()                      {}
func (n nop) MdcPut(_, _ string)             {}
func (n nop) MdcGet(_ string) (string, bool) { return "", false }
func (n nop) MdcRemove(_ string)             {}
func (n nop) MdcCopy() map[string]string     { return map[string]string{} }
