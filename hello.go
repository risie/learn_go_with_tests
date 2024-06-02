package main

import "fmt"

const (
	englishHelloPrefix = "Hello"
	spanishHelloPrefix = "Hola"
	frenchHelloPrefix  = "Bonjour"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == "Spanish" {
		return fmt.Sprintf("%s, %s!", spanishHelloPrefix, name)
	}

	if language == "French" {
		return fmt.Sprintf("%s, %s!", frenchHelloPrefix, name)
	}

	return fmt.Sprintf("%s, %s!", englishHelloPrefix, name)
}

func main() {
	fmt.Println(Hello("world", "Spanish"))
}
