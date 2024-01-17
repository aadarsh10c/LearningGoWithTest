package storecontext

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type spyStore struct {
	response  string
	cancelled bool
}

func (s *spyStore) Fetch() string {
	time.Sleep(1 * time.Second)
	return s.response
}
func (s *spyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	t.Run("Fetch hello world response", func(t *testing.T) {
		data := "Hello,World"
		store := &spyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
	})
	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "Hellow World"
		store := &spyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellinCTX, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellinCTX)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if !store.cancelled {
			t.Error("store was not told to cancel")
		}

	})
}
