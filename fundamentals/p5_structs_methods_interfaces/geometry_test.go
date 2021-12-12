package p5_structs_methods_interfaces

import (
	"learn-go-with-tests/utils"
	"testing"
)

func assertHelper(t testing.TB, got, want float32) {
	t.Helper()

	// shouldn't use this direct floating point
	if !utils.AreEqual(got, want) {
		t.Errorf("got %g, want %g", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, shape Shape, want float32) {
		t.Helper()
		got := shape.Perimeter()
		assertHelper(t, got, want)
	}

	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{11, 9}
		want := float32(40)
		checkPerimeter(t, rectangle, want)
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10}
		want := float32(62.831856)
		checkPerimeter(t, circle, want)
	})
}

func TestArea(t *testing.T) {
	// normal tests
	checkArea := func(t testing.TB, shape Shape, want float32) {
		t.Helper()
		got := shape.Area()
		assertHelper(t, got, want)
	}

	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{11, 9}
		want := float32(99)
		checkArea(t, rectangle, want)
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10}
		want := float32(314.159265359)
		checkArea(t, circle, want)
	})

	// table tests
	areaTests := []struct {
		shape Shape
		want  float32
	}{
		{Rectangle{11, 9}, 99},
		{Circle{10}, 314.159265359},
		{Triangle{3, 14}, 21},
	}

	for _, test := range areaTests {
		checkArea(t, test.shape, test.want)
	}
}
