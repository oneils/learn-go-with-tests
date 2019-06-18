package iteration

import "testing"

func TestRepeat(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("should repeat 1 time", func(t *testing.T) {
		repeated := Repeat("a", 1)
		expected := "a"

		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("should repeat by specified times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertCorrectMessage(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
