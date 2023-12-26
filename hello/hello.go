package hello 

import "fmt"

const (
	spanish       = "Spanish"
	french        = "French"

	prefixEnglish = "Hello, "
	prefixSpanish = "Hola, "
	prefixFrench  = "Bonjour, "

	suffix        = "."
)

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
