package stack

import "errors"

type Stack struct {
	data []interface{}
}

func (s *Stack) pop() (interface{}, error) {
	l := len(s.data)
	if l > 0 {
		out := s.data[l-1]
		s.data = s.data[:l-1]
		return out, nil
	} else {
		return nil, errors.New("empty")
	}
}

func (s *Stack) push(data interface{}) {
	s.data = append(s.data, data)
}

func (s *Stack) len() int {
	return len(s.data)
}

func (s *Stack) top() (interface{}, error) {
	l := len(s.data)
	if l == 0 {
		return nil, errors.New("empty")
	}
	return s.data[l-1], nil
}

func GetStack() *Stack {
	return &Stack{data: make([]interface{}, 0)}
}
