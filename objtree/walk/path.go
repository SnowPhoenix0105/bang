package walk

import (
	"fmt"
	"strings"
)

type pathRecorder struct {
	pathStack []string
}

func (p *pathRecorder) EnterFiled(fieldName string) {
	p.pathStack = append(p.pathStack, fieldName)
}

func (p *pathRecorder) EnterIndex(index int) {
	p.pathStack = append(p.pathStack, fmt.Sprintf("[%d]", index))
}

func (p *pathRecorder) Exit() {
	p.pathStack = p.pathStack[:len(p.pathStack)-1]
}

func (p *pathRecorder) String() string {
	builder := strings.Builder{}
	for _, name := range p.pathStack {
		if name[0] == '[' {
			builder.WriteString(name)
			continue
		}
		builder.WriteRune('.')
		builder.WriteString(name)
	}
	return builder.String()
}
