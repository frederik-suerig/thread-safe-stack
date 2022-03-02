package threadsafestack

import "errors"

const ErrEmptyStack string = "stack is empty"

type StringStack struct {
	Value []string
}

func (s *StringStack) push(value string) {
	s.Value = append(s.Value, value)
}

func (s *StringStack) pop() (string, error) {
	if len(s.Value) == 0 {
		return "", errors.New(ErrEmptyStack)
	}

	lastIndex := len(s.Value) - 1
	lastElement := s.Value[lastIndex]

	s.Value = s.Value[:lastIndex]

	return lastElement, nil

}

func (s *StringStack) isEmpty() bool {
	return len(s.Value) == 0
}
