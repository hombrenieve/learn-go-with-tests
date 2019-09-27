package slices

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}

		got := Sum(numbers)
		want := 10

		if want != got {
			t.Errorf("expected %d but got %d, given %v", want, got, numbers)
		}
	})
}
