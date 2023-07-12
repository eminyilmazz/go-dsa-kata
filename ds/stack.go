package ds

import "errors"

type Stack struct {
	elements []interface{}
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.elements[s.Size()-1]
}

func (s *Stack) Push(e interface{}) {
	s.elements = append(s.elements, e)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("empty stack")
	}
	val := s.Peek()
	s.elements = s.elements[:s.Size()-1]
	return val, nil
}
