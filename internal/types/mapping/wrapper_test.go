package mapping

import (
	"testing"

	"github.com/snowphoenix0105/bang/internal/types/value"
	"github.com/stretchr/testify/assert"
)

func TestWrap(t *testing.T) {
	raw := map[string]interface{}{
		"A": 1,
		"B": map[string]int{},
	}
	w := Wrap(raw)
	assert.Equal(t, raw, w.Raw())

	assert.Equal(t, 2, w.Len())
	assert.Equal(t, 2, w.Cap())
	assert.Equal(t, nil, w.GetOrDefault("C"))
	assert.Equal(t, nil, w.Get("C"))
	assert.Equal(t, int64(123), w.GetOr("C", int64(123)))
	assert.Equal(t, 1, w.Get("A"))
	assert.True(t, w.ContainsValue(1))
	assert.False(t, w.ContainsValue(2))
	assert.True(t, w.All(value.IsNotZero[any]))

	val, exists := w.TryGet("B")
	assert.True(t, exists)
	assert.Equal(t, map[string]int{}, val)

	_, exists = w.TryGet("C")
	assert.False(t, exists)
	assert.False(t, w.Contains("C"))
	assert.False(t, w.ContainsKey("C"))
}

func TestMap_Set(t *testing.T) {
	raw := map[string]interface{}{
		"A": 1,
		"B": map[string]int{},
	}
	w := Wrap(raw)

	w.Set("D", 123)
	assert.Equal(t, 123, w.Get("D"))

	merge := func(old interface{}) interface{} {
		if val, ok := old.(int); ok {
			return val + 234
		}
		return 135
	}
	w.Merge("D", 234, merge)
	assert.Equal(t, 357, w.Get("D"))

	w.Merge("E", 234, merge)
	assert.Equal(t, 234, w.GetOr("E", 789))
}

func TestMap_Find(t *testing.T) {
	raw := map[string]interface{}{
		"A": 1,
		"B": map[string]int{},
	}
	w := Wrap(raw)

	key, ok := w.Find(1)
	assert.True(t, ok)
	assert.Equal(t, "A", key)

	_, ok = w.FindFn(func(_ any) bool { return false })
	assert.False(t, ok)
}

func TestMap_ForAll(t *testing.T) {
	raw := map[string]interface{}{
		"A": 1,
		"B": map[string]int{},
	}
	w := Wrap(raw)

	w2 := MakeWithCap[string, any](2)
	w.ForAll2(func(k string, v any) {
		w2.Set(k, v)
	})
	assert.Equal(t, w.Raw(), w2.Raw())

}
