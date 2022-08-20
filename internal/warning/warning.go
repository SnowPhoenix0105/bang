package warning

var globalOnWarning func(msg string)

func SetOnWarning(fn func(msg string)) {
	globalOnWarning = fn
}

func Notify(msg string) {
	globalOnWarning(msg)
}
