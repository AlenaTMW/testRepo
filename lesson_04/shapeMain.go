package main

import (
	"lesson_04/shapedescription"
)

func main() {
	c := shapedescription.Circle{Radius: 8.0}
	r := shapedescription.Rectangle{
		Height: 9.0,
		Width:  3.0,
	}
	q := shapedescription.Square{Length: 2.0}
	shapedescription.DescribeShape(c)
	shapedescription.DescribeShape(r)
	shapedescription.DescribeShape(q)
}
