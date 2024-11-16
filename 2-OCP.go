// pen/Closed Principle (OCP)

// The Open/Closed Principle states that software entities (classes, modules, functions, etc.)
// should be open for extension but closed for modification.
// This principle encourages developers to write code that is flexible
//and can be extended without the need for significant modifications.

package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func Area(rectangle *Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

type Shape interface {
	Area() float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func main() {

}
