package clockface

import (
	"testing"
	"time"
)

func TestClockface(t *testing.T) {
	t.Run("Second hand at midnight", func(t *testing.T) {
		tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

		got := SecondHand(tm)
		want := Point{X: 150, Y: 150 - 90}

		assertStruct(got, want, t)
	})
	t.Run("Second hand at 30 secs", func(t *testing.T) {
		tm := time.Date(337, time.January, 1, 0, 0, 30, 0, time.UTC)

		got := SecondHand(tm)
		want := Point{150, 150 + 90}

		assertStruct(got, want, t)
	})
}

func assertStruct(got, want interface{}, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("Got %v , want %v", got, want)
	}
}
