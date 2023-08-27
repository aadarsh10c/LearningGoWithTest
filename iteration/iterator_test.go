package iteration

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	repeated := Iterator("a", 6)
	expected := "aaaaaa"

	if expected != repeated {
		t.Errorf("expected %q but got %q", expected, repeated)
	}

}

func BenchmarkIterator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iterator("a", 5)
	}
}

func ExampleIterator() {
	answer := Iterator("a", 5)
	fmt.Println(answer)
	//Output: aaaaa
}
