package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("greets with the right name", func(t *testing.T) {
		got := Hello("Doge", "")
		want := "Hello, Doge"
		assertCorrectMessage(t, got, want)
	},
	)

	t.Run("say 'Hello, Cruel World' if empty string is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, Cruel World"
		assertCorrectMessage(t, got, want)
	},
	)

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Gandalf", "Spanish")
		want := "Hola, Gandalf"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Legolas", "French")
		want := "Bonjour, Legolas"
		assertCorrectMessage(t, got, want)
	})
}
