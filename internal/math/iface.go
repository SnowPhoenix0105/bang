package math

type SignedComparable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

type UnsignedComparable interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | uintptr
}

type Comparable interface {
	SignedComparable | UnsignedComparable
}

type Number interface {
	Comparable | complex64 | complex128
}
