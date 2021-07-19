package main

import (
	"fmt"
	"lesson_03/slices/1/thelongestString"
)

var s = []string{"one", "two", "three", "本日本日本"}

func main() {
	fmt.Println(thelongestString.Max(s))
}
