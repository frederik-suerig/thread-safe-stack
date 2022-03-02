package threadsafestack

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPushStack(t *testing.T) {
	t.Run("Push to stack", func(t *testing.T) {
		s := StringStack{Value: make([]string, 0)}

		s.push("Peter")

		want := StringStack{Value: []string{"Peter"}}

		if !cmp.Equal(s, want) {
			t.Errorf(cmp.Diff(s, want))
		}
	})
}

func TestPopStack(t *testing.T) {
	t.Run("Pop the last element from the stack", func(t *testing.T) {
		s := StringStack{}

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
		s := StringStack{}

		_, err := s.pop()
		want := errors.New(ErrEmptyStack)

		if err == nil {
			t.Error("Received no error when it should return one")
		}

		if err.Error() != want.Error() {
			t.Errorf(cmp.Diff(err.Error(), want.Error()))
		}
	})
}
