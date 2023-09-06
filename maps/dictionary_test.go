package maps

import "testing"

func TestSearch(t *testing.T) {
	//create adn instance of dictinary, woth values
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		//call the method on dictionary instance and record value
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		//test the response
		assertStrings(got, want, t)
	})

	t.Run("unkown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrNotFound
		//i error is not given
		if err == nil {
			t.Fatal("Word is present in the dictionary")
		}
		assertError(err, want, t)
	})

}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	want := "this is just a test"
	got, err := dictionary.Search("test")

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(got, want, t)
}

func assertError(got, want error, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, wanted %q, given %q", got, want, "test")

	}
}

func assertStrings(got, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, wanted %q, given %q", got, want, "test")
	}
}
