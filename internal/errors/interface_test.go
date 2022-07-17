package errors

import (
	"testing"
)

func TestInterface(t *testing.T) {
	staticAssertIsStackTraceSpanNode(t, &errorWithStack{})
	staticAssertIsStackTraceSpanNode(t, &errorWithMark{})
	staticAssertIsStackTraceSpanNode(t, &errorWithMessage{})
}

/*
staticAssertIsStackTraceSpanNode assert that err is stackTraceSpanNode while compiling
*/
func staticAssertIsStackTraceSpanNode(t *testing.T, err stackTraceSpanNode) {}
