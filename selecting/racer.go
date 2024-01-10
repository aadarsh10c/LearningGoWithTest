package selecting

import (
	"net/http"
	"time"
)

func Racer(a, b string) string {
	durationA := measureTime(a)
	durationB := measureTime(b)

	if durationA < durationB {
		return a
	}
	return b
}

func measureTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	duration := time.Since(start)
	return duration
}
