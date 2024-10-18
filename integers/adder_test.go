package integers

import "testing"

func TestAdder(t *testing.T) {

	t.Run("should be able to add two integers", func(t *testing.T) {
		sum := Add(1, 1)
		var expected float32 = 2.0
		assert(t, sum, expected)
	})

	t.Run("should be able to add two floating numbers", func(t *testing.T) {

		sum := Add(1, 1.2)
		var expected float32 = 2.2
		assert(t, sum, expected)
	})

}

func assert(t testing.TB, got float32, expected float32) {
	if got != expected {
		t.Errorf("expected %f. But got %f", expected, got)
	}
}
