package slice

func CopyOf[T any](slice []T) []T {
	cpy := make([]T, 0, cap(slice))
	copy(cpy, slice)
	return cpy
}
