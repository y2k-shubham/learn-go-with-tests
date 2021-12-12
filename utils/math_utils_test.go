package utils

import (
	"math"
	"testing"
)

func assertHelper(t testing.TB, got, want bool, f1, f2 float32) {
	if got != want {
		// format specifier for boolean https://stackoverflow.com/a/7059760/3679900
		t.Errorf("got %t, want %t for f1=%.2f, f2=%.2f", got, want, f1, f2)
	}
}

func TestAreEqual(t *testing.T) {
	t.Run("equal numbers", func(t *testing.T) {
		pi := float32(22) / 7

		got := AreEqual(pi, pi)
		want := true

		assertHelper(t, got, want, pi, pi)
	})

	t.Run("approx equal numbers 1", func(t *testing.T) {
		pi1 := float32(22) / 7
		pi2 := (float32(22) / 7) + 1e-10

		got := AreEqual(pi1, pi2)
		want := true

		assertHelper(t, got, want, pi1, pi1)
	})

	t.Run("approx equal numbers 2", func(t *testing.T) {
		pi1 := float32(22) / 7
		pi2 := (float32(22) / 7) + 1e-9

		got := AreEqual(pi1, pi2)
		want := true

		assertHelper(t, got, want, pi1, pi1)
	})

	t.Run("approx non equal numbers 1", func(t *testing.T) {
		pi1 := float32(22) / 7
		pi2 := (float32(22) / 7) + 1e-5

		got := AreEqual(pi1, pi2)
		want := false

		assertHelper(t, got, want, pi1, pi1)
	})

	t.Run("approx equal numbers 2", func(t *testing.T) {
		pi1 := float32(22) / 7
		pi2 := float32(math.Sqrt(10))

		got := AreEqual(pi1, pi2)
		want := false

		assertHelper(t, got, want, pi1, pi1)
	})
}
