package mocking

import (
	"fmt"
	"io"
)

const countdownStart = 3
const finalValue = "go"

type Sleeper interface {
	Sleep()
}

func CountDown(w io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		sleeper.Sleep()
	}
	fmt.Fprintln(w, finalValue)
}
