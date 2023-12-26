package main 

import "fmt"

func main() {
    fmt.Println(Hello("Chris"))
}

func Hello(name string) string {
    const prefix = "Hello, "
    return prefix + name + "."
}
