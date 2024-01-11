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

		got, _ := Racer(slowUrl, fastUrl)
		want := fastUrl

		if got != want {
			t.Fatalf("Got %q ,want %q", got, want)
		}

	})
	t.Run("Expect error after 10 seconds", func(t *testing.T) {
		slowServer := makeDelayer(12 * time.Second)
		fastServer := makeDelayer(11 * time.Second)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		_, err := Racer(slowURL, fastURL)

		if err == nil {
			t.Fatal("Expected an error but did not recieve one")
		}

	})
}

func makeDelayer(duration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
}
