package stack

import "errors"

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack *Stack) Push(x interface{}) {
	*stack = append(*stack, x)
}

func (stack *Stack) Pop() (interface{}, error) {
	_stack := *stack
	if _stack.Len() == 0 {
		return nil, errors.New("Stack is empty")
	}
	x := _stack[_stack.Len() -1]
	*stack = _stack[:_stack.Len() -1]
	return x, nil
}