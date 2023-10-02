package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("Collection of 5 size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Fatalf("Got %d , wanted %d, given %v", got, want, "test")
		}
	})
	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Fatalf("Got %d , wanted %d, given %v", got, want, "test")
		}
	})

}
