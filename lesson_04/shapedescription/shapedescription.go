package shapedescription

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Height float64
	Width  float64
}

type Square struct {
	Length float64
}

func DescribeShape(s Shape) {
	fmt.Println(s)
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func (c Circle) String() string {
	return fmt.Sprintf("\nCirlce: Radius %.2f", c.Radius)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("\nRectangle with height %.2f and width %.2f", r.Height, r.Width)
}

func (q Square) String() string {
	return fmt.Sprintf("\nSquare with length %.2f", q.Length)
}

func (C Circle) Area() float64 {
	return math.Pi * math.Pow(C.Radius, 2)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (q Square) Area() float64 {
	return math.Pow(q.Length, 2)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}
func (q Square) Perimeter() float64 {
	return 4 * q.Length
}
