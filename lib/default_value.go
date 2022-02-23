package lib

func IfIsNotExistGetDefaultIntValue(value int, defaultValue int) int {
	if value == 0 {
		value = defaultValue
	}
	return value
}
