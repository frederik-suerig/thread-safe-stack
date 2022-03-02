package threadsafestack

import (
	"errors"
	"sync"
)

const ErrEmptyStack string = "stack is empty"

type StringStack struct {
	mu    sync.Mutex
	Value []string
}

func newStringStack() StringStack {
	return StringStack{Value: []string{}}
}

func (s *StringStack) push(value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Value = append(s.Value, value)
}

func (s *StringStack) pop() (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
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
