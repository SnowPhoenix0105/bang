package stack

type StringStack struct {
	data []string
}

func NewString() *StringStack {
	return &StringStack{}
}

func (ss *StringStack) Raw() []string {
	return ss.data
}

func (ss *StringStack) Push(str string) {
	ss.data = append(ss.data, str)
}

func (ss *StringStack) Top() string {
	return ss.data[ss.Depth()-1]
}

func (ss *StringStack) Pop() {
	ss.data = ss.data[:ss.Depth()-1]
}

func (ss *StringStack) Depth() int {
	return len(ss.data)
}
