package main

import "testing"

// `t` is the hook into testing framework
func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		want := "Hello, Chris."
		got := Hello("Chris")

        assertCorrectMessage(t, got, want)
	})

    t.Run("say 'Hello, World' for empty string", func(t *testing.T) {
        want := "Hello, World."
        got := Hello("")

        assertCorrectMessage(t, got, want)
    })
}

// TB means that we can call this helper from both tests and benchmarks
func assertCorrectMessage(t testing.TB, got string, want string) {
    t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
