package main

import (
	"bytes"
	"reflect"
	"testing"
)

type SpyCountdownOperation struct {
	Calls []string
}

func (ss *SpyCountdownOperation) Sleep() {
	ss.Calls = append(ss.Calls, sleep)
}

func (ss *SpyCountdownOperation) Write(p []byte) (n int, err error) {
	ss.Calls = append(ss.Calls, write)
	return
}

const (
	sleep = "sleep"
	write = "write"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountdownOperation{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleeperPrinter := &SpyCountdownOperation{}
		Countdown(spySleeperPrinter, spySleeperPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleeperPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleeperPrinter.Calls)
		}
	})
}
