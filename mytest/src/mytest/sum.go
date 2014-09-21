package mytest

func MySum(xs ...int) int {
	acc := 0

	for _, val := range xs {
		acc = acc + val
	}

	return acc
}
