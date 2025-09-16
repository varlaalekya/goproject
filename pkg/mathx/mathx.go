package mathx

func Sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func IsEven(n int) bool {
	return n%2 == 0
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	    return b
}