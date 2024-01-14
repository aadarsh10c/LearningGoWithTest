package synctest

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementting the counter 3 times", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
		assertCounter(&counter, 3, t)
	})
	t.Run("it runs safley concurrently ", func(t *testing.T) {
		counter := Counter{}
		wantedCount := 1000
		var wg sync.WaitGroup
		wg.Add(wantedCount)
		for i := 0; i < 1000; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
		assertCounter(&counter, wantedCount, t)

	})

}

func assertCounter(counter *Counter, want int, t *testing.T) {
	if counter.Value() != want {
		t.Errorf("%d, want %d", counter.Value(), want)
	}
}
