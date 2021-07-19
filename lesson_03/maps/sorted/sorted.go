package sorted

import (
	"fmt"
	"sort"
)

func PrintSorted(name1 map[int]string) {

	keys := make([]int, 0, len(name1))
	for k := range name1 {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		fmt.Print(name1[k], " ")
	}
}
