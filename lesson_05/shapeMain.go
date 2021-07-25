package main

import (
	"lesson_05/shapedescription"
)

func main() {
	c := shapedescription.Circle{Radius: 8.0}
	r := shapedescription.Rectangle{
		Height: 9.0,
		Width:  3.0,
	}
	shapedescription.DescribeShape(c)
	shapedescription.DescribeShape(r)
}
