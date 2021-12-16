package sum

func Sum(arr []int) (got int) {
	for _, number := range arr {
		got += number
	}
	return
}

func SumAll(numbersToSum ...[]int) (got []int) {
	for _, numbers := range numbersToSum {
		got = append(got, Sum(numbers))
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (got []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			numbers = append(numbers, 0)
		}
		tail := numbers[1:]
		got = append(got, Sum(tail))
	}
	return
}
