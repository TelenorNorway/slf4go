package slf4go

import "github.com/telenornorway/slf4go/internal/caller"

type Logger struct {
	name string
}

func GetLogger() *Logger {
	return &Logger{name: caller.Get(1).Sig.Package}
}

// Name returns the name of the logger
func (l *Logger) Name() string {
	return l.name
}

// Level returns the log level of the logger
func (l *Logger) Level() LogLevel {
	return configuredDriver.GetLevelForLoggerWithName(l.name)
}

func (l *Logger) output(level LogLevel, format string, args ...any) {
	configuredDriver.Write(newLogPayload(l.name, level, format, args))
}

func (l *Logger) isEnabled(level LogLevel) bool {
	return l.Level() >= level
}

// Trace logs a message at the trace level
func (l *Logger) Trace(format string, args ...any) {
	l.output(Trace, format, args...)
}

// TraceIf logs a message at the trace level if the expression is true
func (l *Logger) TraceIf(expression bool, format string, args ...any) {
	if expression {
		l.output(Trace, format, args...)
	}
}

// TraceUnless logs a message at the trace level if the expression is false
func (l *Logger) TraceUnless(expression bool, format string, args ...any) {
	if !expression {
		l.output(Trace, format, args...)
	}
}

// IsTraceEnabled returns true if the trace level is enabled
func (l *Logger) IsTraceEnabled() bool {
	return l.isEnabled(Trace)
}

// Debug logs a message at the debug level
func (l *Logger) Debug(format string, args ...any) {
	l.output(Debug, format, args...)
}

// DebugIf logs a message at the debug level if the expression is true
func (l *Logger) DebugIf(expression bool, format string, args ...any) {
	if expression {
		l.output(Debug, format, args...)
	}
}

// DebugUnless logs a message at the debug level if the expression is false
func (l *Logger) DebugUnless(expression bool, format string, args ...any) {
	if !expression {
		l.output(Debug, format, args...)
	}
}

// IsDebugEnabled returns true if the debug level is enabled
func (l *Logger) IsDebugEnabled() bool {
	return l.isEnabled(Debug)
}

// Info logs a message at the info level
func (l *Logger) Info(format string, args ...any) {
	l.output(Info, format, args...)
}

// InfoIf logs a message at the info level if the expression is true
func (l *Logger) InfoIf(expression bool, format string, args ...any) {
	if expression {
		l.output(Info, format, args...)
	}
}

// InfoUnless logs a message at the info level if the expression is false
func (l *Logger) InfoUnless(expression bool, format string, args ...any) {
	if !expression {
		l.output(Info, format, args...)
	}
}

// IsInfoEnabled returns true if the info level is enabled
func (l *Logger) IsInfoEnabled() bool {
	return l.isEnabled(Info)
}

// Warn logs a message at the warn level
func (l *Logger) Warn(format string, args ...any) {
	l.output(Warn, format, args...)
}

// WarnIf logs a message at the warn level if the expression is true
func (l *Logger) WarnIf(expression bool, format string, args ...any) {
	if expression {
		l.output(Warn, format, args...)
	}
}

// WarnUnless logs a message at the warn level if the expression is false
func (l *Logger) WarnUnless(expression bool, format string, args ...any) {
	if !expression {
		l.output(Warn, format, args...)
	}
}

// IsWarnEnabled returns true if the warn level is enabled
func (l *Logger) IsWarnEnabled() bool {
	return l.isEnabled(Warn)
}

// Error logs a message at the error level
func (l *Logger) Error(format string, args ...any) {
	l.output(Error, format, args...)
}

// ErrorIf logs a message at the error level if the expression is true
func (l *Logger) ErrorIf(expression bool, format string, args ...any) {
	if expression {
		l.output(Error, format, args...)
	}
}

// ErrorUnless logs a message at the error level if the expression is false
func (l *Logger) ErrorUnless(expression bool, format string, args ...any) {
	if !expression {
		l.output(Error, format, args...)
	}
}

// IsErrorEnabled returns true if the error level is enabled
func (l *Logger) IsErrorEnabled() bool {
	return l.isEnabled(Error)
}

// Fatal logs a message at the fatal level
func (l *Logger) Fatal(format string, args ...any) {
	l.output(Fatal, format, args...)
}

// FatalIf logs a message at the fatal level if the expression is true
func (l *Logger) FatalIf(expression bool, format string, args ...any) {
	if expression {
		l.output(Fatal, format, args...)
	}
}

// FatalUnless logs a message at the fatal level if the expression is false
func (l *Logger) FatalUnless(expression bool, format string, args ...any) {
	if !expression {
		l.output(Fatal, format, args...)
	}
}

// IsFatalEnabled returns true if the fatal level is enabled
func (l *Logger) IsFatalEnabled() bool {
	return l.isEnabled(Fatal)
}
