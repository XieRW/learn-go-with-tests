package structsAndMethodsAndInterfaces

import "math"

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.width * rectangle.height
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

type Shape interface {
	Area() float64
}

type Triangle struct {
	width  float64
	height float64
}

func (t Triangle) Area() (got float64) {
	got = t.height * t.width / 2
	return
}
