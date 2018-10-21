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

// Calculates the totals of the tails of given slices 
func SumAllTails(slices ...[]int) []int {
	var sums []int
	for _, numbers := range slices {
		sum := 0
		if len(numbers) != 0 {
			tail := numbers[1:]
			sum = Sum(tail)
		}
		sums = append(sums, sum)
	}
	return sums
}
