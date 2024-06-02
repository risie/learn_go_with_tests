package main

import (
	"fmt"
)

const (
	french        = "French"
	spanish       = "Spanish"
	swedish       = "Swedish"
	englishPrefix = "Hello"
	spanishPrefix = "Hola"
	frenchPrefix  = "Bonjour"
	swedishPrefix = "Hej"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return fmt.Sprintf("%s, %s!", greetingPrefix(language), name)
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	case swedish:
		prefix = swedishPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", "Spanish"))
}
