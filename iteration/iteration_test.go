
package iteration

import "testing"

func TestRepeat(t *testing.T) {

	t.Run("testing if the character repeats 10 times", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"
		AssertCorrectRepetition(t, repeated, expected)
	})

	t.Run ("testing if the function defaults to character if 'times' < 1", func (t *testing.T) {
		repeated := Repeat("b", -1)
		expected := "b"
		AssertCorrectRepetition(t, repeated, expected)
	})
}


func AssertCorrectRepetition(t testing.TB, repeated, expected string) {
	t.Helper()
	if expected != repeated {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}


func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 10)
	}
}
