package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	name := "name"
	englishGreeting := "Hello"
	spanishGreeting := "Hola"
	frenchGreeting := "Bonjour"
	swedishGreeting := "Hej"

	t.Run("saying hello to people", func(t *testing.T) {
		actual := Hello(name, "")
		expected := buildName(t, englishGreeting, name)

		if expected != actual {
			assertCorrectMessage(t, actual, expected)
		}
	})

	t.Run("default to 'world' when an empty string is supplied", func(t *testing.T) {
		actual := Hello("", "")
		expected := buildName(t, englishGreeting, "world")

		if expected != actual {
			assertCorrectMessage(t, actual, expected)
		}
	})

	t.Run("should default to english if language is empty", func(t *testing.T) {
		actual := Hello(name, "")
		expected := buildName(t, englishGreeting, name)

		if expected != actual {
			assertCorrectMessage(t, actual, expected)
		}
	})

	t.Run("should default to english if a unsupported language is entered", func(t *testing.T) {
		actual := Hello(name, "Swedish")
		expected := buildName(t, englishGreeting, name)

		if expected != actual {
			assertCorrectMessage(t, actual, expected)
		}
	})

	t.Run("should return a greeting in Spanish", func(t *testing.T) {
		actual := Hello(name, "Spanish")
		expected := buildName(t, spanishGreeting, name)

		if expected != actual {
			assertCorrectMessage(t, actual, expected)
		}
	})

	t.Run("should return a greeting in French", func(t *testing.T) {
		actual := Hello(name, "French")
		expected := buildName(t, frenchGreeting, name)

		if expected != actual {
			assertCorrectMessage(t, actual, expected)
		}
	})

	t.Run("should return a greeting in Swedish", func(t *testing.T) {
		actual := Hello(name, "Swedish")
		expected := buildName(t, swedishGreeting, name)

		if expected != actual {
			assertCorrectMessage(t, actual, expected)
		}
	})
}

func assertCorrectMessage(t testing.TB, actual, expected string) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}

func buildName(t testing.TB, greeting string, name string) string {
	t.Helper()
	return fmt.Sprintf("%s, %s!", greeting, name)
}
