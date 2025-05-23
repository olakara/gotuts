package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Abdel", "")
		want := "Hello, Abdel!"
		assertMessage(t, want, got)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertMessage(t, want, got)
	})

	t.Run("say hello to people in French", func(t *testing.T) {
		got := Hello("Abdel", "French")
		want := "Bonjour, Abdel!"
		assertMessage(t, want, got)
	})
}

func assertMessage(t testing.TB, want string, got string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
