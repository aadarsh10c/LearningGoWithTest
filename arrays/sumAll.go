package arrays

//Takes vriable arrays and returns an array with sum of each array
func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, number := range numbersToSum {
		sums = append(sums, Sum(number))
	}

	return
}
