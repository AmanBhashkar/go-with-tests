package structsmethodsandinterfaces

import "math"

type Shape interface {
	Area() float64
}
type Rectangle struct {
	Length float64
	Width  float64
}
type Circle struct {
	Redius float64
}
type Triangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Redius * c.Redius
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
