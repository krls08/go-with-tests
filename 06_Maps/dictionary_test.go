package main

import (
	"testing"
)

func TestDefaultSearch(t *testing.T) {
	value := "this is a test value"
	dictioary := map[string]string{"test": value}

	got := defaultSearch(dictioary, "test")
	want := value

	assertStrings(t, got, want)
	//if got != want {
	//	t.Errorf("got %q want %q, given %q", got, want, "test")
	//}
}

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unkwown")
		//want := "could not find the word you were looking for"
		if err == nil {
			t.Fatal("expected to get an error.")
		}

		//assertStrings(t, err.Error(), want)
		assertError(t, err, ErrNoWordFound)
	})

}

func TestAdd(t *testing.T) {
	dict := Dictionary{}

	t.Run("add new word", func(t *testing.T) {
		word := "newWord"
		definition := "this is a test"

		err := dict.Add(word, definition)
		if err != nil {
			t.Errorf("new value %q, should not return an error", "new value")
		}
		assertDefinition(t, dict, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dict := Dictionary{word: definition}

		err := dict.Add(word, "new test")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, word, definition)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dict := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dict.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dict, word, newDefinition)

	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dict := Dictionary{}

		err := dict.Update(word, definition)
		assertError(t, err, ErrNoWordFound)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dict := Dictionary{word: definition}

		err := dict.Delete(word)
		assertError(t, err, nil)

		_, err = dict.Search(word)
		assertError(t, err, ErrNoWordFound)
	})

	t.Run("detele unexisting word", func(t *testing.T) {
		word := "test"
		dict := Dictionary{}

		err := dict.Delete(word)
		assertError(t, err, ErrNoWordFound)

	})
}

func assertDefinition(t testing.TB, dict Dictionary, word, definition string) {
	t.Helper()
	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("Sound find added word:", err)
	}
	assertStrings(t, got, definition)

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	//if !errors.Is(got, want) {
	//	t.Errorf("got error: %q, wanted %q", got, want)
	//}
	if got != want {
		t.Errorf("got error: %q, wanted %q", got, want)
	}
}
