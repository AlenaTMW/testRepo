package fibonacci

import "fmt"

func Number(z int) {

	n1, n2 := 0, 1
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
