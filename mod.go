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

var installedDriver Driver = nil

func UseDriver(driver Driver) {
	if installedDriver != nil {
		panic(fmt.Sprintf("Driver already installed: %s", installedDriver.Name()))
	}
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
