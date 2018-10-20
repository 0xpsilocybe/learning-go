package slices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	slicesCount := len(slices)
	sums := make([]int, slicesCount)
	for i, numbers := range slices {
		sums[i] = Sum(numbers)
	}
	return sums
}

