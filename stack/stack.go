// This package contains an abstract implementation of Stack
// data structure
package stack

import "errors"

// template type Stack(TValue)
type TValue *int

// A structure that represents a stack data structure
//
// Example:
// stack := &stack.Stack{}
// stack.Push(new(TValue))
// value, err := stack.Pop()
type Stack struct {
	data []TValue
}

// Adds an element on top of the stack
func (s *Stack) Push(value TValue) {
	s.data = append(s.data, value)
}

// Removes an element from top of the stack.
// If the stack is empty, it returns an error.
func (s *Stack) Pop() (TValue, error) {
	length := len(s.data)
	if length == 0 {
		return nil, errors.New("Stack is empty")
	}

	value := s.data[length-1]
	s.data = s.data[:length-1]
	return value, nil
}
