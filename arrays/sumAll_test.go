package arrays

import (
	"fmt"
	"testing"
)

func TestSumAll(t *testing.T) {
	t.Run("function take 2 slices", func(t *testing.T) {
		got := SumAll([]int{1, 1, 1}, []int{0, 9})
		want := []int{3, 9}
		fmt.Printf("git %v", got)
		AssertEqual(got, want, t)

	})
}

func AssertEqual(got []int, want []int, t *testing.T) {
	t.Helper()
	for index, number := range got {
		if number != want[index] {
			t.Errorf("Got %v, wanted %v", got, want)
		}
	}
}

func ExampleSumAll() {
	answer := SumAll([]int{1, 2, 3}, []int{4, 5}, []int{6, 7, 8})
	fmt.Printf("%v", answer)
	//Output: [6 9 21]
}
