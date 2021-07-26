package shapedescription

import (
	"errors"
	"fmt"
	"math"
)

type Shape interface {
	Area() (float64, error)
	Perimeter() (float64, error)
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Height float64
	Width  float64
}

func DescribeShape(s Shape) {
	fmt.Println(s)
	a, err := s.Area()
	if err == nil {
		fmt.Printf("Area: %.2f\n", a)
	} else {
		fmt.Println(err)
	}
	p, err := s.Perimeter()
	if err == nil {
		fmt.Printf("Perimeter: %.2f\n", p)
	} else {
		fmt.Println(err)
	}

}

func (c Circle) String() string {
	return fmt.Sprintf("\nCirlce: Radius %.2f", c.Radius)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("\nRectangle with height %.2f and width %.2f", r.Height, r.Width)
}

func (c Circle) Area() (float64, error) {
	if c.Radius == 0 {
		return 0.0, errors.New("value is invalid")
	}

	if c.Radius < 0 {
		return 0.0, errors.New("value is invalid")
	}
	return math.Pi * math.Pow(c.Radius, 2), nil
}

func (r Rectangle) Area() (float64, error) {
	if r.Height <= 0 || r.Width <= 0 {
		return 0.0, errors.New("value is invalid")
	}
	return r.Height * r.Width, nil
}

func (c Circle) Perimeter() (float64, error) {
	if c.Radius == 0 {
		return 0.0, errors.New("value is invalid")
	}

	if c.Radius < 0 {
		return 0.0, errors.New("value is invalid")
	}

	return 2 * math.Pi * c.Radius, nil

}

func (r Rectangle) Perimeter() (float64, error) {
	if r.Height <= 0 || r.Width <= 0 {
		return 0.0, errors.New("value is invalid")
	}
	return 2 * (r.Height + r.Width), nil
}
