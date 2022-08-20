package bang

import "github.com/snowphoenix0105/bang/internal/warning"

func SetOnWarning(fn func(msg string)) {
	warning.SetOnWarning(fn)
}
