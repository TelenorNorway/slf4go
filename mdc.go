package slf4go

func MdcPut(key, value string) {
	configuredDriver.MdcPut(key, value)
}
func MdcGet(key string) (string, bool) {
	return configuredDriver.MdcGet(key)
}
func MdcRemove(key string) {
	configuredDriver.MdcRemove(key)
}
func MdcClear() {
	configuredDriver.MdcClear()
}
func MdcCopy() map[string]string {
	return configuredDriver.MdcCopy()
}
