package queue

type Interface interface {
	Dequeue() any
	Enqueue(x any)
	Front() any
	End() any
}

type IntQueue []int

func NewIntQueue() IntQueue {
	iq := make(IntQueue, 0)
	return iq
}

func (iq *IntQueue) Dequeue() any {
	if len(*iq) == 0 {
		return nil
	}
	old := *iq
	n := len(old)
	x := old[0]
	*iq = old[1:n]
	return x
}

func (iq *IntQueue) Enqueue(x any) {
	*iq = append(*iq, x.(int))
}

func (iq IntQueue) Front() any {
	if len(iq) == 0 {
		return nil
	}
	return iq[0]
}

func (iq IntQueue) End() any {
	if len(iq) == 0 {
		return nil
	}
	return iq[len(iq)-1]
}
