package main 

import "testing"

// `t` is the hook into testing framework
func TestHello(t *testing.T) {
    want := "Hello, world."
    got := Hello()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
