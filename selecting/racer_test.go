package selecting

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("fast url", func(t *testing.T) {
		slowServer := makeDelayer(20 * time.Millisecond)
		fastServer := makeDelayer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		got := Racer(slowUrl, fastUrl)
		want := fastUrl

		if got != want {
			t.Fatalf("Got %q ,want %q", got, want)
		}

	})
}

func makeDelayer(duration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
}
