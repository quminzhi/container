package stack

type Interface interface {
	Pop() any
	Push(x any)
	Peek() any
}

type IntStack []int

func InitIntStack() IntStack {
	is := make(IntStack, 0)
	return is
}

func (is IntStack) Len() int {
	return len(is)
}

func (is IntStack) Peek() any {
	return is[len(is)-1]
}

func (is *IntStack) Pop() any {
	old := *is
	n := len(old)
	x := old[n-1]
	*is = old[0 : n-1]
	return x
}

func (is *IntStack) Push(x any) {
	*is = append(*is, x.(int))
}
