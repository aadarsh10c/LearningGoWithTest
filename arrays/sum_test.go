package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("Test with slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("Got %d, expected %d, given %v", got, want, numbers)
		}
	})

}