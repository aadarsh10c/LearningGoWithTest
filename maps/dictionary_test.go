package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrigs(got, want, t)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"
		if err == nil {
			t.Fatal("expeted to get an error.")
		}
		assertStrigs(err.Error(), want, t)
	})

}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	want := "this is just a test"
	got, err := dictionary.Search("test")

	//if error is given show error
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	//if errr not given compare
	assertStrigs(got, want, t)
}

func assertStrigs(got string, want string, t *testing.T) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
