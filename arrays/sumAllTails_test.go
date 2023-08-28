package arrays

import "testing"

func TestSumAllTails(t *testing.T) {
	t.Run("Takes 2 arrays,", func(t *testing.T) {
		got := SumAllTails([]int{2, 9}, []int{9, 0})
		want := []int{9, 0}
		AssertEqual(got, want, t)
	})
	t.Run("One of the array is empty", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3})
		want := []int{0, 5}
		AssertEqual(got, want, t)
	})
}
