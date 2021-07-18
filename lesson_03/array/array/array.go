package array

import "fmt"

func ArraySum(a [8]int) {
	sum := 0
	for i := range a {
		sum += a[i]
	}
	fmt.Printf("%.1f\n", float64(sum)/float64(len(a)))
}
