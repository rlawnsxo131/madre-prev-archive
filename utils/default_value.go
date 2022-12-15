package utils

func Contains[T int | string](ss []T, value T) bool {
	for _, s := range ss {
		if s == value {
			return true
		}
	}
	return false
}

func IfIsNotExistGetDefaultIntValue(value int, defaultValue int) int {
	if value == 0 {
		value = defaultValue
	}
	return value
}

func ParseOtionalString(ss ...string) string {
	var result string
	if len(ss) > 0 {
		for _, v := range ss {
			result += v
		}
	}
	return result
}
