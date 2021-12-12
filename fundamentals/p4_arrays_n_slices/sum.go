package p4_arrays_n_slices

func Sum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}

	return sum
}

func SumAll(slcs [][]int) []int {
	if len(slcs) == 0 {
		return []int{}
	}

	sum := make([]int, len(slcs))
	for idx, slc := range slcs {
		sum[idx] = Sum(slc)
	}

	return sum
}

func SumAllTails(slcs [][]int) []int {
	if len(slcs) == 0 {
		return []int{}
	}

	sum := make([]int, len(slcs))
	for idx, slc := range slcs {
		sum[idx] = Sum(slc[1:])
	}

	return sum
}
