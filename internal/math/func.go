package math

func GreaterThanZero[T SignedComparable](val T) bool {
	return val > 0
}

func Sum[T Number](val ...T) T {
	var ret T
	for _, v := range val {
		ret += v
	}
	return ret
}

func Less[T Comparable](left, right T) bool {
	return left < right
}

func Greater[T Comparable](left, right T) bool {
	return left > right
}

func Compare[T Comparable](left, right T) int {
	if left == right {
		return 0
	}

	if left > right {
		return 1
	}

	return -1
}

func Min[T Comparable](val ...T) (max T) {
	return MaxFn(Greater[T], val...)
}

func Max[T Comparable](val ...T) (max T) {
	return MaxFn(Less[T], val...)
}

func MinFn[T any](less func(T, T) bool, val ...T) T {
	greater := func(l, r T) bool {
		return !less(r, l)
	}
	return MaxFn(greater, val...)
}

func MaxFn[T any](less func(T, T) bool, val ...T) T {
	var max T
	if len(val) == 0 {
		return max
	}
	max = val[0]
	for i := 1; i < len(val); i++ {
		if less(max, val[i]) {
			max = val[i]
		}
	}

	return max
}
