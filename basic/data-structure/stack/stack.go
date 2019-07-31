package stack

import "errors"

type Stack []interface{}

func (stack Stack) Cap() int {
	return cap(stack)
}
func (stack Stack) Len() int {
	return len(stack)
}
func (stack *Stack) Push(value interface{}) {
	*stack = append(*stack, value)
}
func (stack *Stack) Pop() (interface{}, error) {
	length := len(*stack)
	if length == 0 {
		return nil, errors.New("out of index, len is 0")
	}

	last := length - 1

	value := (*stack)[last]
	*stack = (*stack)[:last]

	return value, nil
}