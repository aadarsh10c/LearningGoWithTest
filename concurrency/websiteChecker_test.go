package concurrency

import (
	"reflect"
	"testing"
	// "time"
)

func mockWebsiteCheck(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestCheckWebsites(t *testing.T) {
	t.Run("Check websites", func(t *testing.T) {
		websites := []string{
			"http://google.com",
			"http://blog.gypsydave5.com",
			"waat://furhurterwe.geds",
		}
		got := CheckWebsites(mockWebsiteCheck, websites)
		want := map[string]bool{
			"http://google.com":          true,
			"http://blog.gypsydave5.com": true,
			"waat://furhurterwe.geds":    false,
		}

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("wanted %v, got %v", want, got)
		}

	})
}

func slowStubWebsiteChecket(_ string) bool {
	// time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < 100; i++ {
		urls[i] = "a true"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecket, urls)
	}
}
