package iterations

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	result := "aaaaa"

	if repeated != result {
		t.Errorf("Expected %q but got %q", result, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
