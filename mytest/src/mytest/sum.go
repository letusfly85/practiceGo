package mytest

func MySum(xs ...int) int {
	acc := 0

	for _, val := range xs {
		acc = acc + val
	}
	return acc
}

func MyFoldSum(xs []int) int {
	switch {
	case len(xs) == 1:
		return xs[0]

	default:
		head, tail := xs[0], xs[1:]
		return (head + MyFoldSum(tail))
	}
}
