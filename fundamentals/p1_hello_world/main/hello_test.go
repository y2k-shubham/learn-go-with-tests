package main

import "testing"

func assertHelper(t testing.TB, got, want string) {
	// t.Helper() is needed to tell the test suite that this method is a helper.
	// By doing this when it fails the line number reported will be in our function call
	// rather than inside our test helper.
	t.Helper()

	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}

func TestGreeting(t *testing.T) {
	t.Run("greeting in English", func(t *testing.T) {
		var got string = Greeting(English)
		var want string = "Hello"
		assertHelper(t, got, want)
	})

	t.Run("greeting in Spanish", func(t *testing.T) {
		var got string = Greeting(Spanish)
		var want string = "Hola"
		assertHelper(t, got, want)
	})

	t.Run("greeting in French", func(t *testing.T) {
		var got string = Greeting(French)
		var want string = "Bonjour"
		assertHelper(t, got, want)
	})

	t.Run("greeting in unknown lang", func(t *testing.T) {
		var got string = Greeting(0)
		var want string = "Hello"
		assertHelper(t, got, want)
	})
}

func TestWorld(t *testing.T) {
	t.Run("world in English", func(t *testing.T) {
		var got string = World(English)
		var want string = "World"
		assertHelper(t, got, want)
	})

	t.Run("world in Spanish", func(t *testing.T) {
		var got string = World(Spanish)
		var want string = "Mundo"
		assertHelper(t, got, want)
	})

	t.Run("world in French", func(t *testing.T) {
		var got string = World(French)
		var want string = "Monde"
		assertHelper(t, got, want)
	})

	t.Run("world in unknown lang", func(t *testing.T) {
		var got string = World(0)
		var want string = "World"
		assertHelper(t, got, want)
	})
}

func TestHello(t *testing.T) {
	t.Run("greeting Hello to people", func(t *testing.T) {
		var got string = Hello("Chris", English)
		var want string = "Hello Chris!"
		assertHelper(t, got, want)
	})

	t.Run("greeting in different language", func(t *testing.T) {
		var got string = Hello("Tulika", French)
		var want string = "Bonjour Tulika!"
		assertHelper(t, got, want)
	})

	t.Run("saying Hello World when there's no person to greet", func(t *testing.T) {
		var got string = Hello("", Spanish)
		var want string = "Hola Mundo!"
		assertHelper(t, got, want)
	})

	t.Run("falling back to English when there's no language specified", func(t *testing.T) {
		var got string = Hello("", 0)
		var want string = "Hello World!"
		assertHelper(t, got, want)
	})
}
