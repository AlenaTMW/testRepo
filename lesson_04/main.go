package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	height float64
	width  float64
}

func DescribeShape(s Shape) {
	fmt.Println(s)
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

//func (c Cat) String() string {return fmt.Sprintf("Cat has %d legs\n", c.legs)}
func (c Circle) String() string {
	return fmt.Sprintf("\nCirlce: radius %.2f", c.radius)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("\nRectangle with height %.2f and width %.2f", r.height, r.width)
}

func main() {
	// choose your own dimensions
	c := Circle{radius: 8.0}
	r := Rectangle{
		height: 9.0,
		width:  3.0,
	}
	DescribeShape(c)
	DescribeShape(r)
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}
