package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const prefixEnglish = "Hello, "
const prefixSpanish = "Hola, "
const prefixFrench = "Bonjour, "
const suffix = "."

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name + suffix
}

func greetingPrefix(language string) string {
	switch language {
	case spanish:
        return prefixSpanish 
	case french:
		return prefixFrench 
	default:
		return prefixEnglish 
	}
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
