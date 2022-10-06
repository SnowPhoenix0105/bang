package base

type Iterator[T any] interface {
	Next() bool
	Get() T
}

type Iterable[T any] interface {
	Iterator() Iterator[T]
}

type BackCollection[T any] interface {
	Len() int
	PushBack(elem T)
	PopBack()
	GetBack() T
}

type ReadableKVCollection[TK, TV any] interface {
	Len() int
	TryGet(TK) (TV, bool)
	GetOr(TK, TV) TV
	GetOrDefault(TK) TV
	ContainsKey(TK) bool
}

type ReadWriteKVCollection[TK, TV any] interface {
	ReadableKVCollection[TK, TV]
	Set(TK, TV)
	Merge(TK, TV, func(TV) TV) bool
}
