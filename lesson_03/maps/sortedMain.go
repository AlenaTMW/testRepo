package main

import (
	"fmt"
	"lesson_03/maps/sorted"
)

var name1 = map[int]string{2: "a", 0: "b", 1: "c"}
var name2 = map[int]string{10: "aa", 0: "bb", 500: "cc"}

func main() {
	sorted.PrintSorted(name1)
	fmt.Println()
	sorted.PrintSorted(name2)
}
