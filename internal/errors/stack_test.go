package errors

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStackEqual(t *testing.T) {
	err := testFuncStdErrorf()

	marked, markedBase := CollectMarkedStackTraceFromError(err)
	raw, rawBase := CollectStackTrace(err)

	require.Equal(t, markedBase, rawBase)
	require.Equal(t, len(marked), len(raw))

	for i := 0; i < len(raw); i++ {
		assert.Equal(t, marked[i].Frame(), raw[i])
	}
}

func TestCast(t *testing.T) {
	err := testFuncStdErrorf()
	trace, _ := CollectStackTrace(err)
	assert.Equal(t, trace, PCListToStackTrace(StackTraceToPCList(trace)))
}

func TestFrame_Format(t *testing.T) {
	err := testFuncStdErrorf()
	trace, base := CollectStackTrace(err)
	t.Log(base)

	builder := &strings.Builder{}
	for _, frame := range trace {
		fmt.Fprintf(builder, "%n\n\t@ %v\n", frame, frame)
	}
	t.Log(builder.String())

	builder.Reset()
	for _, frame := range trace {
		fmt.Fprintf(builder, "%+s", frame)
	}
	t.Log(builder.String())

	builder.Reset()
	for _, frame := range trace {
		fmt.Fprintf(builder, "%n\n\t@ %s:%d\n", frame, frame.FullName(), frame)
	}
	t.Log(builder.String())

	builder.Reset()
	fmt.Fprintf(builder, "%+v", trace)
	t.Log(builder.String())

	builder.Reset()
	fmt.Fprintf(builder, "%v", trace)
	t.Log(builder.String())

	builder.Reset()
	fmt.Fprintf(builder, "%s", trace)
	t.Log(builder.String())

	builder.Reset()
	fmt.Fprintf(builder, "%n", trace)
	t.Log(builder.String())
}
