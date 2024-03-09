package integration

import (
	"github.com/telenornorway/slf4go"
)

type Driver interface {
	Name() string
	GetLogger() slf4go.Logger
	MdcClear()
	MdcPut(key, value string)
	MdcGet(key string) (string, bool)
	MdcRemove(key string)
	MdcCopy() map[string]string
}
