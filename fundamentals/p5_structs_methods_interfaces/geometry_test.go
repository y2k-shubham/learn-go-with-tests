package p5_structs_methods_interfaces

import "testing"

func assertHelper(t testing.TB, got, want float32) {
	if want != got {
		t.Errorf("got %f, want %f", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	got := Perimeter(11, 9)
	want := float32(40)

	assertHelper(t, got, want)
}
