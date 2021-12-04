package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMsg := func(t testing.TB, got, want string) {
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
