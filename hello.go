package main

import "fmt"

func main() {
	fmt.Println(Hello("Chris", ""))
}

func Hello(name, language string) string {
	const spanish = "Spanish"
	const prefixEnglish = "Hello, "
	const prefixSpanish = "Hola, "
	const suffix = "."

	if name == "" {
		name = "World"
	}

	if language == spanish {
		return prefixSpanish + name + suffix
	}

	return prefixEnglish + name + suffix
}
