package hello

import "testing"
import "fmt"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Gabriel", "")
		want := "Hello, Gabriel"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Barb", "Spanish")
		want := "Hola, Barb"
		assertCorrectMessage(t, got, want)
	})
	
	t.Run("in French", func(t *testing.T) {
		got := Hello("Jean-Paul", "French")
		want := "Salut, Jean-Paul"
		assertCorrectMessage(t, got, want)
	})

	t.Run("unknown language defaults to english", func(t *testing.T) {
		got := Hello("Julia", "Polish")
		want := "Hello, Julia"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Fatalf("got: %q | want: %q", got, want)
	}
}

func ExampleHello() {
	greeting := Hello("Gabriel", "French")
	fmt.Println(greeting)
	// Output: Salut, Gabriel
}
