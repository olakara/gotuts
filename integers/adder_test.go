package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(1, 1)
	expected := 2

	if sum != expected {
		t.Errorf("expected %q. But got %q", expected, sum)
	}
}
