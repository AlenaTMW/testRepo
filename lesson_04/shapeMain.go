package main

import (
	"lesson_04/shapepkg"
)

func main() {
	c := shapepkg.Circle{Radius: 8.0}
	r := shapepkg.Rectangle{
		Height: 9.0,
		Width:  3.0,
	}
	shapepkg.DescribeShape(c)
	shapepkg.DescribeShape(r)
}
