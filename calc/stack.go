package main

import "errors"

type (
	Stack struct {
		top *node
		length int
	}
	node struct {
		value int
		prev *node
	}
)

func New() *Stack {
	return &Stack{nil,0}
}

func (stack *Stack) Pop() (int, error) {
	if stack.length == 0 {
		return 0, errors.New("stack is empty")
	}

	n := stack.top
	stack.top = n.prev
	stack.length--

	return n.value, nil
}

func (stack *Stack) Push(value int) {
	n := &node{value,stack.top}
	stack.top = n
	stack.length++
}