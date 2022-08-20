package errors

import "strings"

func FormatErrorWithFullStackTrace(err error) string {
	builder := strings.Builder{}

	markedStackTrace, baseError := CollectMarkedStackTraceFromError(err)

	builder.WriteString(err.Error())
	builder.WriteString("\nStacktrace:\n")
	markedStackTrace.WriteMultiLine(&builder)
	builder.WriteString("Message:\t")
	builder.WriteString(baseError.Error())

	return builder.String()
}
