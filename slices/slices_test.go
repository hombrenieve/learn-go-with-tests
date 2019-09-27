package slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4}

	got := Sum(numbers)
	want := 10

	if want != got {
		t.Errorf("expected %d but got %d, given %v", want, got, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 4, 9})
	want := []int{3, 13}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v when expecting %v", got, want)
	}
}
