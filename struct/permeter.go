package structs

import "math"

//interface
type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

// method for reactangle
func (r Rectangle) Area() float64 {
	return (r.height * r.width)
}

type Circle struct {
	radius float64
}

//method for circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() float64 {
	return (.5 * t.base * t.height)
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}

// func Area(rectangle Rectangle) float64 {
// 	return (rectangle.width * rectangle.height)
// }
