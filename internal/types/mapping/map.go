package mapping

func Make[TK comparable, TV any]() Map[TK, TV] {
	return make(Map[TK, TV])
}

func MakeWithCap[TK comparable, TV any](capacity int) Map[TK, TV] {
	return make(Map[TK, TV], capacity)
}

func Wrap[TK comparable, TV any](m map[TK]TV) Map[TK, TV] {
	return m
}
