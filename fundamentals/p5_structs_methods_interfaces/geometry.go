package p5_structs_methods_interfaces

import "math"

type Shape interface {
	Perimeter() float32

	Area() float32
}

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

type Triangle struct {
	Base   float32
	Height float32
}

func (t Triangle) Perimeter() float32 {
	return 0
}

func (t Triangle) Area() float32 {
	return (t.Base * t.Height) / 2
}
