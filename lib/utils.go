package lib

type utils struct{}

func NewUtils() *utils {
	return &utils{}
}

// default value
func (u *utils) IfIsNotExistGetDefaultIntValue(value int, defaultValue int) int {
	if value == 0 {
		value = defaultValue
	}
	return value
}
