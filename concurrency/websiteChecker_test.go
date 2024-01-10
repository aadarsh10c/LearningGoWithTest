package concurrency

import (
	"reflect"
	"testing"
)

type testWebsiteCallback struct{}

func (t *testWebsiteCallback) CallBack(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}
	t.Run("Websites map", func(t *testing.T) {
		test := &testWebsiteCallback{}
		got := CheckWebsites(test, websites)
		want := map[string]bool{
			"http://google.com":          true,
			"http://blog.gypsydave5.com": true,
			"waat://furhurterwe.geds":    false,
		}
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("wanted %v, got %v", want, got)
		}
	})
}
