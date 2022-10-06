package errors

import "fmt"

type PanicError struct {
	PanicMsg interface{}
}

func (pe *PanicError) Error() string {
	return fmt.Sprintf("panic: %v", pe.PanicMsg)
}

func CatchPanic(fn func()) (finalErr error) {
	defer func() {
		iface := recover()
		if iface == nil {
			finalErr = nil
			return
		}

		if err, ok := iface.(error); ok {
			finalErr = err
			return
		}

		finalErr = &PanicError{iface}
		return
	}()

	fn()
	return
}
