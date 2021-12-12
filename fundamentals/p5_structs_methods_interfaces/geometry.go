package p5_structs_methods_interfaces

func Perimeter(rectangle Rectangle) float32 {
	return 2 * (rectangle.Length + rectangle.Width)
}

func Area(rectangle Rectangle) float32 {
	return rectangle.Length * rectangle.Width
}

type Rectangle struct {
	Length float32
	Width  float32
}
