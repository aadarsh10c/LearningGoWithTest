package stringstest

import (
	"strings"
	"testing"
)

func TestComapre(t *testing.T) {
	got := strings.Compare("a", "b")
	want := -1

	if got != want {
		t.Errorf("Got %d ,wanted %d ", got, want)
	}
}

func TestContainsRune(t *testing.T) {
	got := strings.ContainsRune("rdvrcka", 97)
	want := true

	if got != want {
		t.Errorf("Got %t ,wanted %t ", got, want)
	}
}
