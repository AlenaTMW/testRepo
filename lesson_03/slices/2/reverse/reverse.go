package reverse

func Reverse(y []int64) []int64 {
	x := make([]int64, len(y))
	for i := range y {
		x[len(y)-1-i] = y[i]
	}
	return x

}
