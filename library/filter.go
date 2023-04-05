package library

func Filter[T any](arr []T, check func(el T) bool) []T {
	new_arr := []T{}
	for _, item := range arr {
		if check(item) {
			new_arr = append(new_arr, item)
		}
	}
	return new_arr
}
