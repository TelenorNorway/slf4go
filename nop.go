package slf4go

type NopDriver struct{}

func (d NopDriver) Name() string                                { return "nop" }
func (d NopDriver) GetLevelForLoggerWithName(_ string) LogLevel { return 6 }
func (d NopDriver) Write(_ LogPayload)                          {}
func (d NopDriver) MdcPut(_, _ string)                          {}
func (d NopDriver) MdcGet(_ string) (string, bool)              { return "", false }
func (d NopDriver) MdcRemove(_ string)                          {}
func (d NopDriver) MdcClear()                                   {}
func (d NopDriver) MdcCopy() map[string]string                  { return nil }
