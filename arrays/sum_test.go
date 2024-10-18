package arrays

import "testing"

func TestSum(t *testing.T) {

	t.Run("should return sum of integers", func(b *testing.T) {
		sum := Sum()
		expected := 40

		if sum != expected {
			t.Errorf("Expected to get %d, but got %d", expected, sum)
		}
	})
}
