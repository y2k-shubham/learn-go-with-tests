package p5_structs_methods_interfaces

import "math"

type Rectangle struct {
	Length float32
	Width  float32
}

func (rectangle Rectangle) Perimeter() float32 {
	return 2 * (rectangle.Length + rectangle.Width)
}

func (rectangle Rectangle) Area() float32 {
	return rectangle.Length * rectangle.Width
}

type Circle struct {
	Radius float32
}

func (circle Circle) Perimeter() float32 {
	return 2 * math.Pi * circle.Radius
}

func (circle Circle) Area() float32 {
	return float32(math.Pi * math.Pow(float64(circle.Radius), 2))
}
