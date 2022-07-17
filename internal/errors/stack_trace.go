package errors

import (
	"fmt"
	"io"
	"runtime"
	"strings"
)

func FormatStackTrace(err error) string {
	builder := strings.Builder{}

	formatStackTrace(err, &builder)

	return builder.String()
}

func formatStackTrace(err error, writer io.Writer) {
	trace := stackTrace{}
	trace.applyError(err)

	trace.printMessage(writer)

	writer.Write([]byte{'\n', '\n'})

	trace.printStackFrames(writer)
}

type runtimeFrame struct {
	pc      uintptr
	msgList []string
}

type stackTraceErrorMessageMode int

const (
	stackTraceErrorMessageModeInvalid stackTraceErrorMessageMode = iota
	stackTraceErrorMessageModeNormal
	stackTraceErrorMessageModeCompatible
)

type stackTrace struct {
	originError           error
	runtimeStackFrameList []runtimeFrame
	wrappedMessageList    []string
	baseError             error
	errorMessageMode      stackTraceErrorMessageMode
}

func (p *stackTrace) applyError(err error) {
	collector := stackTraceCollector{}

	collector.doCollect(err)

	p.originError = err
	p.wrappedMessageList = collector.wrapMsgList
	p.runtimeStackFrameList = collector.frameList
	p.baseError = collector.baseError

	if collector.stackTraceSpanHasExternalWrapper {
		p.errorMessageMode = stackTraceErrorMessageModeCompatible
	} else {
		p.errorMessageMode = stackTraceErrorMessageModeNormal
	}
}

func (p *stackTrace) printMessage(writer io.Writer) {
	// printStackFrames error message
	switch p.errorMessageMode {
	case stackTraceErrorMessageModeNormal:
		for _, msg := range p.wrappedMessageList {
			io.WriteString(writer, msg)
			io.WriteString(writer, ErrorMessageSplitter)
		}
		io.WriteString(writer, p.baseError.Error())
	case stackTraceErrorMessageModeCompatible:
		fallthrough
	default:
		io.WriteString(writer, p.originError.Error())
	}
}

func (p *stackTrace) printStackFrames(writer io.Writer) {
	printer := stackTraceFramePrinter{
		stackTrace: p,
		writer:     writer,
	}
	printer.doPrint()
}

type stackTraceCollector struct {
	frameList   []runtimeFrame
	wrapMsgList []string

	/*
		baseError record the first external error outside the stackTrace span:

			1. internal -> internal -> internal -> *external* -> external
			2. internal -> external -> internal -> *external* -> external
			3. external -> internal -> *external* -> external
			4. external -> internal -> *external*

	*/
	baseError error

	/*
		stackTraceSpanHasExternalWrapper describe whether there is an external error type in the stackTrace span:

			* internal -> internal -> internal -> external 	=> false
			* internal -> external -> internal -> external 	=> true
			* internal -> external -> internal 				=> true
			* internal -> internal -> external				=> false
	*/
	stackTraceSpanHasExternalWrapper bool
}

func (c *stackTraceCollector) doCollect(err error) {
	c.collect(err)
}

func (c *stackTraceCollector) collect(err error) bool {
	switch realError := err.(type) {
	case *errorWithStack:
		c.collectSubFromStackTraceSpanNode(realError)

		if c.frameList == nil {
			c.setupFrameList(realError.runtimeStackPCList)
		}

		return true

	case *errorWithMark:
		c.collectSubFromStackTraceSpanNode(realError)

		var targetIndex int
		if c.frameList == nil {
			c.setupFrameList(realError.runtimeStackPCList)
			targetIndex = 0
		} else {
			targetIndex = c.getMinFrameIndexOfCommonPC(realError.runtimeStackPCList)
		}

		frame := &c.frameList[targetIndex]
		frame.msgList = append(frame.msgList, realError.msg)

		return true

	case *errorWithMessage:
		c.wrapMsgList = append(c.wrapMsgList, realError.msg)

		c.collectSubFromStackTraceSpanNode(realError)

		return true

	default:
		c.tryCollectSub(err)

		return false
	}
}

func (c *stackTraceCollector) getMinFrameIndexOfCommonPC(pcList []uintptr) int {
	frameIndex := len(c.frameList) - 1
	pcIndex := len(pcList) - 1

	for frameIndex >= 0 && pcIndex >= 0 && c.frameList[frameIndex].pc == pcList[pcIndex] {
		frameIndex--
		pcIndex--
	}

	return frameIndex
}

func (c *stackTraceCollector) setupFrameList(pcList []uintptr) {
	frameList := make([]runtimeFrame, len(pcList))
	for i, pc := range pcList {
		frameList[i].pc = pc
	}

	c.frameList = frameList
}

func (c *stackTraceCollector) collectSubFromStackTraceSpanNode(err stackTraceSpanNode) {
	subError := err.Unwrap()

	if DEBUG {
		if subError == nil {
			panic("subError is nil")
		}
	}

	subIsInternalWrapper := c.collect(subError)
	if !subIsInternalWrapper {
		c.baseError = subError
	}
}

func (c *stackTraceCollector) tryCollectSub(err error) bool {
	if subError := Unwrap(err); subError != nil {
		subIsInternalWrapper := c.collect(subError)

		if subIsInternalWrapper {
			c.stackTraceSpanHasExternalWrapper = true
		}

		return true
	}
	return false
}

type stackTraceFramePrinter struct {
	stackTrace *stackTrace
	writer     io.Writer
}

func (p *stackTraceFramePrinter) doPrint() {
	for _, frame := range p.stackTrace.runtimeStackFrameList {
		fn := runtime.FuncForPC(frame.pc)
		file, line := fn.FileLine(frame.pc)
		fmt.Fprintf(p.writer, "%s \n\t@ %s:%d\n", funcName(fn), file, line)

		for _, msg := range frame.msgList {
			fmtMsg := strings.ReplaceAll(msg, "\n", "\n\t* ")
			fmt.Fprintf(p.writer, "\t> %s\n", fmtMsg)
		}
	}
}
