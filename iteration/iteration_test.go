package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 4)
	expected := "aaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	iterate := 1
	for i := 0; i < b.N; i++ {
		Repeat("b", iterate)
		iterate += 1
		if iterate == 10 {
			iterate = 1
		}
	}
}
