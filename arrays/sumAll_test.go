package arrays

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 3, 9}, []int{3, 9})
	want := []int{13, 12}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, wantd %v", got, want)
	}
}
