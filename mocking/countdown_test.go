package mocking

import (
	"bytes"
	"testing"
)

type spySleeper struct {
	Calls int
}

func (s *spySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
// 	t.Run("Print 3", func(t *testing.T) {

// 		CountDown(buffer)
// 		got := buffer.String()
// 		want := "3"

// 		compareString(got, want, t)

// 	})
// 	t.Run("Print 3,2,1,go", func(t *testing.T) {
// 		CountDown(buffer)

// 		got := buffer.String()
// 		want := `3
// 2
// 1
// go
// `

// 		compareString(got, want, t)

// 	})
	t.Run("Mocking", func(t *testing.T) {
		spy := &spySleeper{}
		CountDown(buffer, spy)
		got := spy.Calls
		want := 3
		if got != want {
			t.Fatalf("Got %d, wanted %d", got, want)
		}
	})
}

// func compareString(got string, want string, t *testing.T) {
// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }
