package main

import (
	"os"
	"time"

	m "github.com/aadarsh10c/LearningGoWithTest/mocking"
)

type DefaultSleep struct{}

func (d *DefaultSleep) Sleep(){
	time.Sleep(1 * time.Second)
}

func main() {
	d := &DefaultSleep{}
	m.CountDown(os.Stdout, d)
}
