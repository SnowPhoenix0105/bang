package slice

func Raw[T any](elem ...T) []T {
	return elem
}

func Make[T any]() Slice[T] {
	return []T{}
}

func MakeWithCap[T any](capacity int) Slice[T] {
	return make([]T, 0, capacity)
}

func MakeWithVal[T any](elems ...T) Slice[T] {
	return Wrap(elems)
}

func New[T any]() *Slice[T] {
	return new(Slice[T])
}

func NewWithCap[T any](capacity int) *Slice[T] {
	ret := MakeWithCap[T](capacity)
	return &ret
}

func NewWithVal[T any](elems ...T) *Slice[T] {
	return WrapPtr(&elems)
}

func Wrap[T any](slice []T) Slice[T] {
	return slice
}

func WrapPtr[T any](slice *[]T) *Slice[T] {
	return (*Slice[T])(slice)
}

func Get[T any](slice []T, index int) T {
	return WrapPtr(&slice).Get(index)
}

func TryGet[T any](slice []T, index int) (T, bool) {
	return WrapPtr(&slice).TryGet(index)
}

func GetOrDefault[T any](slice []T, index int) T {
	return WrapPtr(&slice).GetOrDefault(index)
}

func GetOr[T any](slice []T, index int, defaultValue T) T {
	return WrapPtr(&slice).GetOr(index, defaultValue)
}

func Set[T any](slice []T, index int, elem T) {
	WrapPtr(&slice).Set(index, elem)
}

func TrySet[T any](slice []T, index int, elem T) bool {
	return WrapPtr(&slice).TrySet(index, elem)
}

func SetLen[T any](ptr *[]T, length int) bool {
	return WrapPtr(ptr).SetLen(length)
}

func Append[T any](ptr *[]T, elem ...T) {
	WrapPtr(ptr).Append(elem...)
}

func Contains[T any](slice []T, elem T) bool {
	return WrapPtr(&slice).Contains(elem)
}

func Any[T any](slice []T, fn func(T) bool) bool {
	return WrapPtr(&slice).Any(fn)
}

func All[T any](slice []T, fn func(T) bool) bool {
	return WrapPtr(&slice).All(fn)
}

func Find[T any](slice []T, elem T) (int, bool) {
	return WrapPtr(&slice).Find(elem)
}

func FindFn[T any](slice []T, fn func(T) bool) (int, bool) {
	return WrapPtr(&slice).FindFn(fn)
}

func Len[T any](slice []T) int {
	return WrapPtr(&slice).Len()
}

func Cap[T any](slice []T) int {
	return WrapPtr(&slice).Cap()
}

func PushBack[T any](ptr *[]T, elem T) {
	WrapPtr(ptr).PushBack(elem)
}

func PopBack[T any](ptr *[]T) {
	WrapPtr(ptr).PopBack()
}

func GetBack[T any](slice []T) T {
	return WrapPtr(&slice).GetBack()
}

func ForEach[T any](slice []T, fn func(T) bool) {
	WrapPtr(&slice).ForEach(fn)
}

func ForEach2[T any](slice []T, fn func(int, T) bool) {
	WrapPtr(&slice).ForEach2(fn)
}

func ForAll[T any](slice []T, fn func(T)) {
	WrapPtr(&slice).ForAll(fn)
}

func ForAll2[T any](slice []T, fn func(int, T)) {
	WrapPtr(&slice).ForAll2(fn)
}
