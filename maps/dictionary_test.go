package dictionary

import (
	"testing"
)

func TestDelete(t *testing.T) {
	word := "test"
	definition := "word's definition"
	dictionary := Dictionary{word: definition}
	dictionary.Delete(word)

	_, err := dictionary.Search(word)

	if err != ErrNotFound {
		t.Errorf("expected %q to be deleted", word)
	}
}

func TestUpdate(t *testing.T) {
	t.Run("word exists", func(t *testing.T) {
		word := "test"
		definition := "initial definition"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)

		assertDefinition(t, dictionary, word, newDefinition)
		assertError(t, err, nil)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "initial definition"
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}

		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
		assertError(t, err, nil)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "it's a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "it's another test")

		assertDefinition(t, dictionary, word, definition)
		assertError(t, err, ErrWordExists)
	})
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "it's a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "it's a test"
		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		if err == nil {
			t.Fatal("Expected to get error but didn't get one.")
		}
		assertError(t, err, ErrNotFound)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if got != definition {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
