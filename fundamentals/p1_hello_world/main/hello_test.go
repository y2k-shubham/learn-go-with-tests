package main

import "testing"

func TestHello(t *testing.T) {
	// For helper functions, it's a good idea to accept a testing.TB which is an interface
	// that *testing.T and *testing.B both satisfy, so you can call helper functions from a test, or a benchmark.
	assertCorrectMsg := func(t testing.TB, got, want string) {
		// t.Helper() is needed to tell the test suite that this method is a helper.
		// By doing this when it fails the line number reported will be in our function call
		// rather than inside our test helper.
		t.Helper()

		if got != want {
			t.Errorf("Got %q, want %q", got, want)
		}
	}

	t.Run("greeting Hello to people", func(t *testing.T) {
		var got string = Hello("Chris")
		var want string = "Hello Chris!"
		assertCorrectMsg(t, got, want)
	})

	t.Run("saying Hello World when there's no person to greet", func(t *testing.T) {
		var got string = Hello("")
		var want string = "Hello World!"
		assertCorrectMsg(t, got, want)
	})
}
