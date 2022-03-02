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
			t.Errorf(cmp.Diff(want, s))
		}
	})
}
