package p5_structs_methods_interfaces

import "testing"

func assertHelper(t testing.TB, got, want float32) {
	// shouldn't use this direct floating point
	if want != got {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{11, 9}

	got := Perimeter(rectangle)
	want := float32(40)

	assertHelper(t, got, want)
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{11, 9}

	got := Area(rectangle)
	want := float32(99)

	assertHelper(t, got, want)
}
