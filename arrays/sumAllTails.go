package arrays

func SumAllTails(numList ...[]int) []int {
	var sums []int

	for _, value := range numList {
		if len(value) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(value[1:]))
		}
	}
	return sums
}
