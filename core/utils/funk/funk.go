package funk

func Map[S any, R any](
	ss []S,
	f func(v S, idx int, ss []R) R,
) []R {
	result := make([]R, 0, len(ss))

	for idx, v := range ss {
		result = append(result, f(v, idx, result))
	}

	return result
}

func Filter[S any](
	ss []S,
	f func(v S, idx int, ss []S) bool,
) []S {
	result := make([]S, 0, len(ss))

	for idx, v := range ss {
		if f(v, idx, result) {
			result = append(result, v)
		}
	}

	return result
}

func Reduce[S any, R any](
	ss []S,
	f func(acc R, cur S, idx int, ss []S) R,
	initValue R,
) R {
	acc := initValue
	for idx, v := range ss {
		acc = f(acc, v, idx, ss)
	}
	return acc
}

func Contains[T comparable](ss []T, value T) bool {
	for _, s := range ss {
		if s == value {
			return true
		}
	}
	return false
}
