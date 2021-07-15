package main

import (
	"fmt"

	"lesson_02/lesson_02/fibonacci"
)

func main() {
	defer fmt.Println("Completed")
	fmt.Println("Hello!")
	fibonacci.Number(5)
}
