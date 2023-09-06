package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}
	got := Search(dictionary, "test")
	want := "this is just a test"

	assertStrings(got, want, t)
}

func assertStrings(got string, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, wanted %q, given %q", got, want, "test")
	}
}
