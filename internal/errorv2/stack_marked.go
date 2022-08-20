package errorv2

import (
	"fmt"
	"io"
	"runtime"
	"strings"
)

type MarkedFrame struct {
	frame Frame
	marks []string
}

func (mf MarkedFrame) Frame() Frame {
	return mf.frame
}

func (mf MarkedFrame) Marks() []string {
	ret := make([]string, len(mf.marks))
	copy(ret, mf.marks)
	return ret
}

func (mf MarkedFrame) WriteMultiLine(writer io.Writer) (int, error) {
	fn := runtime.FuncForPC(mf.frame.PC())
	file, line := fn.FileLine(mf.frame.PC())
	cnt, err := fmt.Fprintf(writer, "%s\n\t@ %s:%d\n", funcName(fn), file, line)
	if err != nil {
		return cnt, err
	}

	for _, msg := range mf.marks {
		fmtMsg := strings.ReplaceAll(msg, "\n", "\n\t* ")
		n, err := fmt.Fprintf(writer, "\t> %s\n", fmtMsg)
		cnt += n
		if err != nil {
			return cnt, err
		}
	}

	return cnt, nil
}

type MarkedStackTrace []MarkedFrame

func CollectMarkedStackTraceFromError(err error) (MarkedStackTrace, error) {
	collector := markedStackTraceCollector{}
	collector.buildTraceOf(err)
	return collector.Result, collector.BaseError
}

type markedStackTraceCollector struct {
	Result    MarkedStackTrace
	BaseError error
}

func (c *markedStackTraceCollector) buildTraceOf(err error) {
	if err == nil {
		return
	}

	wrapErr, ok := err.(WrapError)
	if !ok {
		return
	}
	c.buildTraceOf(wrapErr.Unwrap())

	stackErr, ok := wrapErr.(StackError)
	if !ok {
		if len(c.Result) == 0 {
			c.BaseError = err
		}
		return
	}
	rawStack := stackErr.StackTrace()
	if len(rawStack) < 0 {
		return
	}

	isFirst := false
	if len(c.Result) == 0 {
		c.buildResult(rawStack)
		isFirst = true
	}

	markErr, ok := stackErr.(MarkError)
	if !ok {
		return
	}

	frameIndex := 0
	if !isFirst {
		frameIndex = c.getMarkedFrameIndex(rawStack)
		if frameIndex < 0 {
			return
		}
	}

	c.Result[frameIndex].marks = append(c.Result[frameIndex].marks, markErr.Mark())
}

func (c *markedStackTraceCollector) getMarkedFrameIndex(stack StackTrace) int {
	frameIndex := len(c.Result) - 1
	pcIndex := len(stack) - 1

	for frameIndex >= 0 && pcIndex >= 0 && c.Result[frameIndex].frame == stack[pcIndex] {
		frameIndex--
		pcIndex--
	}

	return frameIndex
}

func (c *markedStackTraceCollector) buildResult(rawStack StackTrace) {
	c.Result = make(MarkedStackTrace, len(rawStack))
	for i, frame := range rawStack {
		c.Result[i].frame = frame
	}
}

func (mt MarkedStackTrace) WriteMultiLine(writer io.Writer) (int, error) {
	cnt := 0
	length := len(mt)
	for i := 0; i < length; i++ {
		n, err := mt[length-i-1].WriteMultiLine(writer)
		cnt += n
		if err != nil {
			return cnt, err
		}
	}

	return cnt, nil
}
