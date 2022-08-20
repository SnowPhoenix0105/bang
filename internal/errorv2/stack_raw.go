package errorv2

import (
	"fmt"
	"io"
	"path"
	"runtime"
	"strconv"

	"github.com/snowphoenix0105/bang/internal/caster"
)

type Frame uintptr

func (f Frame) PC() uintptr {
	return uintptr(f)
}

func (f Frame) FullName() string {
	return runtime.FuncForPC(f.PC()).Name()
}

func (f Frame) Name() string {
	return funcName(runtime.FuncForPC(f.PC()))
}

func (f Frame) File() string {
	file, _ := f.FileLine()
	return file
}

func (f Frame) Line() int {
	_, line := f.FileLine()
	return line
}

func (f Frame) FileLine() (string, int) {
	return runtime.FuncForPC(f.PC()).FileLine(f.PC())
}

// Format formats the frame according to the fmt.Formatter interface.
//
//    %s    source file
//    %d    source line
//    %n    function name
//    %v    equivalent to %s:%d
//
// Format accepts flags that alter the printing of some verbs, as follows:
//
//    %+s   function name and path of source file relative to the compile time
//          GOPATH separated by \n\t (<funcname>\n\t<path>)
//    %+v   equivalent to %+s:%d
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

type StackTrace []Frame

var (
	stackTraceCaster = caster.MustNew[[]uintptr, StackTrace](nil)
	pcListCaster     = caster.MustNew[StackTrace, []uintptr](nil)
)

func PCListToStackTrace(pcList []uintptr) StackTrace {
	return stackTraceCaster(pcList)
}

func StackTraceToPCList(trace StackTrace) []uintptr {
	return pcListCaster(trace)
}
func (st StackTrace) Format(state fmt.State, verb rune) {
	switch verb {
	case 'v':
		switch {
		case state.Flag('+'):
			for _, frame := range st {
				fmt.Fprintf(state, "\n%+v", frame)
			}
		}
	}
}
