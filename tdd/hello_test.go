package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("")

	want := "Hello, World!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloWithName(t *testing.T) {
	got := Hello("Abdel")
	want := "Hello, Abdel!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
