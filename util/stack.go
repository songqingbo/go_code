package util

type Stack struct {
	size   int
	top    int
	values []interface{}
}

func MakeStack(size int) Stack {
	q := Stack{}
	q.size = size
	q.values = make([]interface{}, size)
	return q
}

func (stack *Stack) Push(element interface{}) bool {
	stack.values[stack.top] = element
	stack.top ++
	return true
}

// 判断是否满了
func (stack *Stack) IsFull() bool {
	return stack.size == stack.top
}

// 判断是否为空
func (stack *Stack) IsEmpty() bool {
	return stack.top == 0
}

func (stack *Stack) Pop() (r interface{}, err error) {
	stack.top --
	response := stack.values[stack.top]
	return response, nil
}
