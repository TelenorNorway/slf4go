package slf4go

import (
	"github.com/telenornorway/slf4go/internal/caller"
)

type LogPayload struct {
	// Information about the logger

	name string

	// Information about the log entry

	level  LogLevel
	format string
	args   []any
	fields map[string]string

	// Information about the caller

	callerInfo *caller.Info
}

func (p *LogPayload) Name() string              { return p.name }
func (p *LogPayload) Level() LogLevel           { return p.level }
func (p *LogPayload) Format() string            { return p.format }
func (p *LogPayload) Args() []any               { return p.args }
func (p *LogPayload) Fields() map[string]string { return p.fields }
func (p *LogPayload) PackageName() string       { return p.callerInfo.Sig.Package }
func (p *LogPayload) HasReceiver() bool         { return p.callerInfo.Sig.HasReceiver }
func (p *LogPayload) Receiver() string          { return p.callerInfo.Sig.Receiver }
func (p *LogPayload) Function() string          { return p.callerInfo.Sig.Function }
func (p *LogPayload) File() string              { return p.callerInfo.File }
func (p *LogPayload) Line() int                 { return p.callerInfo.Line }

func newLogPayload(
	name string,
	level LogLevel,
	format string,
	args []any,
) LogPayload {
	p := LogPayload{
		name:       name,
		level:      level,
		format:     format,
		args:       args,
		fields:     configuredDriver.MdcCopy(),
		callerInfo: caller.Get(2),
	}
	return p
}
