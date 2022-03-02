package threadsafestack

import (
	"errors"
	"fmt"
	"sync"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPushStack(t *testing.T) {
	t.Run("Push to stack", func(t *testing.T) {
		s := newStringStack()

		s.push("Peter")

		want := []string{"Peter"}

		if !cmp.Equal(s.Value, want) {
			t.Errorf(cmp.Diff(s.Value, want))
		}
	})

	t.Run("Runs safe concurrently", func(t *testing.T) {
		wantedCount := 1000
		s := newStringStack()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(i int) {
				defer wg.Done()
				s.push(fmt.Sprintf("GoRoutine %d", i))
			}(i)
		}

		wg.Wait()

		if len(s.Value) != 1000 {
			t.Errorf("Not all goroutine completed")
		}
	})
}

func TestPopStack(t *testing.T) {
	t.Run("Pop the last element from the stack", func(t *testing.T) {
		s := newStringStack()

		s.push("Peter")

		got, err := s.pop()
		want := "Peter"

		if !cmp.Equal(err, nil) {
			t.Errorf(cmp.Diff(err, nil))
		}

		if !cmp.Equal(got, want) {
			t.Errorf(cmp.Diff(got, want))
		}
	})

	t.Run("Return error on pop if stack is empty", func(t *testing.T) {
		s := newStringStack()

		_, err := s.pop()
		want := errors.New(ErrEmptyStack)

		if err == nil {
			t.Error("Received no error when it should return one")
		}

		if err.Error() != want.Error() {
			t.Errorf(cmp.Diff(err.Error(), want.Error()))
		}
	})

	t.Run("Runs safe concurrently", func(t *testing.T) {
		wantedCount := 1000
		s := newFilledStringStack(wantedCount)

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(i int) {
				defer wg.Done()
				s.pop()
			}(i)
		}

		wg.Wait()

		if len(s.Value) != 0 {
			t.Errorf("Not all goroutine completed")
		}
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("Return true for an empty stack", func(t *testing.T) {
		s := newStringStack()

		emt := s.isEmpty()

		if !emt {
			t.Errorf("isEmpty() returnd False when it should return True")
		}
	})

	t.Run("Return false for an non-empty stack", func(t *testing.T) {
		s := newStringStack()

		s.push("Peter")

		emt := s.isEmpty()

		if emt {
			t.Errorf("isEmpty() returnd True when it should return False")
		}
	})

	t.Run("Return True after all elements where poped", func(t *testing.T) {
		s := newStringStack()

		s.push("Peter")
		s.pop()

		emt := s.isEmpty()

		if !emt {
			t.Errorf("isEmpty() returnd False when it should return True")
		}
	})
}

func newFilledStringStack(size int) StringStack {
	s := newStringStack()

	for i := 0; i < size; i++ {
		s.push(fmt.Sprintf("GoRoutine %d", i))
	}

	return s
}
