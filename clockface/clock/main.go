package main

import (
	"os"
	"time"

	"github.com/aadarsh10c/LearningGoWithTest/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
