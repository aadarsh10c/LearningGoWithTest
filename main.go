package main

import (
	"os"

	d "github.com/aadarsh10c/LearningGoWithTest/dinjection"
)

func main() {
	d.Greet(os.Stdout, "Elodie")
}
