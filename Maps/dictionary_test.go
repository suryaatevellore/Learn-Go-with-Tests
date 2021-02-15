package main

import "testing"

func TestDictionary(t *testing.T) {
	// t.Fatal("not implemented")
	t.Run("test known word", func(t *testing.T) {
		dictionary := Dictionary{"heart-warming": "sweet", "purge": "remove"}

		got, _ := dictionary.Search("heart-warming")
		want := "sweet"
		assertStringsEqual(t, got, want)
	})

	t.Run("test unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test"}
		_, err := dictionary.Search("not-known-test")
		assertError(t, err, ErrNotFound)
	})

	t.Run("add new words", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		meaning := "this is not a test"
		dictionary.Add(word, meaning)
		assertDefinition(t, dictionary, word, meaning)
	})

	t.Run("add a word that exists already", func(t *testing.T) {
		word := "test"
		meaning := "this is a test"
		dictionary := Dictionary{word: meaning}
		err := dictionary.Add("test", "this is the second test")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, meaning)
	})

	t.Run("Update the definition of a word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is not a test"}
		dictionary.Update("test", "change meaning")
		assertDefinition(t, dictionary, "test", "change meaning")
	})

	t.Run("Update the definition of a word that does not exist", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is not a test"}
		err := dictionary.Update("name", "change meaning")
		assertError(t, err, ErrWordDoesNotExist)
	})
	t.Run("Delete a word that exists", func(t *testing.T) {
		word := "test"
		meaning := "this is not a test"
		dictionary := Dictionary{word: meaning}
		dictionary.Delete(word)
		_, err := dictionary.Search(word)
		assertError(t, err, ErrNotFound)
	})
}

func assertStringsEqual(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got != nil {
		if want != nil {
			return
		}
		t.Fatal("expected to get an error")
	}

}

func assertDefinition(t *testing.T, d Dictionary, word, meaning string) {
	t.Helper()
	got, err := d.Search(word)
	if err != nil {
		t.Errorf("Unable to add a new word to dictionary")
	}
	want := meaning
	assertStringsEqual(t, got, want)
}
