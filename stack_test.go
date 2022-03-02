package threadsafestack

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStack(t *testing.T) {
	t.Run("Push to stack", func(t *testing.T) {
		s := StringStack{Value: make([]string, 0)}

		s.push("Peter")

		want := StringStack{Value: []string{"Peter"}}

		if !cmp.Equal(s, want) {
			t.Errorf(cmp.Diff(s, want))
		}
	})

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
}
