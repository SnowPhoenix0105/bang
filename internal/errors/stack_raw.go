package errors

import (
	"fmt"
	"io"
	"path"
	"runtime"
	"strconv"

	"github.com/snowphoenix0105/bang/internal/caster"
)

type Frame uintptr

// PC returns the pc of the frame
func (f Frame) PC() uintptr {
	return uintptr(f)
}

// FullName returns the full function name of the Frame
func (f Frame) FullName() string {
	return runtime.FuncForPC(f.PC()).Name()
}

// Name returns the function name fo the Frame
func (f Frame) Name() string {
	return funcName(runtime.FuncForPC(f.PC()))
}

// File returns the filename of the Frame
func (f Frame) File() string {
	file, _ := f.FileLine()
	return file
}

// Line returns the line num of the Frame
func (f Frame) Line() int {
	_, line := f.FileLine()
	return line
}

// FileLine returns the filename and line of the Frame
func (f Frame) FileLine() (string, int) {
	return runtime.FuncForPC(f.PC()).FileLine(f.PC())
}

/*
Format formats the frame according to the fmt.Formatter interface.

   %s    source file
   %d    source line
   %n    function name
   %v    equivalent to %s:%d

Format accepts flags that alter the printing of some verbs, as follows:

   %+s   function name and path of source file relative to the compile time
         GOPATH separated by \n\t (<funcname>\n\t<path>)
   %+v   equivalent to %+s:%d
*/
func (f Frame) Format(state fmt.State, verb rune) {
	switch verb {
	case 's':
		switch {
		case state.Flag('+'):
			io.WriteString(state, f.FullName())
			io.WriteString(state, "\n\t")
			io.WriteString(state, f.File())
		default:
			io.WriteString(state, path.Base(f.File()))
		}
	case 'd':
		io.WriteString(state, strconv.Itoa(f.Line()))
	case 'n':
		io.WriteString(state, f.Name())
	case 'v':
		f.Format(state, 's')
		io.WriteString(state, ":")
		f.Format(state, 'd')
	}
}

/*
StackTrace TODO
*/
type StackTrace []Frame

func (st StackTrace) Format(state fmt.State, verb rune) {
	switch verb {
	case 'v':
		switch {
		case state.Flag('+'):
			for _, frame := range st {
				fmt.Fprintf(state, "\n%+v", frame)
			}
			return
		}
	}
	fmt.Fprintf(state, fmt.Sprintf("%%%c", verb), ([]Frame)(st))
}

var (
	stackTraceCaster = caster.MustNew[[]uintptr, StackTrace](nil)
	pcListCaster     = caster.MustNew[StackTrace, []uintptr](nil)
)

// PCListToStackTrace cast the []uintptr to StackTrace (they share the memory)
func PCListToStackTrace(pcList []uintptr) StackTrace {
	return stackTraceCaster(pcList)
}

// StackTraceToPCList cast the StackTrace to []uintptr (they share the memory)
func StackTraceToPCList(trace StackTrace) []uintptr {
	return pcListCaster(trace)
}

// CollectStackTrace returns the deepest StackTrace and the baseError of it
func CollectStackTrace(err error) (stack StackTrace, baseError error) {
	record := true
	for {
		if record {
			baseError = err
			record = false
		}

		if stackErr, ok := err.(StackError); ok {
			stack = stackErr.StackTrace()
			err = stackErr.Unwrap()
			record = true
			continue
		}

		if err = Unwrap(err); err != nil {
			continue
		}

		return stack, baseError
	}
}
