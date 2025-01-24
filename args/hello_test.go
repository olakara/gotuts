package hello

import "testing"

func TestSay(t *testing.T) {
	want := "Hello John"
	got := Say("John")
	if got != want {
		t.Errorf("Say() = %q, want %q", got, want)
	}
}
