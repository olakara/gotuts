package iterations

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("should repeat the character 5 times when input is 5", func(t *testing.T) {
		repeated := Repeat("a", 5)
		result := "aaaaa"

		if repeated != result {
			t.Errorf("Expected %q but got %q", result, repeated)
		}
	})

	t.Run("should repeat the character twice when input is 2", func(t *testing.T) {
		repeated := Repeat("a", 2)
		result := "aa"

		if repeated != result {
			t.Errorf("Expected %q but got %q", result, repeated)
		}
	})

	t.Run("should return empty string if asked to repeat 0 times", func(t *testing.T) {
		repeated := Repeat("a", 0)
		result := ""

		if repeated != result {
			t.Errorf("Expected 0 but got %q", repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", b.N)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("b", 3))
	//Output: bbb
}
