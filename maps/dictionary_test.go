package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "just a test"}

	key := "test"
	got := Search(dictionary, key)
	want := "just a test"

	if got != want {
		t.Errorf("got %q want %q given %q", got, want, key)
	}
}
