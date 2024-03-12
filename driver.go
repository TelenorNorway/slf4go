package slf4go

import (
	"fmt"
)

type Driver interface {
	Name() string
	GetLevelForLoggerWithName(name string) LogLevel
	Write(payload LogPayload)
	MdcPut(key, value string)
	MdcGet(key string) (string, bool)
	MdcRemove(key string)
	MdcClear()
	MdcCopy() map[string]string
}

var configuredDriver Driver = NopDriver{}
var isConfigured bool

func UseDriver(driver Driver) {
	if isConfigured {
		panic(fmt.Sprintf("Cannot install driver '%s', driver '%s' is already installed", driver.Name(), configuredDriver.Name()))
	}
	isConfigured = true
	configuredDriver = driver
}
