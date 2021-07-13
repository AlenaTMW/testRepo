package main

import (
	"fmt"

	"lesson_02/lesson_02/fibonacci"
)

func main() {
	defer sayComplete()
	fmt.Println("Hello!")
	fibonacci.Number(5)
}

func sayComplete() {
	fmt.Println("Completed")
}
