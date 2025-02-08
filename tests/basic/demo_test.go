package main

import "testing"

func Test_IsOdd_ReturnsTrueWhenOdd(t *testing.T) {
	t.Parallel()
	want := true
	got := isOdd(3)
	if got != want {
		t.Fatalf("isOdd(3): got %v want %v", got, want)
	}
}

func Test_IsOdd_ReturnsFalseWhenEven(t *testing.T) {
	t.Parallel()
	want := false
	got := isOdd(4)
	if got != want {
		t.Fatalf("isOdd(4): got %v want %v", got, want)
	}
}
