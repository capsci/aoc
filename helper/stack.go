package helper

import "errors"

type Stack []interface{}

var ErrEmptyStack error = errors.New("Stack is empty")

func (s *Stack) Push(value interface{}) {
	*s = append(*s, value)
}

func (s *Stack) Pop(value interface{}) (elem interface{}, err error) {
	if !s.Empty() {
		elem = (*s)[s.Size()-1]
		(*s) = (*s)[:s.Size()-1]
		return elem, nil
	}
	return nil, ErrEmptyStack
}

func (s *Stack) Peek() (elem interface{}, err error) {
	if !s.Empty() {
		return (*s)[0], nil
	}
	return nil, ErrEmptyStack
}

func (s *Stack) Empty() bool {
	return s.Size() == 0
}

func (s *Stack) Size() int {
	return len(*s)
}
