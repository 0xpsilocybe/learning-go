package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{
		"test": "this is just a test",
	}
	
	t.Run("search known word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "this is just a test"
		assertNoError(t, err)
		assertStrings(t, got, want)
	})

	t.Run("search unknown word", func(t *testing.T) {
		_, err := dictionary.Search("bulbulator")
		want := "could not find the word you're looking for"
		assertError(t, err, ErrNotFound)
		assertStrings(t, err.Error(), want)
	})

}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	if got != want {
		t.Errorf("got error '%s', want '%s'", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("got an error '%s' but didnt want one", err)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

