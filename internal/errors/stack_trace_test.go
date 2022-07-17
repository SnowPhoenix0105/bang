package errors

import (
	"strings"
	"testing"
)

func TestStackTracePrint(t *testing.T) {
	collection := stackTrace{}

	fn1 := func() {
		pcList := getRuntimeStackPCList(1)
		collection.runtimeStackFrameList = make([]runtimeFrame, len(pcList))
		for i, pc := range pcList {
			collection.runtimeStackFrameList[i].pc = pc
		}
	}

	fn2 := func() {
		fn1()
	}

	fn3 := func() {
		fn2()
	}

	fn3()

	collection.runtimeStackFrameList[0].msgList = []string{
		"msg1",
		"msg2",
	}

	buf := strings.Builder{}
	printer := stackTraceFramePrinter{
		stackTrace: &collection,
		writer:     &buf,
	}

	printer.doPrint()

	t.Logf("\n%s", buf.String())
}
