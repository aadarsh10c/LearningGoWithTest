package arrays

import "testing"

func TestSumAll(t *testing.T) {
	t.Run("function take 2 slices", func(t *testing.T) {
		got := SumAll([]int{1, 1, 1}, []int{0, 9})
		want := []int{3, 9}

		assertEqual(got, want, t)

	})
}

func assertEqual(got []int, want []int, t *testing.T) {
	t.Helper()
	for index, number := range got {
		if number != want[index] {
			t.Errorf("Got %v, wanted %v", got, want)
		}
	}
}
