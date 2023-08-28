package arrays

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, number := range numbersToSum {
		if len(number) == 0 {
			sums = append(sums, 0)
		} else {
			tail := number[1:]
			sums = append(sums, Sum(tail))
		}

	}
	return
}
