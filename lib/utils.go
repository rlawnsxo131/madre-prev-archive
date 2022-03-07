package lib

type Utils interface {
	IfIsNotExistGetDefaultIntValue(value int, defaultValue int) int
}

type utils struct{}

var utilManager *utils

func NewUtils() Utils {
	if utilManager == nil {
		utilManager = &utils{}
	}
	return utilManager
}

// default value
func (u *utils) IfIsNotExistGetDefaultIntValue(value int, defaultValue int) int {
	if value == 0 {
		value = defaultValue
	}
	return value
}
