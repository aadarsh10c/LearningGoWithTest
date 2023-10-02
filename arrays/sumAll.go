package arrays

func SumAll(numList ...[]int) []int {
	var sumAll []int

	for _, value := range numList {
		sumAll = append(sumAll, Sum(value))
	}
	return sumAll
}
