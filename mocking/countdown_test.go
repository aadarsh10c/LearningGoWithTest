package main

import (
	"bytes"
	"testing"
)

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}
	CountDown(buffer, spySleeper)
	got := buffer.String()
	want := "3\n2\n1\nGo!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
	}
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}
