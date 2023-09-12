package main

import (
	"reflect"
	"testing"
)

func TestCountDown(t *testing.T) {
	// buffer := &bytes.Buffer{}
	spySleeper := &SpyCountDownOperations{}
	CountDown(spySleeper, spySleeper)
	// got := buffer.String()
	// want := "3\n2\n1\nGo!"
	want := []string{
		write,
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
	}

	if !reflect.DeepEqual(want, spySleeper.Calls) {
		t.Errorf("wanted calls %v got %v", want, spySleeper.Calls)
	}
}

type SpyCountDownOperations struct {
	Calls []string
}

func (s *SpyCountDownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountDownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"
