package arr

func SumInt(ints ...int) int {
	sum := 0

	for i := range ints {
		sum += i
	}

	return sum
}
