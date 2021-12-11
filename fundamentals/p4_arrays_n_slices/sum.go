package p4_arrays_n_slices

func Sum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}

	return sum
}
