package p5_structs_methods_interfaces

import (
	"learn-go-with-tests/utils"
	"testing"
)

func assertHelper(t testing.TB, got, want float32) {
	// shouldn't use this direct floating point
	if !utils.AreEqual(got, want) {
		t.Errorf("got %g, want %g", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{11, 9}

		got := rectangle.Perimeter()
		want := float32(40)

		assertHelper(t, got, want)
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10}

		got := circle.Perimeter()
		want := float32(62.831856)

		assertHelper(t, got, want)
	})
}

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{11, 9}

		got := rectangle.Area()
		want := float32(99)

		assertHelper(t, got, want)
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10}

		got := circle.Area()
		want := float32(314.159265359)

		assertHelper(t, got, want)
	})
}
