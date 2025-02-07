package lib

func AppendAll[T any](slices ...[]T) []T {
	appendedSlice := []T{}
	for _, s := range slices {
		appendedSlice = append(appendedSlice, s...)
	}
	return appendedSlice
}
