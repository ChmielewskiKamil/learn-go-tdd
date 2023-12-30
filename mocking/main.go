package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
    Countdown(os.Stdout)
}

func Countdown(out io.Writer) {
    fmt.Fprint(out, "3\n")
    fmt.Fprint(out, "2\n")
    fmt.Fprint(out, "1\n")
    fmt.Fprint(out, "Go!")
}
