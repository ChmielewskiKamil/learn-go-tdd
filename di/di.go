package di

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprint(writer, "Hello, ", name)
}
