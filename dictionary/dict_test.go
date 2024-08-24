package dictionary

import "testing"

func TestSearch(t *testing.T) {
	custom_dict := MyDict{"test": "this is a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := custom_dict.Search("test")
		want := "this is a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := custom_dict.Search("unknown")
		// want := "could not find the word you were looking for"
		if got == nil {
			t.Fatal("expected to get an error")
		}
		assertError(t, got, ErrNotFound)
	})

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q ", got, want)
	}
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		myDict := MyDict{}
		word := "test"
		def := "this is a test"
		err := myDict.Add(word, def)
		assertError(t, err, nil)
		assertDefinition(t, myDict, word, def)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is a test"
		dict := MyDict{word: def}
		err := dict.Add(word, "new test")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, word, def)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		def := "this is a test"
		dict := MyDict{}
		err := dict.Update(word, def)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertDefinition(t testing.TB, dict MyDict, word, def string) {
	t.Helper()
	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should find added word: ", err)
	}
	assertStrings(t, got, def)
}

func TestUpdate(t *testing.T) {
	word := "test"
	def := "this is a test"
	dict := MyDict{word: def}

	newDef := "new definition"
	dict.Update(word, newDef)
	assertDefinition(t, dict, word, newDef)
}

func TestDelete(t *testing.T) {
	word := "test"
	dict := MyDict{word: "test definition"}
	dict.Delete(word)
	_, err := dict.Search(word)
	assertError(t, err, ErrNotFound)
}
