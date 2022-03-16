package threadsafestack

import (
	"errors"
	"sync"
)

const ErrEmptyStack string = "stack is empty"

type Stack[T any] struct {
	mu    sync.Mutex
	Value []T
}

func newStringStack() Stack[string] {
	return Stack[string]{}
}

func (s *Stack[T]) push(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Value = append(s.Value, value)
}

func (s *Stack[T]) pop() (T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.Value) == 0 {
		var empty T
		return empty, errors.New(ErrEmptyStack)
	}

	lastIndex := len(s.Value) - 1
	lastElement := s.Value[lastIndex]

	s.Value = s.Value[:lastIndex]

	return lastElement, nil

}

func (s *Stack[T]) isEmpty() bool {
	return len(s.Value) == 0
}
