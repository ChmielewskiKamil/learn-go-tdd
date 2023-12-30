package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
    Sleep()
}

type NormalSleeper struct {}

func (ns *NormalSleeper) Sleep() {
    time.Sleep(1 * time.Second)
}

const (
	countdownStart = 3
	finalWord      = "Go!"
)

func main() {
	Countdown(os.Stdout, &NormalSleeper{})
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
        sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}
