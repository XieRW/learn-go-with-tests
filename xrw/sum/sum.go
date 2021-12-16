package sum

func Sum(arr []int) (got int) {
	for _, number := range arr {
		got += number
	}
	return
}
