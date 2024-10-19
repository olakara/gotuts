package arrays

import "testing"

func TestSum(t *testing.T) {

	t.Run("should return sum of integers", func(b *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		sum := Sum(numbers[:])
		expected := 15

		if sum != expected {
			t.Errorf("Expected to get %d, but got %d", expected, sum)
		}
	})

	t.Run("should return 0 when there is no numbers passed", func(t *testing.T) {
		numbers := []int{}
		sum := Sum(numbers)
		expected := 0
		if sum != expected {
			t.Errorf("Expected to get 0, but got %d", sum)
		}

	})
}
