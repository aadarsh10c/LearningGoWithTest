package main

import (
	"log"
	"os"

	d "github.com/aadarsh10c/LearningGoWithTest/dinjection"

	"net/http"
)

func main() {
	d.Greet(os.Stdout, "Test")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(d.MyGreeterHandler)))
}
