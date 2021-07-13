package fibonacci

import "fmt"

func Number(z int) {

	var n1 int = 0
	var n2 int = 1
	var n int

	for i := 1; i <= z; i++ {
		if i == 1 {
			fmt.Print(n1, " ")
			continue
		}
		if i == 2 {
			fmt.Print(n2, " ")
			continue
		}
		fmt.Print(n1+n2, " ")
		n = n1
		n1 = n2
		n2 = n + n2

	}
	fmt.Println("")
}
