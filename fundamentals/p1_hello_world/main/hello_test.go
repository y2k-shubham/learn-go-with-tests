package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("greeting Hello to people", func(t *testing.T) {
		var got string = Hello("Chris")
		var want string = "Hello Chris!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
