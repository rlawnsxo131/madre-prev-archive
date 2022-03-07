package lib

var utilManager *utils

type Utils interface {
	IfIsNotExistGetDefaultIntValue(value int, defaultValue int) int
}

type utils struct{}

func GetUtils() Utils {
	once.Do(func() {
		utilManager = &utils{}
	})
	return utilManager
}

// default value
func (u *utils) IfIsNotExistGetDefaultIntValue(value int, defaultValue int) int {
	if value == 0 {
		value = defaultValue
	}
	return value
}
