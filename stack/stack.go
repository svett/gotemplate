package stack

import "errors"

// template type Stack(TValue)
type TValue *int

type Stack struct {
	data []TValue
}

func (s *Stack) Push(value TValue) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() (TValue, error) {
	length := len(s.data)
	if length == 0 {
		return nil, errors.New("Stack is empty")
	}

	value := s.data[length-1]
	s.data = s.data[:length-1]
	return value, nil
}
