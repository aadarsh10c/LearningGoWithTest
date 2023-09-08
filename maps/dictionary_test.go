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
	word := "test"
	definiton := "this is just a test"
	t.Run("new Word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add(word, definiton)
		assertError(err, nil, t)
		asserDefinition(dictionary, word, t, definiton)
	})
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{word: definiton}
		err := dictionary.Add(word, "new word")
		assertError(err, ErrWordExists, t)
		asserDefinition(dictionary, word, t, definiton)
	})

}

func asserDefinition(dictionary Dictionary, word string, t *testing.T, definiton string) {
	t.Helper()
	got, _ := dictionary.Search(word)
	assertStrings(got, definiton, t)
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
