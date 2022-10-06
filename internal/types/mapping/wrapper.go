package mapping

import (
	"github.com/snowphoenix0105/bang/internal/types/equal"
	"github.com/snowphoenix0105/bang/internal/types/value"
)

type Map[TK comparable, TV any] map[TK]TV

func (m Map[TK, TV]) Raw() map[TK]TV {
	return m
}

func (m Map[TK, TV]) Get(key TK) TV {
	return m[key]
}

func (m Map[TK, TV]) TryGet(key TK) (TV, bool) {
	val, exists := m[key]
	return val, exists
}

func (m Map[TK, TV]) GetOrDefault(key TK) TV {
	return m.Get(key)
}

func (m Map[TK, TV]) GetOr(key TK, defaultValue TV) TV {
	val, exists := m[key]
	if !exists {
		return defaultValue
	}

	return val
}

func (m Map[TK, TV]) Set(key TK, value TV) {
	m[key] = value
}

func (m Map[TK, TV]) Merge(key TK, value TV, merge func(old TV) TV) bool {
	old, exists := m.TryGet(key)
	if exists {
		m.Set(key, merge(old))
		return true
	}

	m.Set(key, value)
	return false
}

func (m Map[TK, TV]) ContainsKey(key TK) bool {
	_, exists := m[key]
	return exists
}

func (m Map[TK, TV]) ContainsValue(val TV) bool {
	return m.Any(func(v TV) bool {
		return equal.Of(v, val)
	})
}

func (m Map[TK, TV]) Contains(key TK) bool {
	return m.ContainsKey(key)
}

func (m Map[TK, TV]) All(fn func(TV) bool) bool {
	return m.All2(func(_ TK, v TV) bool {
		return fn(v)
	})
}

func (m Map[TK, TV]) All2(fn func(TK, TV) bool) bool {
	return !m.Any2(func(k TK, v TV) bool {
		return !fn(k, v)
	})
}

func (m Map[TK, TV]) Any(fn func(TV) bool) bool {
	return m.Any2(func(_ TK, v TV) bool {
		return fn(v)
	})
}

func (m Map[TK, TV]) Any2(fn func(TK, TV) bool) bool {
	for k, v := range m {
		if fn(k, v) {
			return true
		}
	}

	return false
}

func (m Map[TK, TV]) Find(value TV) (TK, bool) {
	return m.FindFn(func(v TV) bool {
		return equal.Of(v, value)
	})
}

func (m Map[TK, TV]) FindFn(fn func(TV) bool) (TK, bool) {
	for k, v := range m {
		if fn(v) {
			return k, true
		}
	}

	return value.Zero[TK](), false
}

func (m Map[TK, TV]) Len() int {
	return len(m)
}

func (m Map[TK, TV]) Cap() int {
	return m.Len()
}

func (m Map[TK, TV]) ForAll(fn func(TV)) {
	m.ForAll2(func(_ TK, v TV) {
		fn(v)
	})
}

func (m Map[TK, TV]) ForAll2(fn func(TK, TV)) {
	m.ForEach2(func(k TK, v TV) bool {
		fn(k, v)
		return true
	})
}

func (m Map[TK, TV]) ForEach(fn func(TV) bool) {
	m.ForEach2(func(_ TK, v TV) bool {
		return fn(v)
	})
}

func (m Map[TK, TV]) ForEach2(fn func(TK, TV) bool) {
	for k, v := range m {
		if !fn(k, v) {
			break
		}
	}
}
